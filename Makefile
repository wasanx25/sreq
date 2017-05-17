glide:
		go get github.com/Masterminds/glide

deps: glide
		glide install

update: glide
		glide update

install: main.go deps
		go install
