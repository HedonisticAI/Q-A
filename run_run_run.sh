go mod tidy
docker compose up --no-start
docker compose start db
goose -dir ./migrations -table goose_migrations up
#docker start app
