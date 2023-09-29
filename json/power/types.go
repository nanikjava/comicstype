package power

import "comics/type/json/common"

const Power common.DataType = "Power"

type PowerResponseData struct {
	DataType common.DataType `json:"datatype"`
	RawData  *MainType       `json:"rawdata"`
}
