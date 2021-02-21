OUTPUT=build

.PHONY: clean
all: clean test build

clean:
	@echo -e "\nCLEANING $(OUTPUT) DIRECTORY"
	rm -rf ./$(OUTPUT)

build: clean
	@echo -e "\nBUILDING $(OUTPUT)/juggler BINARY" 
	mkdir -p ./$(OUTPUT) && go build -o $(OUTPUT)/juggler

test:
	@echo -e "\nTESTING"
	go test -v ./...

run:
	go run main.go