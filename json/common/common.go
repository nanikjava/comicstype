package common

import client "github.com/nanikjava/comicstype/http"

type DataType string

type CommonStruct struct {
	Offset int
	Limit  int
	*client.HttpClient
}

const (
	COMICVIEW_BASEURL = "https://comicvine.gamespot.com/api/"
)
