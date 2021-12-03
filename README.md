# Capstone Project

API made in GO using clean architecture, this project is meant as an exercise to learn the basics of the language.

## Getting started

Clone this repo in your preferred location

``` bash
git clone https://github.com/FernandoGal25/academy-go-q42021
```

### Prerequisites

- `Go` installed in your machine.

- Rename `config.yml.dist` to `config.yml` in `config` directory and replace `default values` with the required values.

## Running the project

- Get/update dependencies by running:

``` bash
go mod tidy
```

- Run the project:

``` bash
go run main.go
```

## Basic endpoints

`hostname` and `port` variables are the ones inside `config.yml` file, under `server` field.

- Get all pokemons in store:

``` bash
GET <hostname>:<port>/pokemons
```

- Get a single Pokemon by it's ID:

``` bash
GET <hostname>:<port>/pokemons/{id}
```

## TODOs

## Contact

Fernando Galindo Arroyo - fernando.gali25@gmail.com