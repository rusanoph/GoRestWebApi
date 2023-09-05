# GoWebApiApp

**Must have:**
* (Go Best Practices)[https://github.com/codeship/go-best-practices]

---

Создание REST API.

Golang WebApp. ApiServer

Урок 1. Веб-сервер
* Узнал о makefile и научился с ними работать. Узнал о флагах .SILENT И .PHONY
* Научился создавать простые тести для http сервера c помощью библиотеки testify
* Научился конфигурировать компилятор go. Необходимость возникла при попытке тестирования приложения, так как при вызове команды `go test -v -race <path>` компилятор ругался сначала на отсутствие 64-битной версии gcc компилятора, а после установки нужного gcc под x86_64 ему не нравилось, что не установлена переменная окружения go компилятора CGO_ENABLED=0, он требовал 1. После установки переменной команда исправно отработала и начали прогонятся тести. Например изменение переменной окружения: `go env -w <variable>=<value>`, где `-w` флаг записи write. `go env` - выводит список всех переменных. `go env <variable>` - выводит значение переменной
* Узнал о библиотеке для роутинга веб сервера mux. Ничего не понятно, надо углублятся
* Узнал о формате .toml и об одноименной библиотеке для работы с этим конфиг форматом. Есть ещё варианты конфигурировать приложение с помощью .yml и .json
* Узнал о библиотеке для логгирования работы сервера logrus, как и с mux, ничего не ясно, необходимо будет дополнительно углубится, почему нельзя использовать для этого стандартный log. Что за библиотека zerolog
* Также изучил способ передавать в уже собранное приложение флаги с помощью библиотеки flags. Установил флаг, позволяющий передавать файл конфига. Подробнее об этой библе (тут)[https://pkg.go.dev/flag#Parse]

Список использованных библиотек: testify (testing, assert), mux, toml, logrus, zerolog, flags, net (http, http/httptest), io, log

---

Урок 2. Работа с БД

* На данный момент научился создавать подключение к бд PostgreSQL, расширил конфиги
* Научился создавать миграции, для этого использовал инструмент migrate. Для этого скачал менеджер пакетов scoop, он аналогичен choco. При создании миграций возникали проблемы: 
    * С неправильной строкой подключения (тут чисто я накосячил). Решение: postgres://postgres:pass@localhost:5432/restapi_db?sslmode=disable; 
    * Error: Dirty database version X. Fix and force version. Решение: можно снести саму базу, либо удалить в ней системную таблицу мигнаций (должны быть более цивилизованные способы решения проблемы, чем полное удаление);
    * No migration found for version X. Решение: использовать команду force для переключения версии.  
* Приобрел навык работы с postgres через терминал. Например команды: \dt, \l, \c <dbname>

Полезные ссылки по миграциям: 
* (Golang-migrate returns "no change")[https://stackoverflow.com/questions/75053949/golang-migrate-returns-no-change]
* (Migrate psql tutorial)[https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md]
* (Чиним dirty migration с помощью restore api)[https://support.hashicorp.com/hc/en-us/articles/15166228647571--no-migration-found-for-version-Error-Encountered-After-Using-the-Restore-API-to-Restore-Terraform-Enterprise]
* (FreeCodeCamp migration guide)[https://www.freecodecamp.org/news/database-migration-golang-migrate/]

Список новых tool'ов: migrate, scoop, psql, lib/pq

Также есть инфа по использованию lib'ы (database/sql)[http://go-database-sql.org/]
