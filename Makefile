all : parser.go

parser.go: parser.go.y
	goyacc -o parser.go parser.go.y
	gofmt -s -w .

test : parser.go
	go test ./...