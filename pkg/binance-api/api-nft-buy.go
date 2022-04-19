package bapi

import (
	"fmt"
	"encoding/json"
	binance_struct "github.com/1makarov/binance-nft-buy/internal/domain/binance-api"
	"github.com/valyala/fasthttp"
)

const (
	orderCreateUrl = "https://www.binance.com/bapi/nft/v1/private/nft/nft-trade/order-create"
)

const (
	orderQueryUrl = "https://www.binance.com/bapi/nft/v1/private/nft/nft-trade/order-query"
)

func (api *Api) orderCreateRequest(body []byte) *fasthttp.Request {
	r := fasthttp.AcquireRequest()
	api.request.CopyTo(r)
	r.Header.SetMethod(fasthttp.MethodPost)
	r.Header.SetContentType("application/json")
	r.Header.SetRequestURI(orderCreateUrl)
	r.SetBody(body)
	fmt.Println(r)
	return r
}

func (api *Api) orderQueryRequest(body []byte) *fasthttp.Request {
	r := fasthttp.AcquireRequest()
	api.request.CopyTo(r)
	r.Header.SetMethod(fasthttp.MethodPost)
	r.Header.SetContentType("application/json")
	r.Header.SetRequestURI(orderQueryUrl)
	r.SetBody(body)
	fmt.Println(r)
	return r
}

func (api *Api) NFTMysteryBoxBuy(req *fasthttp.Request) (*fasthttp.Response, error) {
	response, err := api.postRequest(req)
	if err = handleError(response, err); err != nil {
		return nil, err
	}
	return response, nil
}

func MarshalMysteryBoxBuy(productID string, amount int) ([]byte, error) {
	b, err := json.Marshal(binance_struct.BuyRequest{ProductID: productID, Amount: amount})
	if err != nil {
		return nil, err
	}
	return b, nil
}
