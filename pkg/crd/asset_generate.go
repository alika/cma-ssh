// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/samsung-cnct/ims-kaas/pkg/crd"
)

func main() {
	if err := vfsgen.Generate(crd.Crd, vfsgen.Options{
		PackageName:  "crd",
		BuildTags:    "!dev",
		VariableName: "Crd",
	}); err != nil {
		log.Fatalln(err)
	}
}
