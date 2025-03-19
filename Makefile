.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook ./cmd/webook/main.go
	@docker rmi -f zhubao/webook:v0.0.1
	@docker build -t zhubao/webook:v0.0.1 .