dev: 
	make -j5 live/templ live/server live/tailwind live/sync_assets

format:
	go fmt myapp/...

build:
	go build -o bin/main main.go

run:
	go run main.go

templ-watch:
	templ generate --watch --proxy="http://localhost:3001" --cmd="go run main.go"

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:3001" --proxybind="localhost" --proxyport="3000" --open-browser=false

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
live/esbuild:
	npx esbuild js/index.ts --bundle --outdir=assets/ --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

