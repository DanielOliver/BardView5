// +build tools

package tools

//cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %
import (
	_ "github.com/OneOfOne/struct2ts"
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	//_ "github.com/jteeuwen/go-bindata"
	_ "github.com/kevinburke/go-bindata/go-bindata"
)

