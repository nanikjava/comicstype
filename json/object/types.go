package object

import "github.com/nanikjava/comicstype/json/common"

const Object common.DataType = "Object"

type ObjectResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
