PHONY: all
all:
	go build

armv6:
	CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=6 go build

install:
	go install

clean:
	rm icka
