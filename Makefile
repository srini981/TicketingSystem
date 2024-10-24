# Define the microservice directories
MICROSERVICES := chatService dbService notificationService  ticketService userService rbacService routingService

#build-all:
#		go build -o build/chatService ./chatService/cmd/main.go
#		go build -o build/dbService ./dbService/cmd/main.go
#		go build -o build/notificationService ./notificationService/cmd/main.go
#		go build -o build/rbacService ./rbacService/cmd/main.go
#		go build -o build/routingService ./routingService/cmd/main.go
#		go build -o build/ticketService ./ticketService/cmd/main.go
#		go build -o build/userService ./userService/cmd/main.go

# Build and run all microservices
.PHONY: build-all
build-all: $(MICROSERVICES)
	@for service in $(MICROSERVICES); do \
		echo "building $$service..."; \
		go build -o  build/$$service ./$$service/cmd/main.go & \
	done; \
	wait

# build a specific microservice
.PHONY: build
run:
	@read -p "Enter microservice name (e.g., service1): " service; \
	if [ -d "$$service" ]; then \
		echo "building $$service..."; \
		go build -o build/$$service ./$$service/cmd/main.go; \
	else \
		echo "Service $$service does not exist."; \
	fi


# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf ./build

# Run a specific microservice
.PHONY: run
run:
	@read -p "Enter microservice name (e.g., service1): " service; \
	if [ -d "$$service" ]; then \
		echo "Running $$service..."; \
		go run ./$$service/cmd/main.go; \
	else \
		echo "Service $$service does not exist."; \
	fi

# Build and run all microservices
.PHONY: run-all
run-all: $(MICROSERVICES)
	@for service in $(MICROSERVICES); do \
		echo "Running $$service..."; \
        cd ./$$service;\
		go run cmd/main.go & \
		cd ..;\
	done; \
	wait

up:
	docker-compose up
down:
	docker-compose down
# Print help message
.PHONY: help
help:
	@echo "Makefile for Go Microservices"
	@echo "Usage:"
	@echo "  make build-all        - Build all microservices"
	@echo "  make build            - Build a specfic microservices"
	@echo "  make clean            - Remove build artifacts"
	@echo "  make run-all     	   - run all microservices"
	@echo "  make up      		   - start the containers required for all microservices"
	@echo "  make down      	   - stop the containers required for all microservices"