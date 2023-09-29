package characters

import "comics/type/json/common"

const Characters common.DataType = "Characters"

type CharactersArray []MainType

type CharactersResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
