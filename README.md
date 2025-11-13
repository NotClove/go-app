# Go Test Application

Simple Go application for CI/CD testing with Jenkins, SonarQube, and Nexus.

## Features

- Go 1.21
- Unit tests
- SonarQube code analysis
- Docker containerization
- Nexus repository integration

## Requirements

- Go 1.21+

## Build

```bash
go build -o app .
```

## Test

```bash
go test -v ./...
```

## Run

```bash
./app
```

## Docker

```bash
docker build -t go-app:latest .
docker run go-app:latest
```

