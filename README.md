# arithmetic-rest-api
## Пояснение к решению
Команда для сборки образа:
```docker
  docker build -t arithmetic-rest-api .
```
Команда для создания и запуска контейнера в интерактивном режиме:
```docker
  docker run -it -p 8081:8080 arithmetic-rest-api
```
Можно запустить тест во время работы контейнера командой:
```go
  go test -v
```