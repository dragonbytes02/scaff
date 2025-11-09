##build: build the aplication 
.PHONY: build
build: 
	go mod tidy
	if not exist bin mkdir bin
	go build -o ./bin/scaff.exe .

##run: execute bin with create command
.PHONY: run
run:
	./bin/scaff.exe create project

##brun: build and run
.PHONY: brun
brun: build run