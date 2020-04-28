package main

import (
	"os/exec"
)

const (
	_getSwaggerGen = "go get -u github.com/gs02048/micro/tool/protobuf/protoc-gen-bswagger"
	_swaggerProtoc = "protoc --bswagger_out=:."
)

func installSwaggerGen() error {
	if _, err := exec.LookPath("protoc-gen-bswagger"); err != nil {
		if err := goget(_getSwaggerGen); err != nil {
			return err
		}
	}
	return nil
}

func genSwagger(files []string) error {
	return generate(_swaggerProtoc, files)
}
