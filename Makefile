
TARGET_DIR = $(CURDIR)/target
BUILD_DIR = $(TARGET_DIR)/build

$(TARGET_DIR):
	mkdir -p $(TARGET_DIR)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)


.PHONY: clean
clean:
	rm -fr $(TARGET_DIR)

.PHONY: build
build:
	go build -o $(BUILD_DIR)/bin/generator github.com/dmgo1014/interviewing-golang.git/cmd/generator
	go build -o $(BUILD_DIR)/bin/loader github.com/dmgo1014/interviewing-golang.git/cmd/loader

.PHONY: down_env
down_env:
	docker-compose -f env/docker-compose.yaml down


.PHONY: start_env
start_env:
	docker-compose -f env/docker-compose.yaml up -d


.PHONY: update_env
update_env:
	docker-compose -f env/docker-compose.yaml pull

test_10k: clean build
	$(BUILD_DIR)/bin/generator 10000 test.json
	$(BUILD_DIR)/bin/loader postgresql://test:test@localhost:5432/test?sslmode=disable test.json

