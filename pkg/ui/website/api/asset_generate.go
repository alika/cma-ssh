// +build ignore

package main

import (
	"github.com/samsung-cnct/ims-kaas/pkg/ui/website/api"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	if err := vfsgen.Generate(api.Api, vfsgen.Options{
		PackageName:  "api",
		BuildTags:    "!dev",
		VariableName: "Api",
	}); err != nil {
		log.Fatalln(err)
	}

}
