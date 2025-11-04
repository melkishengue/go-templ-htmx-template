package middleware

import (
	"bytes"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type CachedResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
}

var reqCache = cache.New(5*time.Minute, 10*time.Minute)
var ttl = time.Minute * 15

type responseWriterWithCache struct {
	gin.ResponseWriter
	buf        *bytes.Buffer
	statusCode int
}

func (w *responseWriterWithCache) Write(b []byte) (int, error) {
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriterWithCache) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// in-memory cache middleware
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		noCacheParam := c.Query("no_cache")
		if noCacheParam != "" || c.Request.Method != "GET" || !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		cacheKey := c.Request.URL.String()

		if cached, found := reqCache.Get(cacheKey); found {
			cachedResponse := cached.(*CachedResponse)
			c.Header("Content-Type", cachedResponse.ContentType)
			c.Data(cachedResponse.StatusCode, cachedResponse.ContentType, cachedResponse.Body)
			c.Abort()
			return
		}

		rw := &responseWriterWithCache{
			ResponseWriter: c.Writer,
			buf:            &bytes.Buffer{},
			statusCode:     200,
		}
		c.Writer = rw

		c.Next()

		contentType := rw.Header().Get("Content-Type")

		if strings.Contains(contentType, "application/json") && rw.buf.Len() > 0 {
			responseToCache := &CachedResponse{
				Body:        rw.buf.Bytes(),
				ContentType: contentType,
				StatusCode:  rw.statusCode,
			}
			reqCache.Set(cacheKey, responseToCache, ttl)
		}
	}
}
