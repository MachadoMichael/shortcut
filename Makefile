build:
	go build -o ./bin/shortcut

alias:
	chmod +x ./set_alias.sh
	./set_alias.sh

all: build alias
