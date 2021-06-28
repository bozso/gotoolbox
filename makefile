BIN ?= /tmp/gotoolbox
DEPLOY_TARGET ?= /home/istvan/mount/eejit/packages/usr/bin

.PHONY: $(BIN)


test:
	go test ./hash

$(BIN):
	go build -o $(BIN)

deploy: $(BIN)
	cp $(BIN) $(DEPLOY_TARGET)

install: $(BIN)
	cp $(BIN) /home/istvan/packages/usr/bin
