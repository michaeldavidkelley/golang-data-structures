all:
	@make test

test: 
	@go test ./...

watch:
	@fswatch -o . | xargs -n1 -I{} make test
