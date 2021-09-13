package main

import "github.com/cheekybits/genny/generic"

type ApiModel generic.Type

type ApiModelResponseType struct {
	Data ApiModel
}
