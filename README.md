запуск бд для разработки
docker run --name=todo-db -e POSTGRES_PASSWORD='devel' -d -p 5432:5432 --rm postgres

миграции

с помощью утилиты goilang-migrate

установка:
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 $GOPATH/bin/migrate

Создаем миграцию
migrate create -ext sql -dir ./schema -seq init

Прогоняем миграцию
migrate -path ./schema -database 'postgres://postgres:devel@192.168.0.21:5432/postgres?sslmode=disable' up