package character

import "github.com/nanikjava/comicstype/json/common"

const Character common.DataType = "Character"

type CharacterArray []*MainType

type CharacterResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
