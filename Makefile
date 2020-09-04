default:
	go build ./...

install:
	go install ./...

croscompile:
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...

clean:
	rm -rf doc-gen