all:
	CGO_ENABLED=0 go build -ldflags '-w' -o bin/kube-scheduler cmd/main.go \
	&& sudo docker build . -t danielbgmitre/reverse-scheduler \
	&& sudo docker push danielbgmitre/reverse-scheduler \
	&& sudo docker pull danielbgmitre/reverse-scheduler