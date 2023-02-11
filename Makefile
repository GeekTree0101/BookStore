build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-w -s" -o bookStore main.go