MAIN_PKG = cmd/main.go

BUILD_DIR = dist

CONFIG_PATH = config/local.yaml

serve:
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}
	${BUILD_DIR}/main.exe -config=${CONFIG_PATH}

build: 
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}

swag:
	swag init --parseDependency --parseInternal --parseDepth 2 -d "./internal/http-server" -g "http-server.go" -o "./docs"
