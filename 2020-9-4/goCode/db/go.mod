module db

go 1.14

replace redis => ./redis

replace sqlx => ./sqlx

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	redis v0.0.0-00010101000000-000000000000 // indirect
	sqlx v0.0.0-00010101000000-000000000000 // indirect
)
