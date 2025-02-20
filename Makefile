build:
	GOOS=linux GOARCH=amd64 go build -o target/logger-daemon server/server.go
	GOOS=linux GOARCH=amd64 go build -o target/client-app client/client.go

proto-compile:
	protoc --go_out=. --go-grpc_out=. proto/log.proto

test:
	go test ./server

docker-build:
	docker build -t log-img -f  docker/Dockerfile .

docker-run:
	docker run -d --name log-cnt log-img

docker-logview:
	docker exec -it log-cnt tail -f /var/log/syslog

docker-testlog1:
	docker exec -it log-cnt /usr/local/bin/app1 "This is a test log from app1"

docker-testlog2:
	docker exec -it log-cnt /usr/local/bin/app2 "This is a test log from app2"

docker-testlog3:
	docker exec -it log-cnt /usr/local/bin/app3 "This is a test log from app3"

clean:
	rm target/*
