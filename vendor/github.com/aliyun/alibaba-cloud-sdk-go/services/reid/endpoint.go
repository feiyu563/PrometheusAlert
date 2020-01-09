package reid

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "reid.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "reid.aliyuncs.com",
			"cn-beijing":                  "reid.aliyuncs.com",
			"cn-shanghai-inner":           "reid.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "reid.aliyuncs.com",
			"cn-north-2-gov-1":            "reid.aliyuncs.com",
			"cn-yushanfang":               "reid.aliyuncs.com",
			"cn-qingdao-nebula":           "reid.aliyuncs.com",
			"cn-beijing-finance-pop":      "reid.aliyuncs.com",
			"cn-wuhan":                    "reid.aliyuncs.com",
			"cn-zhangjiakou":              "reid.aliyuncs.com",
			"us-west-1":                   "reid.aliyuncs.com",
			"rus-west-1-pop":              "reid.aliyuncs.com",
			"cn-shanghai-et15-b01":        "reid.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "reid.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "reid.aliyuncs.com",
			"ap-northeast-1":              "reid.aliyuncs.com",
			"cn-shanghai-et2-b01":         "reid.aliyuncs.com",
			"ap-southeast-1":              "reid.aliyuncs.com",
			"ap-southeast-2":              "reid.aliyuncs.com",
			"ap-southeast-3":              "reid.aliyuncs.com",
			"ap-southeast-5":              "reid.aliyuncs.com",
			"us-east-1":                   "reid.aliyuncs.com",
			"cn-shenzhen-inner":           "reid.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "reid.aliyuncs.com",
			"cn-beijing-gov-1":            "reid.aliyuncs.com",
			"ap-south-1":                  "reid.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "reid.aliyuncs.com",
			"cn-haidian-cm12-c01":         "reid.aliyuncs.com",
			"cn-qingdao":                  "reid.aliyuncs.com",
			"cn-hongkong-finance-pop":     "reid.aliyuncs.com",
			"cn-shanghai":                 "reid.aliyuncs.com",
			"cn-shanghai-finance-1":       "reid.aliyuncs.com",
			"cn-hongkong":                 "reid.aliyuncs.com",
			"eu-central-1":                "reid.aliyuncs.com",
			"cn-shenzhen":                 "reid.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "reid.aliyuncs.com",
			"eu-west-1":                   "reid.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "reid.aliyuncs.com",
			"eu-west-1-oxs":               "reid.aliyuncs.com",
			"cn-beijing-finance-1":        "reid.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "reid.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "reid.aliyuncs.com",
			"cn-shenzhen-finance-1":       "reid.aliyuncs.com",
			"me-east-1":                   "reid.aliyuncs.com",
			"cn-chengdu":                  "reid.aliyuncs.com",
			"cn-hangzhou-test-306":        "reid.aliyuncs.com",
			"cn-hangzhou-finance":         "reid.aliyuncs.com",
			"cn-beijing-nu16-b01":         "reid.aliyuncs.com",
			"cn-edge-1":                   "reid.aliyuncs.com",
			"cn-huhehaote":                "reid.aliyuncs.com",
			"cn-fujian":                   "reid.aliyuncs.com",
			"ap-northeast-2-pop":          "reid.aliyuncs.com",
			"cn-hangzhou":                 "reid.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
