SOCK_PATH?=driver.sock
TRACEFILE?=

.PHONY: help
help:
	@echo "Usage: make [run.sock | run.stdio | run.yolo]"
	
# Start server and listen UNIX socket
.PHONY: run.sock
run.sock:
	@go run ./cmd/gopackagesdriver -serve -sock "$(SOCK_PATH)" -trace "$(TRACEFILE)"

# Start server and listen for stdio requests
.PHONY: run.stdio
run.stdio:
	@go run ./cmd/gopackagesdriver -serve -trace "$(TRACEFILE)"

# Run as regular go packages driver program that reads arguments from command line and immediately returns result to stdout.
.PHONY: run.yolo
run.yolo:
	@go run ./cmd/gopackagesdriver

