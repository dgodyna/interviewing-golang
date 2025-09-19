
TARGET_DIR = $(CURDIR)/target
BUILD_DIR = $(TARGET_DIR)/build

$(TARGET_DIR):
	mkdir -p $(TARGET_DIR)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)


# Clean all generated build artifacts and output files
.PHONY: clean
clean:
	rm -fr $(TARGET_DIR)

# Build the network event data generator binary
.PHONY: build
build:
	go build -o $(BUILD_DIR)/bin/generator github.com/dmgo1014/interviewing-golang.git/cmd/generator

# Generate 10,000 events - Quick validation test (~75ms baseline)
test_10k: clean build
	$(BUILD_DIR)/bin/generator 10000 test.json

# Generate 100,000 events - Medium scale test (~529ms baseline)
test_100k: clean build
	$(BUILD_DIR)/bin/generator 100000 test.json

# Generate 1,000,000 events - Large scale validation (~5s baseline)
test_1M: clean build
	$(BUILD_DIR)/bin/generator 1000000 test.json

# Generate 10,000,000 events - Performance stress test (~65s baseline)
test_10M: clean build
	$(BUILD_DIR)/bin/generator 10000000 test.json
