unit_tests:
	go test -v ./internal/tests/unit_tests/*

migration:
	go run cmd/migration/main.go

run:
	go run cmd/transaction-system/main.go

integration_tests:
	go test -v ./internal/tests/integration_tests/*