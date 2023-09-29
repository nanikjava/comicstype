package people

import "comics/type/json/common"

const Movie common.DataType = "Movie"

type MovieResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
