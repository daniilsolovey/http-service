# http-service

**http-service** is a tool for receiving and updating users


## Installation

```
go get github.com/daniilsolovey/http-service
```


## Usage

##### -c --config \<path>
Read specified config file. [default: config.toml].

##### --debug
Enable debug messages.

##### -v --version
Print version.

#####  -h --help
Show this help.

## Build
For build program use command:

```
make build
```

## RUN
For running program use command:

```
./http-service
```

## RUN TESTS

```
go test ./...
```


## COMMANDS
Request for updating fields of user:

```
curl -X PUT -H "Content-Type: application/json" -d "{\"age\":newAgeInt}" http://localhost:8080/users/{id}
curl -X PUT -H "Content-Type: application/json" -d "{\"name\":"newNameString"}" http://localhost:8080/users/{id}
```

For receiving user by id use command:

```
curl -X GET http://localhost:8080/users/{id}
```

For receiving all users use command:

```
curl -X GET http://localhost:8080/users
```
