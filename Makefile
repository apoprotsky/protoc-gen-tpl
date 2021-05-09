.PHONY: help all tidy build run clean

help:
	@echo "Usage:\n\
	  make       - print this message\n\
	  make help  - print this message\n\
	  make all   - update modules and build executable\n\
	  make tidy  - update modules\n\
	  make build - build executable\n\
	  make run   - generate code in examples\n\
	  make clean - remove built executable\n\
	"

all: tidy build

tidy:
	@go mod tidy

build:
	@go build -buildmode=default -ldflags="-s -w" -trimpath -o dist/protoc-gen-tpl cmd/protoc-gen-tpl/*.go

run:
	@rm -rf examples/out
	@mkdir examples/out
	@protoc \
		--tpl_out=examples/out \
		--tpl_opt=module=github.com/apoprotsky/prototpl/examples/ \
		examples/proto/*.proto \

clean:
	@rm -rf examples/out
	@rm dist/protoc-gen-tpl
