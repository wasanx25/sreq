glide:
		go get github.com/Masterminds/glide

deps: glide
		glide install

update: glide
		glide update

install: main.go deps
		go install

lint: 
		@go get -v github.com/golang/lint/golint
		golint -min_confidence=1.0 $(shell glide nv)
