package main

import (
	"os/exec"
)

const (
	_getEcodeGen = "go get -u github.com/gs02048/micro/tool/protobuf/protoc-gen-ecode"
	_ecodeProtoc = "protoc --ecode_out=:."
)

func installEcodeGen() error {
	if _, err := exec.LookPath("protoc-gen-ecode"); err != nil {
		if err := goget(_getEcodeGen); err != nil {
			return err
		}
	}
	return nil
}

func genEcode(files []string) error {
	return generate(_ecodeProtoc, files)
}
