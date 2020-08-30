.PHONY: vendor
vendor:
	go mod vendor && go mod tidy
