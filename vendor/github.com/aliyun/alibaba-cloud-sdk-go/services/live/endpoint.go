package live

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "live.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "live.aliyuncs.com",
			"cn-beijing":                  "live.aliyuncs.com",
			"cn-shanghai-inner":           "live.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "live.aliyuncs.com",
			"cn-north-2-gov-1":            "live.aliyuncs.com",
			"cn-yushanfang":               "live.aliyuncs.com",
			"cn-qingdao-nebula":           "live.aliyuncs.com",
			"cn-beijing-finance-pop":      "live.aliyuncs.com",
			"cn-wuhan":                    "live.aliyuncs.com",
			"cn-zhangjiakou":              "live.aliyuncs.com",
			"us-west-1":                   "live.ap-southeast-1.aliyuncs.com",
			"rus-west-1-pop":              "live.ap-southeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "live.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "live.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "live.aliyuncs.com",
			"ap-northeast-1":              "live.aliyuncs.com",
			"cn-shanghai-et2-b01":         "live.aliyuncs.com",
			"ap-southeast-1":              "live.aliyuncs.com",
			"ap-southeast-2":              "live.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":              "live.ap-southeast-1.aliyuncs.com",
			"ap-southeast-5":              "live.aliyuncs.com",
			"us-east-1":                   "live.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "live.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "live.aliyuncs.com",
			"cn-beijing-gov-1":            "live.aliyuncs.com",
			"ap-south-1":                  "live.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "live.aliyuncs.com",
			"cn-haidian-cm12-c01":         "live.aliyuncs.com",
			"cn-qingdao":                  "live.aliyuncs.com",
			"cn-hongkong-finance-pop":     "live.aliyuncs.com",
			"cn-shanghai":                 "live.aliyuncs.com",
			"cn-shanghai-finance-1":       "live.aliyuncs.com",
			"cn-hongkong":                 "live.aliyuncs.com",
			"eu-central-1":                "live.aliyuncs.com",
			"cn-shenzhen":                 "live.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "live.aliyuncs.com",
			"eu-west-1":                   "live.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "live.aliyuncs.com",
			"eu-west-1-oxs":               "live.ap-southeast-1.aliyuncs.com",
			"cn-beijing-finance-1":        "live.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "live.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "live.aliyuncs.com",
			"cn-shenzhen-finance-1":       "live.aliyuncs.com",
			"me-east-1":                   "live.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":                  "live.aliyuncs.com",
			"cn-hangzhou-test-306":        "live.aliyuncs.com",
			"cn-hangzhou-finance":         "live.aliyuncs.com",
			"cn-beijing-nu16-b01":         "live.aliyuncs.com",
			"cn-edge-1":                   "live.aliyuncs.com",
			"cn-huhehaote":                "live.aliyuncs.com",
			"cn-fujian":                   "live.aliyuncs.com",
			"ap-northeast-2-pop":          "live.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou":                 "live.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
