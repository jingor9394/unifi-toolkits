PATH_CMD_PRUNE_CLIENTS=./cmd/prune-clients
PATH_CMD_PRINT_MAC_FILTER=./cmd/print-mac-filter
BINARY_NAME_PRUNE_CLIENTS=prune-clients
BINARY_NAME_PRINT_MAC_FILTER=print-mac-filter

build:
	cd ${PATH_CMD_PRUNE_CLIENTS} && go build -o ${BINARY_NAME_PRUNE_CLIENTS} *.go && mv ${BINARY_NAME_PRUNE_CLIENTS} ../../
	cd ${PATH_CMD_PRINT_MAC_FILTER} && go build -o ${BINARY_NAME_PRINT_MAC_FILTER} *.go && mv ${BINARY_NAME_PRINT_MAC_FILTER} ../../

release:
	zip ${BINARY_NAME_PRUNE_CLIENTS}-darwin-${VERSION}.zip ./${BINARY_NAME_PRUNE_CLIENTS}
	zip ${BINARY_NAME_PRINT_MAC_FILTER}-darwin-${VERSION}.zip ./${BINARY_NAME_PRINT_MAC_FILTER}

run_prune:
	cd ${PATH_CMD_PRUNE_CLIENTS} && go run *.go

run_filter:
	cd ${PATH_CMD_PRINT_MAC_FILTER} && go run *.go

format:
	gofmt -s -w .

clean:
	rm -f ${BINARY_NAME_PRUNE_CLIENTS}
	rm -f ${BINARY_NAME_PRINT_MAC_FILTER}
	rm -f ${BINARY_NAME_PRUNE_CLIENTS}-darwin-*.zip
	rm -f ${BINARY_NAME_PRINT_MAC_FILTER}-darwin-*.zip
