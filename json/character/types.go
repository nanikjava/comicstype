package character

import "comics/type/json/common"

const Character common.DataType = "Character"

type CharacterArray []*MainType

type CharacterResponseData struct {
	DataType common.DataType `json:"datatype"`
	Data     *Results        `json:"data"`
}
