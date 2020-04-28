package main

import (
	"os/exec"
)

const (
	_getBMGen = "go get -u github.com/gs02048/micro/tool/protobuf/protoc-gen-bm"
	_bmProtoc = "protoc --bm_out=:."
)

func installBMGen() error {
	if _, err := exec.LookPath("protoc-gen-bm"); err != nil {
		if err := goget(_getBMGen); err != nil {
			return err
		}
	}
	return nil
}

func genBM(files []string) error {
	return generate(_bmProtoc, files)
}