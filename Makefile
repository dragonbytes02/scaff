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

#delete project create with run
.PHONY: clean
clean: 
	rmdir /S /Q project


##brun: build, run and clean
.PHONY: dev
dev: build run clean
