# Makefile for building and running the Docker image and container
.PHONY: run build stop cleanI cleanC exec
.DEFAULT_GOAL:= run

PROJECT_NAME = multi-wordle

# Build the Docker image and run the container
run: cleanC build
	docker run -d --restart=always -p 3000:3000 --name=$(PROJECT_NAME) $(PROJECT_NAME)

# Build the Docker image
build:
	docker build -t $(PROJECT_NAME):latest .

# Stop and remove the Docker container
stop:
	docker stop $(PROJECT_NAME)
	docker rm $(PROJECT_NAME)

# Clean up the Docker image
cleanI:
	docker rmi $(PROJECT_NAME)
	docker builder prune --filter="image=$(PROJECT_NAME)"

# Clean up the Docker container
cleanC:
	docker stop $(PROJECT_NAME)
	docker rm $(PROJECT_NAME)

# Run the application inside the Docker container
exec:
	docker exec -it $(PROJECT_NAME) /app/tarsCron

live:
#	find . -type f \( -name '*.go' -o -name '*.gohtml' \) | entr -r sh -c 'make && docker logs --follow $(APP_NAME)'
	find . -type f \( -name '*.go' -o -name '*.html' \) | entr -r sh -c 'go build -o /tmp/build ./cmd && /tmp/build'