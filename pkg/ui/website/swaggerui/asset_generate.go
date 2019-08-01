// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/samsung-cnct/ims-kaas/pkg/ui/website/swaggerui"
)

func main() {
	if err := vfsgen.Generate(swaggerui.SwaggerUI, vfsgen.Options{
		PackageName:  "swaggerui",
		BuildTags:    "!dev",
		VariableName: "SwaggerUI",
	}); err != nil {
		log.Fatalln(err)
	}
}
