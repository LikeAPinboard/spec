.PHONY: spec clean

spec:
	protoc -I=./spec/ --go_out=plugins=grpc:./entity/ ./spec/*.proto

clean:
	rm -rf ./entity/*

