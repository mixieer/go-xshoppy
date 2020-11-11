package main

import (
	"fmt"
	"go-xshoppy/xssdk"
	"reflect"
)

func main() {

	//shopCheck()
	//getHookList()
	getList()
	//getOrderList()
}

func getHookList() {
	auth := xssdk.App{
		ApiKey:       "",
		Password:     "",
		SharedSecret: "",
	}
	Client := xssdk.NewApiClient(auth)
	d := xssdk.DeleteOption{Topic: "order/create"}
	Client.Hook.Delete(d)
	res, _ := Client.Hook.List()

	fmt.Println(res)
}

func getList() {
	auth := xssdk.App{
		ApiKey:       "",
		Password:     "",
		SharedSecret: "",
	}

	Client := xssdk.NewApiClient(auth)
	list := xssdk.ListOptions{
		Filter: "order_paid",
		Limit:  "1",
		Page:   "1",
	}
	res, _ := Client.Order.List(list)
	fmt.Println(res)
}

func getOrderList() {

	data := xssdk.UpdateOptions{
		Id:    "",
		Email: "",
	}

	auth := xssdk.App{
		ApiKey:       "",
		Password:     "",
		SharedSecret: "",
	}

	Client := xssdk.NewApiClient(auth)
	Client.Order.Update(data)

}

func shopCheck() {
	auth := xssdk.App{
		ApiKey:       "",
		Password:     "",
		SharedSecret: "",
	}

	Client := xssdk.NewApiClient(auth)
	check := xssdk.CheckOption{ShopUrl: "xshoppy后台地址"}
	res, _ := Client.Shop.Check(check)
	fmt.Println(reflect.TypeOf(res.Data))
	a := res.Data

	v1, ok := a.(map[string]interface{})["result"]
	if ok {
		fmt.Println(v1)
	} else {
		fmt.Println("店铺异常")
	}
	fmt.Println(res.Data)
}
