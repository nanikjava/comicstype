package client

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/nanikjava/comicstype/json/common"
	"log"
	"time"
)

type HttpClient struct {
	Client *req.Client
}

func New(apikey string) *HttpClient {
	client := req.C().
		SetCommonQueryParam("api_key", apikey).
		SetTimeout(5 * time.Second)

	return &HttpClient{Client: client}
}

func (h *HttpClient) Call(queryMap map[string]string, resultType interface{}, urlPath string) interface{} {
	resp, err := h.Client.R().
		SetQueryParams(queryMap).
		SetSuccessResult(resultType).
		EnableDump().
		Get(fmt.Sprintf("%s%s", common.COMICVIEW_BASEURL, urlPath))

	if err != nil {
		log.Printf("err: ", err)
		return nil
	}

	return resp.Request.Result
}

func (h *HttpClient) CallSingle(queryMap map[string]string, resultType interface{}, urlPath string) interface{} {
	resp, err := h.Client.R().
		SetQueryParams(queryMap).
		SetSuccessResult(resultType).
		EnableDump().
		Get(fmt.Sprintf("%s", urlPath))

	if err != nil {
		log.Printf("err: ", err)
		return nil
	}

	return resp.Request.Result
}
