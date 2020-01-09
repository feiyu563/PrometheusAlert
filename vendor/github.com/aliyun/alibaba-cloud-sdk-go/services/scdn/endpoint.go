package scdn

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "scdn.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "scdn.aliyuncs.com",
			"cn-beijing":                  "scdn.aliyuncs.com",
			"cn-shanghai-inner":           "scdn.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "scdn.aliyuncs.com",
			"cn-north-2-gov-1":            "scdn.aliyuncs.com",
			"cn-yushanfang":               "scdn.aliyuncs.com",
			"cn-qingdao-nebula":           "scdn.aliyuncs.com",
			"cn-beijing-finance-pop":      "scdn.aliyuncs.com",
			"cn-wuhan":                    "scdn.aliyuncs.com",
			"cn-zhangjiakou":              "scdn.aliyuncs.com",
			"us-west-1":                   "scdn.aliyuncs.com",
			"rus-west-1-pop":              "scdn.aliyuncs.com",
			"cn-shanghai-et15-b01":        "scdn.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "scdn.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "scdn.aliyuncs.com",
			"ap-northeast-1":              "scdn.aliyuncs.com",
			"cn-shanghai-et2-b01":         "scdn.aliyuncs.com",
			"ap-southeast-1":              "scdn.aliyuncs.com",
			"ap-southeast-2":              "scdn.aliyuncs.com",
			"ap-southeast-3":              "scdn.aliyuncs.com",
			"ap-southeast-5":              "scdn.aliyuncs.com",
			"us-east-1":                   "scdn.aliyuncs.com",
			"cn-shenzhen-inner":           "scdn.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "scdn.aliyuncs.com",
			"cn-beijing-gov-1":            "scdn.aliyuncs.com",
			"ap-south-1":                  "scdn.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "scdn.aliyuncs.com",
			"cn-haidian-cm12-c01":         "scdn.aliyuncs.com",
			"cn-qingdao":                  "scdn.aliyuncs.com",
			"cn-hongkong-finance-pop":     "scdn.aliyuncs.com",
			"cn-shanghai":                 "scdn.aliyuncs.com",
			"cn-shanghai-finance-1":       "scdn.aliyuncs.com",
			"cn-hongkong":                 "scdn.aliyuncs.com",
			"eu-central-1":                "scdn.aliyuncs.com",
			"cn-shenzhen":                 "scdn.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "scdn.aliyuncs.com",
			"eu-west-1":                   "scdn.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "scdn.aliyuncs.com",
			"eu-west-1-oxs":               "scdn.aliyuncs.com",
			"cn-beijing-finance-1":        "scdn.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "scdn.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "scdn.aliyuncs.com",
			"cn-shenzhen-finance-1":       "scdn.aliyuncs.com",
			"me-east-1":                   "scdn.aliyuncs.com",
			"cn-chengdu":                  "scdn.aliyuncs.com",
			"cn-hangzhou-test-306":        "scdn.aliyuncs.com",
			"cn-hangzhou-finance":         "scdn.aliyuncs.com",
			"cn-beijing-nu16-b01":         "scdn.aliyuncs.com",
			"cn-edge-1":                   "scdn.aliyuncs.com",
			"cn-huhehaote":                "scdn.aliyuncs.com",
			"cn-fujian":                   "scdn.aliyuncs.com",
			"ap-northeast-2-pop":          "scdn.aliyuncs.com",
			"cn-hangzhou":                 "scdn.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
