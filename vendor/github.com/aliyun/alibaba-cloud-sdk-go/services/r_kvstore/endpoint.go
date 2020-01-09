package r_kvstore

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "r-kvstore.aliyuncs.com",
			"cn-beijing-gov-1":            "r-kvstore.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "r-kvstore.aliyuncs.com",
			"cn-beijing":                  "r-kvstore.aliyuncs.com",
			"cn-shanghai-inner":           "r-kvstore.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "r-kvstore.aliyuncs.com",
			"cn-haidian-cm12-c01":         "r-kvstore.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "r-kvstore.aliyuncs.com",
			"cn-north-2-gov-1":            "r-kvstore.aliyuncs.com",
			"cn-yushanfang":               "r-kvstore.aliyuncs.com",
			"cn-hongkong-finance-pop":     "r-kvstore.aliyuncs.com",
			"cn-qingdao-nebula":           "r-kvstore.aliyuncs.com",
			"cn-shanghai":                 "r-kvstore.aliyuncs.com",
			"cn-shanghai-finance-1":       "r-kvstore.aliyuncs.com",
			"cn-beijing-finance-pop":      "r-kvstore.aliyuncs.com",
			"cn-wuhan":                    "r-kvstore.aliyuncs.com",
			"us-west-1":                   "r-kvstore.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen":                 "r-kvstore.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "r-kvstore.aliyuncs.com",
			"rus-west-1-pop":              "r-kvstore.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "r-kvstore.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "r-kvstore.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "r-kvstore.aliyuncs.com",
			"eu-west-1-oxs":               "r-kvstore.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "r-kvstore.aliyuncs.com",
			"cn-beijing-finance-1":        "r-kvstore.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "r-kvstore.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "r-kvstore.aliyuncs.com",
			"cn-shenzhen-finance-1":       "r-kvstore.aliyuncs.com",
			"cn-hangzhou-test-306":        "r-kvstore.aliyuncs.com",
			"cn-shanghai-et2-b01":         "r-kvstore.aliyuncs.com",
			"cn-hangzhou-finance":         "r-kvstore.aliyuncs.com",
			"cn-beijing-nu16-b01":         "r-kvstore.aliyuncs.com",
			"cn-edge-1":                   "r-kvstore.aliyuncs.com",
			"cn-fujian":                   "r-kvstore.aliyuncs.com",
			"us-east-1":                   "r-kvstore.ap-northeast-1.aliyuncs.com",
			"ap-northeast-2-pop":          "r-kvstore.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "r-kvstore.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "r-kvstore.aliyuncs.com",
			"cn-hangzhou":                 "r-kvstore-cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
