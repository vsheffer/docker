all: mongoconf

pkg: ./pkg/mongo.go
	go install ./pkg

mongoconf: mongoconf.go pkg
	CGO_ENABLED=0 go build -a --ldflags '-w' ./mongoconf.go

container: mongoconf
	sudo docker build -t vsheffer/mongo:mongoconf .

push:
	sudo docker push vsheffer/mongo:mongoconf

clean:
	touch ./pkg/mongo.go
	rm -f mongoconf
