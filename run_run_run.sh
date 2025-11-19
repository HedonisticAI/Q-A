go mod tidy
docker compose up db -d
goose -dir ./migrations -table goose_migrations up
#docker start q-a-app
