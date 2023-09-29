package publisher

import "comics/type/json/common"

const Publisher common.DataType = "Publisher"

type PublisherResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
