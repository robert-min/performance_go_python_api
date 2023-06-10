module github.com/robert-min/performance_go_python_api/go_api_db/server

go 1.19

replace github.com/robert-min/performance_go_python_api/go_api_db/server/handlers => ./handlers

replace github.com/robert-min/performance_go_python_api/go_api_db/server/lib => ./lib

require github.com/go-sql-driver/mysql v1.7.1
