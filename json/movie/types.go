package movie

import "github.com/nanikjava/comicstype/json/common"

const Movie common.DataType = "Movie"

type MovieResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
