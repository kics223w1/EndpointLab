# Variables
DOCKER_REPO=viethuy/endpointlab
PLATFORM=linux/amd64

# Run the application
run: 
	go run main.go

# Build the application
build:
	go build -o endpointlab main.go

# Update the swagger docs
swagger:
	swag init -g api/server.go

# Build the docker image (single platform)
dockerBuild:
	docker build -t $(DOCKER_REPO):latest .

# Push the docker image
dockerPush:
	docker push $(DOCKER_REPO):latest

dockerRemoveImage:
	-docker rmi $(DOCKER_REPO):latest

# Combined remove, build and push
dockerDistribute: dockerRemoveImage dockerBuild dockerPush

.PHONY: run build swagger dockerBuild dockerPush dockerRemoveImage dockerDistribute
