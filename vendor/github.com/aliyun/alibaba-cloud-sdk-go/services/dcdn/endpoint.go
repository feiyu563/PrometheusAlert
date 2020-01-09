package dcdn

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "dcdn.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "dcdn.aliyuncs.com",
			"cn-beijing":                  "dcdn.aliyuncs.com",
			"cn-shanghai-inner":           "dcdn.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "dcdn.aliyuncs.com",
			"cn-north-2-gov-1":            "dcdn.aliyuncs.com",
			"cn-yushanfang":               "dcdn.aliyuncs.com",
			"cn-qingdao-nebula":           "dcdn.aliyuncs.com",
			"cn-beijing-finance-pop":      "dcdn.aliyuncs.com",
			"cn-wuhan":                    "dcdn.aliyuncs.com",
			"cn-zhangjiakou":              "dcdn.aliyuncs.com",
			"us-west-1":                   "dcdn.aliyuncs.com",
			"rus-west-1-pop":              "dcdn.aliyuncs.com",
			"cn-shanghai-et15-b01":        "dcdn.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "dcdn.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "dcdn.aliyuncs.com",
			"ap-northeast-1":              "dcdn.aliyuncs.com",
			"cn-shanghai-et2-b01":         "dcdn.aliyuncs.com",
			"ap-southeast-1":              "dcdn.aliyuncs.com",
			"ap-southeast-2":              "dcdn.aliyuncs.com",
			"ap-southeast-3":              "dcdn.aliyuncs.com",
			"ap-southeast-5":              "dcdn.aliyuncs.com",
			"us-east-1":                   "dcdn.aliyuncs.com",
			"cn-shenzhen-inner":           "dcdn.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "dcdn.aliyuncs.com",
			"cn-beijing-gov-1":            "dcdn.aliyuncs.com",
			"ap-south-1":                  "dcdn.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "dcdn.aliyuncs.com",
			"cn-haidian-cm12-c01":         "dcdn.aliyuncs.com",
			"cn-qingdao":                  "dcdn.aliyuncs.com",
			"cn-hongkong-finance-pop":     "dcdn.aliyuncs.com",
			"cn-shanghai":                 "dcdn.aliyuncs.com",
			"cn-shanghai-finance-1":       "dcdn.aliyuncs.com",
			"cn-hongkong":                 "dcdn.aliyuncs.com",
			"eu-central-1":                "dcdn.aliyuncs.com",
			"cn-shenzhen":                 "dcdn.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "dcdn.aliyuncs.com",
			"eu-west-1":                   "dcdn.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "dcdn.aliyuncs.com",
			"eu-west-1-oxs":               "dcdn.aliyuncs.com",
			"cn-beijing-finance-1":        "dcdn.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "dcdn.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "dcdn.aliyuncs.com",
			"cn-shenzhen-finance-1":       "dcdn.aliyuncs.com",
			"me-east-1":                   "dcdn.aliyuncs.com",
			"cn-chengdu":                  "dcdn.aliyuncs.com",
			"cn-hangzhou-test-306":        "dcdn.aliyuncs.com",
			"cn-hangzhou-finance":         "dcdn.aliyuncs.com",
			"cn-beijing-nu16-b01":         "dcdn.aliyuncs.com",
			"cn-edge-1":                   "dcdn.aliyuncs.com",
			"cn-huhehaote":                "dcdn.aliyuncs.com",
			"cn-fujian":                   "dcdn.aliyuncs.com",
			"ap-northeast-2-pop":          "dcdn.aliyuncs.com",
			"cn-hangzhou":                 "dcdn.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
