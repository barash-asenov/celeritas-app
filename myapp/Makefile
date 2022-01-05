BINARY_NAME=celeritasApp

build:
	@#go mod vendor
	@echo "Building celeritas"
	@go build -o tmp/${BINARY_NAME} .
	@echo "Celeritas built!"

run:  build
	@echo "Starting Celeritas..."
	@./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test: 
	@echo "Testing..."
	@go test ./...
	@ecoh "Done!"

start: run

stop:
	@echo "Stopping celeritas..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start