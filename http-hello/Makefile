.PHONY: module
module: http-hello
	echo '/http-hello: $${MODULE_DIR}/http-hello' > usr.manifest

http-hello: clean
	go build -buildmode=c-shared

clean:
	rm -f http-hello usr.manifest
