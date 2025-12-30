.PHONY: help build run clean

help:
	@echo "Available commands:"
	@echo "  make build  - Build the Docker image"
	@echo "  make run    - Run the container on port 3333"
	@echo "  make clean  - Remove the Docker image"
	@echo "  make help   - Display this help message"

build:
	docker build -t echo:latest .

run:
	docker run --rm -p 3333:3333 echo:latest

clean:
	docker rmi echo:latest