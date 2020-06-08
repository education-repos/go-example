# Простой пример работы с go на примере HTTP веб-сервера

## Запуск

- отредактировать конфиг `internal\configs\dev-config.yml`

- обратите внимание на структуру `Category` - вам необходимо создать таблицу в БД с подобными полями (`id=3`)

- `go mod vendor`

- запустить сервер `go run main.go`

## Описание

По дефолту сервер запускается на `localhost:8080`

Реализовано несколько разных хендлеров для работы сервера:

- MainHandler (`internal\router\handlers\main_handler.go`) - отрабатывает все endpoint's, которые не попадают в другие правила
    
    * http://localhost:8080/
    
    * http://localhost:8080/?param=test&key=153

- ImgHandler (`internal\router\handlers\img_handler.go`) - пример работы со статикой в go
    
    * http://localhost:8080/image
    
- ApiHandler (`internal\router\handlers\api_handler.go`) - пример запроса на внешний сервис и вывода ответа в `json`
    
    * http://localhost:8080/api
    
- DbHandler (`internal\router\handlers\db_handler.go`) - пример работы с mysql базой. делается селект и полученная информация выводится в виде json

    * http://localhost:8080/db
    
За роутинг отвечает файл `internal\router\router.go`

За роутинг статических файлов отвечает `internal\router\handlers\static_handler.go`

За парсинг и работу с конфигами отвечает `internal\configs\config.go`

