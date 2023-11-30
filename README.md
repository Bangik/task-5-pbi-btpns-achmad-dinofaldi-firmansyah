# TASK 5 PBI BTPNS ACHMAD DINOFALDI FIRMANSYAH

## Requirements
- Golang
- Postgres

## Links
- [API](http://157.245.48.137:8080)
- [Swagger Docs](https://app.swaggerhub.com/apis-docs/Bangik/rakamin-btpns/1.0.0)

## How to run
1. Clone this repository
2. Open terminal and go to the directory
3. Setup environment variable in .env file
4. Run `go mod tidy` to install dependencies
5. Run `go run main.go` to run the program

## How to run with docker-compose
1. Clone this repository
2. Open terminal and go to the directory
3. Setup environment variable in docker-compose.yml file
4. Run `docker-compose up` to run the program

## How to run test
1. Open terminal and go to the directory
2. Run `go test -v ./... -coverprofile=cover.out && go tool cover -html=cover.out` to run all test
