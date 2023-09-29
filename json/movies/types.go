package movies

import "comics/type/json/common"

const Movies common.DataType = "Movies"

type MovieResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
