build:
	go build dup
install: build
	sudo mv ./dup /usr/local/bin