.PHONY: all linux macos windows clean

TARGET = gitmon
SRC_DIR = cmd
SRC_FILE = main.go

all: linux windows

linux:
	GOOS=linux GOARCH=amd64 go build -o $(TARGET) $(SRC_DIR)/$(SRC_FILE)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(TARGET).exe $(SRC_DIR)/$(SRC_FILE)

clean:
	rm -f $(TARGET) $(TARGET).exe
