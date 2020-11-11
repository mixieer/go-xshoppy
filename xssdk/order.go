package xssdk

import (
	"fmt"
	"strconv"
)

type OrderService interface {
	List(listOptions interface{}) ([]Order, error)
	Update(updateOptions interface{}) (*CommonResponse, error)
	Count(CountOptions interface{}) (int, error)
}

type OrderServiceOp struct {
	XsClient *XsApiClient
}

type ListOptions struct {
	Limit           string `url:"limit"`
	Page            string `url:"page"`
	Filter          string `url:"filter,omitempty"`
	TimeStart       string `url:"time_start,omitempty"`
	TimeEnd         string `url:"time_end,omitempty"`
	Status          string `url:"status,omitempty"`
	FinancialStatus string `url:"financial_status,omitempty"`
	UpdateTimeStart string `url:"update_time_start,omitempty"`
	UpdateTimeEnd   string `url:"update_time_end,omitempty"`
	PayTimeStart    string `url:"pay_time_start,omitempty"`
	PayTimeEnd      string `url:"pay_time_end,omitempty"`
}

type CountOptions struct {
	Filter          string `url:"filter,omitempty"`
	TimeStart       string `url:"time_start,omitempty"`
	TimeEnd         string `url:"time_end,omitempty"`
	Status          string `url:"status,omitempty"`
	FinancialStatus string `url:"financial_status,omitempty"`
	UpdateTimeStart string `url:"update_time_start,omitempty"`
	UpdateTimeEnd   string `url:"update_time_end,omitempty"`
	PayTimeStart    string `url:"pay_time_start,omitempty"`
	PayTimeEnd      string `url:"pay_time_end,omitempty"`
}

type CountResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data CountResponseData `json:"data"`
	Now  string            `json:"now"`
}

type CountResponseData struct {
	Count string `json:"count"`
}

type UpdateOptions struct {
	Id                  string          `json:"id"`                              //订单id
	Email               string          `json:"email,omitempty"`                 //用户联系邮箱
	TrackingNumber      string          `json:"tracking_number,omitempty"`       //物流单号
	FulfillmentStatus   string          `json:"fulfillment_status,omitempty"`    //物流状态 pending 未发货 fulfilled 已发货 abandoned 废弃的
	FulfillmentSendTime string          `json:"fulfillment_send_time,omitempty"` //物流发货时间
	ShippingAddress     ShippingAddress `json:"shipping_address,omitempty"`      //物流地址
}

type ShippingAddress struct {
	Phone    string `json:"phone,omitempty"`    //用户物流联系手机号
	Address1 string `json:"address1,omitempty"` //用户物流地址1
	Address2 string `json:"address2,omitempty"` //用户物流地址2
	City     string `json:"city,omitempty"`     //用户物流所在市
	Zip      string `json:"zip,omitempty"`      //用户物流所在地邮编
	Province string `json:"province,omitempty"` //用户物流所在州
}

type OrdersResource struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data ListData `json:"data"`
	Now  string   `json:"now"`
}

type ListData struct {
	TotalCount  string  `json:"totalCount"`
	CurrentPage string  `json:"currentPage"`
	Limit       string  `json:"limit"`
	Data        []Order `json:"data"`
}

type OrderResource struct {
	Order *Order `json:"order"`
}

func (xs *OrderServiceOp) List(options interface{}) ([]Order, error) {

	url := fmt.Sprintf("%v%v", HOST, "/order/orders/list")
	resource := &OrdersResource{}
	err := xs.XsClient.Get(url, resource, options)
	return resource.Data.Data, err
}

func (xs *OrderServiceOp) Update(options interface{}) (*CommonResponse, error) {
	url := fmt.Sprintf("%v%v", HOST, "/order/orders/update")
	resource := &CommonResponse{}
	err := xs.XsClient.Put(url, options, resource)
	return resource, err
}

func (xs *OrderServiceOp) Count(options interface{}) (int, error) {
	url := fmt.Sprintf("%v%v", HOST, "/order/orders/count")
	resource := &CountResponse{}
	err := xs.XsClient.Put(url, options, resource)
	count, err := strconv.Atoi(resource.Data.Count)
	return count, err
}
