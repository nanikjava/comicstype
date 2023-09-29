package object

import "comics/type/json/common"

const Object common.DataType = "Object"

type ObjectResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
