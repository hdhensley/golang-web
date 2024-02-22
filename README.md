# Golang Web

## Description

This is a project to get started with Go web applications using HTMX, Tailwind, 
and GORM. Hot reloading is built into the tooling and handled with a Makefile 
for the frontend changes and air for the backend

## Getting Started

### Prerequisites
*See .tool-versions for asdf management of tool versions*
* Docker & Docker Compose - *Optional* for running the application as a container
* Go 1.21.1 - .asdf management recommended
* Node 18.16.0 - .asdf management recommended
* MySQL 8.0 - When using docker-compose, this is included
* Templ [Install](https://templ.guide/quick-start/installation) - `go install github.com/a-h/templ/cmd/templ@latest`
* Air [Install](https://github.com/cosmtrek/air) - `go install github.com/cosmtrek/air@latest`

### Installation

clone and enter the repository root

``` 
asdf install golang 1.21.1
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
go mod tidy
asdf install nodejs 18.16.0
npm i
cp .env.sample .env (update db credentials now if needed)
```

### Running the application

Once the dependencies are installed, run the following commands to start the server and frontend

#### Locally with hot swapping

```
make watchTempl
make devFrontend
make devServer 
```

#### Docker Compose (No hot swapping yet)

```
make runDocker
```

## Building the application

### Docker
```
make buildDockerImage
```

### Build locally
```
make build