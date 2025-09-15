SWAGGER_DIRS=cmd/server,internal/server
SWAGGER_MAIN=cmd/server/main.go

.PHONY: swagger

swrun: swagger
	go run $(SWAGGER_MAIN)

swagger:
	swag init -d $(SWAGGER_DIRS)
