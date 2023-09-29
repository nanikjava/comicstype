package publisher

import "github.com/nanikjava/comicstype/json/common"

const Publisher common.DataType = "Publisher"

type PublisherResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
