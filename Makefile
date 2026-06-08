setup:
	bash setup.sh

build:
	go build -o container .

run:
	sudo ./container run /bin/sh

all: setup build