package characters

import "github.com/nanikjava/comicstype/json/common"

const Characters common.DataType = "Characters"

type CharactersArray []MainType

type CharactersResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
