package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gs02048/micro/tool/protobuf/pkg/gen"
	"github.com/gs02048/micro/tool/protobuf/pkg/generator"
	ecodegen "github.com/gs02048/micro/tool/protobuf/protoc-gen-ecode/generator"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(generator.Version)
		os.Exit(0)
	}

	g := ecodegen.EcodeGenerator()
	gen.Main(g)
}
