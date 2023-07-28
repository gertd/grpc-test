SHELL 	   := $(shell which bash)

NO_COLOR   :=\033[0m
OK_COLOR   :=\033[32;01m
ERR_COLOR  :=\033[31;01m
WARN_COLOR :=\033[36;01m
ATTN_COLOR :=\033[33;01m


.PHONY: all
all: build

build:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	buf build proto
	buf generate proto

.PHONY: run-server
run-server:
	@echo -e "$(ATTN_COLOR)==> $@ $(NO_COLOR)"
	@go run server/server.go

.PHONY: run-client-success
run-client-success:
	@grpcurl -plaintext -d '{"name": "testing 123"}' localhost:50051 helloworld.Greeter.SayHello

.PHONY: run-client-failed
run-client-failed:
	@grpcurl -plaintext -d '{"name": "testing 123"}' localhost:50051 helloworld.Greeter.SayHallo
