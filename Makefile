
source_path := .
output_path := $(shell go env GOPATH)/src

generate-pediction-go:
	protoc -I/usr/local/include/google/protobuf -I$(source_path) --go_out=$(output_path) prediction/v1alpha1/*.proto 