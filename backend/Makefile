SWAGGER_DIRS=cmd/server,internal/server
MAIN_FILE=cmd/server/main.go


.PHONY: run run-swagger

run:
	go run $(MAIN_FILE)

swagger: # swagger docs
	swag init -d $(SWAGGER_DIRS)
