.PHONY: spec clean

spec:
	cd _specfiles && protoc -I=./service/ --go_out=plugins=grpc:../service ./service/*.proto
	cd _specfiles && protoc -I=./sync/ --go_out=plugins=grpc:../sync ./sync/*.proto

clean:
	rm -rf ./*.go

