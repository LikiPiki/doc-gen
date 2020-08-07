default:
	go build ./...

install:
	go install ./...

clean:
	rm -rf doc-gen