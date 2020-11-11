package xssdk

type Order struct {
	ShopUrl             string `json:"shop_url"`     //店铺访问的url
	Id                  string `json:"id"`           //订单号
	OrderNumber         string `json:"order_number"` //订单号带店铺名称
	CustomorId          string `json:"customor_id"`
	Gateway             string `json:"gateway"`
	TotalPrice          string `json:"total_price"`           //总金额
	SubtotalPrice       string `json:"subtotal_price"`        //折后商品总价，不包含折扣
	TotalTax            string `json:"total_tax"`             //总税价
	TotalDiscounts      string `json:"total_discounts"`       //总优惠金额
	TotalFulfillment    string `json:"total_fulfillment"`     //物流总价
	ShippingInsurance   string `json:"shipping_insurance"`    //选中的物流方案
	Currency            string `json:"currency"`              //货币
	ShippingMethod      string `json:"shipping_method"`       //买家自选物流
	Status              string `json:"status"`                //订单状态  init 初始化订单 pending 进行中订单 completed 已完成订单 cancel 取消订单 refunded（退款）
	FinancialStatus     string `json:"currency"`              //支付状态  unpaid 订单未支付 voided 订单支付失败 paid 订单支付成功 pending 待支付
	FulfillmentStatus   string `json:"fulfillment_status"`    //物流状态
	FulfillmentSendTime string `json:"fulfillment_send_time"` //物流发送时间
	BrowserIp           string `json:"browser_ip"`            //下单用户ip
	OrderName           string `json:"order_name"`            //订单名称
	OrderStatusUrl      string `json:"order_status_url"`      //订单详情查询地址
	TrackingNumber      string `json:"tracking_number"`       // 物流单号
	PayTime             string `json:"pay_time"`              //支付时间
	CreatedAt           string `json:"created_at"`            //订单创建时间
	Note                string `json:"note"`                  //订单备注（买家备注）
	products            []Products
	Billing             Billing
	Shipping            Shipping
}

//商品信息
type Products struct {
	Title    string `json:"title"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Sku      string `json:"sku"`
	AttrName string `json:"attr_name"`
	Image    string `json:"image"`
}

//账单信息
type Billing struct {
	Address1     string      `json:"address1"`
	Address2     string      `json:"address2"`
	City         string      `json:"city"`
	Company      interface{} `json:"company"`
	Country      string      `json:"country"`
	CountryCode  string      `json:"country_code"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	Province     string      `json:"province"`
	ProvinceCode interface{} `json:"province_code"`
	Zip          string      `json:"zip"`
}

type Shipping struct {
	Address1     string      `json:"address1"`
	Address2     string      `json:"address2"`
	City         string      `json:"city"`
	Company      interface{} `json:"company"`
	Country      string      `json:"country"`
	CountryCode  string      `json:"country_code"`
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	Province     string      `json:"province"`
	ProvinceCode interface{} `json:"province_code"`
	Zip          string      `json:"zip"`
	TaxNumber    string      `json:"tax_number"`
}
