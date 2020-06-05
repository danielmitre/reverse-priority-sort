all: build
	sudo docker push danielbgmitre/reverse-scheduler
	sudo docker pull danielbgmitre/reverse-scheduler

build: compile
	sudo docker build . -t danielbgmitre/reverse-scheduler

compile: format bin/kube-scheduler
	CGO_ENABLED=0 go build -ldflags '-w' -o bin/kube-scheduler cmd/main.go

format:
	go fmt cmd/main.go

