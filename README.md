# to run locally:
## env
   open .env and change db_host to localhost
  `docker compose up go_db -d`
## to run migartions: 
  `goose -dir ./migrations up`
## run app
    `go run cmd/main.go`
# to run through docker
`docker compose up`
## however there are connection issues between db and app so run locally
Sorry
