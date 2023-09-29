package contract

import "github.com/nanikjava/comicstype/json/common"

type InformationCaller[T any, V any] interface {
	GetData(apiUrl string) error
	GetDataType() common.DataType
	Get(idx int) *V
	Len() int
}
