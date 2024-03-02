tests:
	go test -v ./internal/tests/*

migration:
	go run cmd/migration/main.go