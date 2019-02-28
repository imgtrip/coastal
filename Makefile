# note: call scripts from /scripts

.PHONY: all dep test

test:
	@cd test/coastal && bash test.sh

dep: ## Get the dependencies^
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure
