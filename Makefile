TEMP_DIR=./tmp
TEMP_FILE=file.log
DEMO_FILE=file2.log
BUILD_DIR=./dist
BIN_NAME=inotify-playground
BIN_PATH=$(BUILD_DIR)/$(BIN_NAME)

clean-build:
	rm -rf ./dist/

clean-tmp:
	rm -rf ./tmp/

clean: clean-build clean-tmp

build: clean
	@mkdir $(BUILD_DIR)
	go build -o $(BIN_PATH) *.go

setup: clean
	@mkdir $(TEMP_DIR)
	@touch $(TEMP_DIR)/$(TEMP_FILE)
	@echo "some log" > $(TEMP_DIR)/$(TEMP_FILE)

run: clean setup build 
	exec $(BIN_PATH)

run-demo:
	@echo "some dumb text" >> tmp/file.log
	@touch $(TEMP_DIR)/$(DEMO_FILE)
