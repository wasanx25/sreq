deps:
		go mod vendor

install: deps
		go install

lint:
		@go get -v github.com/golang/lint/golint
		golint -min_confidence=1.0 $(shell glide nv)
