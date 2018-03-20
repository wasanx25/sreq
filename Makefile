godep:
		go get -u github.com/golang/dep/cmd/dep

deps: godep
		dep ensure

update: godep
		dep ensure -update

install: main.go deps
		go install

lint:
		@go get -v github.com/golang/lint/golint
		golint -min_confidence=1.0 $(shell glide nv)
