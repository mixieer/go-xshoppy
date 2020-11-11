package xssdk

import "fmt"

type ShopService interface {
	Check(CheckOption) (*CommonResponse, error)
}

type ShopServiceOp struct {
	XsClient *XsApiClient
}

type CheckOption struct {
	ShopUrl string //店铺后台地址
}

func (s *ShopServiceOp) Check(checkOption CheckOption) (*CommonResponse, error) {
	url := fmt.Sprintf("%v%v", HOST, "/sail/user/verify")
	resource := &CommonResponse{}
	customHeader := make(map[string]string)
	customHeader["DOMAIN"] = checkOption.ShopUrl
	err := s.XsClient.Post(url, nil, resource, customHeader)
	return resource, err
}
