make

go test -i ./command || exit 1
go test -v ./command -race || exit 1
