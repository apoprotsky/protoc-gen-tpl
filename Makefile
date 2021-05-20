.PHONY: help all tidy build link run clean

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
	@go build -buildmode=default -ldflags="-s -w" -trimpath .

link:
	@ln -s $$(pwd)/protoc-gen-tpl /usr/local/bin/

run:
	@rm -rf examples/out
	@mkdir examples/out
	@protoc \
		--proto_path=examples/proto \
		--tpl_out=examples/out \
		--tpl_opt=prefix=github.com/apoprotsky/protoc-gen-tpl/examples/ \
		--tpl_opt=lang=go \
		--tpl_opt=lang=ts \
		--tpl_opt=lang=php \
		examples/proto/*.proto \

clean:
	@rm -rf examples/out
	@rm protoc-gen-tpl
