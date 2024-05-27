.PHONY: run
run:
	go run cmd/main.go cmd/routes.go


UNIT_TEST_PACKAGES=$(shell go list ./... | grep -vE 'cmd|database|mockgen|models|docs|repo|handlers')


.PHONY: test-long
test-long:
	go test $(UNIT_TEST_PACKAGES) -race -count=1 -parallel=8 -timeout=4m -coverprofile=coverage.txt
	go tool cover -html coverage.txt -o coverage.html
