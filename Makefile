.PHONY: spec clean

spec:
	protoc -I=./spec/ --go_out=plugins=grpc:./ ./spec/*.proto

clean:
	rm -rf ./*.go

