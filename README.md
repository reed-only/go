# Go API Server

## Install go

```
brew install golang
```

## Fetch Dependencies

```bash
go get github.com/gorilla/mux
```

## Run API Server

```bash
go run *.go
```

## Try It Out

Visit [localhost:8080/cars](http://localhost:8080/cars)

### Get All Cars

```bash
curl http://localhost:8080/cars
```

### Get a Single Car

```bash
curl http://localhost:8080/cars/1
```

### Create a Car (Not supported yet)

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 3, "make": "Alfa Romeo", "model": "8C", "year": 2017, "transmission": {"type": "Dual Dry Clutch", "gears": 6}}' http://localhost:8080/cars
```

### Delete a Car

```bash
curl -X DELETE http://localhost:8080/cars/1
```