package xssdk

import "fmt"

type HookService interface {
	List() (*ListResponse, error)
	Create(UpdateOption) error
	Update(UpdateOption) error
	Delete(DeleteOption) error
}

type HookServiceOp struct {
	XsClient *XsApiClient
}

type UpdateOption struct {
	Topic    string `json:"topic"`
	Name     string `json:"name"`
	Callback string `json:"callback"`
}



type DeleteOption struct {
	Topic string `json:"topic"`
}

type ListResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data []ListResponseData `json:"data"`
}

type ListResponseData struct {
	Topic      string `json:"topic"`
	Name       string `json:"name"`
	Callback   string `json:"callback"`
	CreateTime string `json:"create_time"`
}

func (h *HookServiceOp) List() (*ListResponse, error) {
	url := fmt.Sprintf("%v%v", HOST, "/hook/hook/list")
	resource := &ListResponse{}
	err := h.XsClient.Get(url, resource, nil)
	return resource, err
}

func (h *HookServiceOp) Create(u UpdateOption) error {
	url := fmt.Sprintf("%v%v", HOST, "/hook/hook/create")
	u.Callback = fmt.Sprintf("%v?act=%v", u.Callback, u.Topic)
	err := h.XsClient.Post(url, u, nil, nil)
	return err
}

func (h *HookServiceOp) Update(u UpdateOption) error {
	url := fmt.Sprintf("%v%v", HOST, "/hook/hook/update")
	u.Callback = fmt.Sprintf("%v?act=%v", u.Callback, u.Topic)
	err := h.XsClient.Put(url, u, nil)
	return err
}

func (h *HookServiceOp) Delete(d DeleteOption) error {
	url := fmt.Sprintf("%v%v", HOST, "/hook/hook/delete")
	err := h.XsClient.Delete(url, d)
	return err
}
