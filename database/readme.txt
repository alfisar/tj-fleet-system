*membuat file migration* 
migrate create -ext sql -dir database/migrations <nama filenya>

*migrasi semua file migration*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable"
" -path database/migrations up

*rollback semua file migration*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -path database/migrations down

*migrasi 1 version*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -path database/migrations up 1

*rollback 1 version*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -path database/migrations down 1

*check version migrate*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -path database/migrations version

*force update migrate yang dirty*
migrate -database "postgres://<username>:<password>@<host>:<port>/<database>?sslmode=disable" -path database/migrations force <angka dari namanya >