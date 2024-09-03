
run:
	go run cmd/main.go

swag-n-go:
	swag init -g cmd/main.go
	go run cmd/main.go

test:
	go test mf-loan/delivery/http/tests mf-loan/repository/tests mf-loan/usecase/tests -v

setup-swag:
	swag init -g cmd/main.go

setup-test:
	mockery --recursive --output=repository/mocks --outpkg=mocks --name=CustomerRepository
	mockery --recursive --output=repository/mocks --outpkg=mocks --name=CustomerUseCase

