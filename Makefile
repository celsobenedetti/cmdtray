PHONY=test clean all

build: 
	go build -o ~/.local/bin/cmdtray

run: build
	~/.local/bin/cmdtray

kill:
	pgrep cmdtray | xargs kill
