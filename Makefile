OUTPUT=output

.PHONY: clean
all: clean test build

clean:
	@echo -e "\nCLEANING $(OUTPUT) DIRECTORY"
	rm -rf ./$(OUTPUT)

build: clean
	@echo -e "\nBUILDING $(OUTPUT)/juggler BINARY" 
	mkdir -p ./$(OUTPUT) && go build -o $(OUTPUT)/juggler cmd/main.go

test:
	@echo -e "\nTESTING"
	go test -v ./... -coverprofile=coverage.out

run:
	go run cmd/main.go

install:
	cp $(OUTPUT)/juggler /usr/local/bin/