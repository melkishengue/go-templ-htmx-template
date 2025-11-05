frontend: 
	air -c cmd/frontend/.air.toml
backend: 
	air -c cmd/backend/.air.toml
openapi: 
	swag init -g cmd/backend/main.go -o spec && docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/spec/swagger.json \
    -g openapi \
    -o /local/spec \
    --openapi-normalizer \
    REMOVE_DUPLICATE_OPERATION_IDS=true && sed -i '' 's|"//localhost|"http://localhost|g' "spec/openapi.json" && sed -i '' 's|"//|"https://|g' "spec/openapi.json"
start: 
	make -j3 openapi backend frontend -d
start-frontend-docker:
	docker build -f ./cmd/frontend/Dockerfile -t myapp-frontend:latest . &&  docker run --env-file .env -p 3010:3010 myapp-frontend:latest
start-backend-docker:
	docker build -f ./cmd/backend/Dockerfile -t myapp-backend:latest . &&  docker run --env-file .env -p 3020:3020 myapp-backend:latest
