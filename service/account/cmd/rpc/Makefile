download:
	@go mod tidy

build:download
	@go build -o cmd

run:download
	@ulimit -HSn 10000
	@go run account.go -f etc/account.yaml

rpc:
	@goctl rpc proto -src account.proto -dir .

model:clean_model
	@goctl model mysql ddl -src ./*.sql -dir ./model . -c 

grpcui:
	@grpcui -plaintext localhost:8090

clean_model:
	@rm -rf ./model/*