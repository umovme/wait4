tools:
	go get -u -v github.com/goreleaser/goreleaser
	go get -u -v github.com/golang/dep/cmd/dep
dependencies: 
	dep ensure
deploy: dependencies
	goreleaser --rm-dist