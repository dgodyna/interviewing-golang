
TARGET_DIR = $(CURDIR)/target
BUILD_DIR = $(TARGET_DIR)/build

$(TARGET_DIR):
	mkdir -p $(TARGET_DIR)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)


# Clean all the generated output
.PHONY: clean
clean:
	rm -fr $(TARGET_DIR)

# Builds generator and loader
.PHONY: build
build:
	go build -o $(BUILD_DIR)/bin/generator github.com/dmgo1014/interviewing-golang.git/cmd/generator
	go build -o $(BUILD_DIR)/bin/loader github.com/dmgo1014/interviewing-golang.git/cmd/loader

# stops pg instance
.PHONY: down_env
down_env:
	docker compose -f env/docker-compose.yaml down


# start pg instance
.PHONY: start_env
start_env:
	docker compose -f env/docker-compose.yaml up -d


# updates compose
.PHONY: update_env
update_env:
	docker compose -f env/docker-compose.yaml pull

# Test generate and load 10_000 of events
test_10k: clean build
	$(BUILD_DIR)/bin/generator 10000 test.json
	#$(BUILD_DIR)/bin/loader postgresql://test:test@localhost:5432/test?sslmode=disable test.json

# Test generate and load 100_000 of events
test_100k: clean build
	$(BUILD_DIR)/bin/generator 100000 test.json
	#$(BUILD_DIR)/bin/loader postgresql://test:test@localhost:5432/test?sslmode=disable test.json

# Test generate and load of 1_000_000 of events
test_1M: clean build
	$(BUILD_DIR)/bin/generator 1000000 test.json
	#$(BUILD_DIR)/bin/loader postgresql://test:test@localhost:5432/test?sslmode=disable test.json
# Test generate and load of 1_000_000 of events
test_1B: clean build
	$(BUILD_DIR)/bin/generator 1000000000 test.json
	#$(BUILD_DIR)/bin/loader postgresql://test:test@localhost:5432/test?sslmode=disable test.json
