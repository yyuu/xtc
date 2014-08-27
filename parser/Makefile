.PHONY: all
.SUFFIXES: .go .go.y
all: parser.go
.go.y.go:
	go tool yacc -o $@ $<
