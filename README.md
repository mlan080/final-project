# Training 6: Final Project

For the last part of the backend training, we have a small project that combines a lot of the skills
from the previous 5 sessions. For this one there’s no skeleton -- you have the skills to create the
whole project from scratch now.


## The service

For this project, you’ll create a service with a REST API, a bit like the ones in session 4 and 5.
It’ll load a database of airports from a CSV file when it starts, and make them available over HTTP.
The file is `airport-data.csv` in this folder.


## API endpoints

Your REST API should have all the endpoints as described in `api.html` file. You can open this file
in a web browser.

## Tasks

1. Create a Go module that loads the airports from the CSV file. We don’t need all data from the CSV
   file: only the `name`, `iata_code`, `latitude_deg` and `longitude_deg` columns. Also, we only
   want airports of type `large_airport`. Make sure you have proper error handling -- if something
   goes wrong (e.g. the file can’t be opened), it should return an `error` value with a suitable
   error message.
2. Create a database type that wraps a map with the airport data. As a minimum, it should have
   methods that make it easy to look up an airport by its IATA code and to add an airport.
3. Create a REST API that lets us access the airports.
4. Write unit tests for your code. 100% test coverage is not required, but make sure that your tests
   are able to catch any potential bugs in the code.
5. Write a Dockerfile and docker-compose.yml file so your service can be run in development with a
   single command i.e. `docker-compose up`.
6. Deploy your service to the `a2b-exp` Kubernetes cluster. Use the namespace `x-backend-training`
   (replace `x` with your first name). The namespaces have already been created.


## Getting started

Use the following structure for your project:
- `cmd` folder that houses the file containing the `main` function.
- `internal` folder that houses all other files under appropriate packages.

Use [Go modules](https://blog.golang.org/using-go-modules) for your project.

Use [Go imports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) to format your code.

## Tips

This exercise is relatively larger than the previous ones and might require the full two weeks
of time that is available for it. So, make sure you start early.
