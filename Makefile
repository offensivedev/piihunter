install:
	go install

test:
	cd cmd && go test -v

lint:
	golint . cmd

format:
	go fmt
	cd cmd && go fmt