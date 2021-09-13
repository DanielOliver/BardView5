package models

import "github.com/cheekybits/genny/generic"

type ApiModel generic.Type


type ApiModelResponseType struct {
	Data ApiModel `json:"data"`
	Errors map[string]interface{} `json:"errors"`
}
