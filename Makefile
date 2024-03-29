OUTPUT=output

.PHONY: clean all run install 
all: clean test build

clean:
	@echo -e "\nCLEANING $(OUTPUT) DIRECTORY"
	rm -rf ./$(OUTPUT)

build: clean test
	@echo -e "\nBUILDING $(OUTPUT)/juggler BINARY" 
	mkdir -p ./$(OUTPUT) && go build -o $(OUTPUT)/juggler cmd/main.go

test:
	@echo -e "\nTESTING"
	go test -v ./... -coverprofile=coverage.out

run:
	go run cmd/main.go

install: build
	sudo cp $(OUTPUT)/juggler /usr/local/bin/