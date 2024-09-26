# kanban-backend

## Инструкция по запуску
### 1) В каталоге config копируем файл `local.example.yaml` в `local.yaml`
### 2) При помощи утилиты make запускаем сервер:
#### a) `make swag` - генерирует сваггер документацию
#### b) `make serve` - запускает сервер с конфигурацией local.yaml
#### c) `make build`- билдит проект
### 3) Без утилиты make:
#### a) `swag init --parseDependency --parseInternal --parseDepth 2 -d "./internal/http-server" -g "http-server.go" -o "./docs"` - генерирует сваггер документацию
#### b) `go build -o dist/main.exe cmd/main.go dist/main.exe -config=config/local.yaml` - запускает сервер с конфигурацией local.yaml
#### c) `go build -o dist/main.exe cmd/main.go`- билдит проект