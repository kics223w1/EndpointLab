run: 
	go run main.go

build:
	go build -o endpointlab main.go

.PHONY: run build