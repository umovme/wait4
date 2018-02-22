tools:
	go get -u -v github.com/goreleaser/goreleaser
deploy:
	goreleaser --rm-dist