package ons

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "ons.aliyuncs.com",
			"cn-beijing-gov-1":            "ons.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "ons.aliyuncs.com",
			"cn-shanghai-inner":           "ons.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "ons.aliyuncs.com",
			"cn-haidian-cm12-c01":         "ons.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "ons.aliyuncs.com",
			"cn-yushanfang":               "ons.aliyuncs.com",
			"cn-hongkong-finance-pop":     "ons.aliyuncs.com",
			"cn-qingdao-nebula":           "ons.aliyuncs.com",
			"cn-beijing-finance-pop":      "ons.aliyuncs.com",
			"cn-wuhan":                    "ons.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "ons.aliyuncs.com",
			"rus-west-1-pop":              "ons.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "ons.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "ons.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "ons.aliyuncs.com",
			"eu-west-1-oxs":               "ons.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "ons.aliyuncs.com",
			"cn-beijing-finance-1":        "ons.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "ons.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "ons.aliyuncs.com",
			"cn-hangzhou-test-306":        "ons.aliyuncs.com",
			"cn-shanghai-et2-b01":         "ons.aliyuncs.com",
			"cn-beijing-nu16-b01":         "ons.aliyuncs.com",
			"cn-edge-1":                   "ons.aliyuncs.com",
			"cn-fujian":                   "ons.aliyuncs.com",
			"ap-northeast-2-pop":          "ons.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "ons.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "ons.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
