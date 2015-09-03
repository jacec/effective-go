default: clean build

build: check
	$(GOPATH)/bin/godep restore
	go build -o ./bin/effective-go

clean:
	rm -f ./bin/*

check:
	./check-dep.sh
