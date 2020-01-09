package adb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "adb.aliyuncs.com",
			"cn-beijing-gov-1":            "adb.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "adb.aliyuncs.com",
			"cn-beijing":                  "adb.aliyuncs.com",
			"cn-shanghai-inner":           "adb.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "adb.aliyuncs.com",
			"cn-haidian-cm12-c01":         "adb.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "adb.aliyuncs.com",
			"cn-north-2-gov-1":            "adb.aliyuncs.com",
			"cn-yushanfang":               "adb.aliyuncs.com",
			"cn-qingdao":                  "adb.aliyuncs.com",
			"cn-hongkong-finance-pop":     "adb.aliyuncs.com",
			"cn-qingdao-nebula":           "adb.aliyuncs.com",
			"cn-shanghai":                 "adb.aliyuncs.com",
			"cn-shanghai-finance-1":       "adb.aliyuncs.com",
			"cn-hongkong":                 "adb.aliyuncs.com",
			"cn-beijing-finance-pop":      "adb.aliyuncs.com",
			"cn-wuhan":                    "adb.aliyuncs.com",
			"us-west-1":                   "adb.aliyuncs.com",
			"cn-shenzhen":                 "adb.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "adb.aliyuncs.com",
			"rus-west-1-pop":              "adb.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "adb.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "adb.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "adb.aliyuncs.com",
			"eu-west-1-oxs":               "adb.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "adb.aliyuncs.com",
			"cn-beijing-finance-1":        "adb.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "adb.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "adb.aliyuncs.com",
			"cn-shenzhen-finance-1":       "adb.aliyuncs.com",
			"me-east-1":                   "adb.ap-northeast-1.aliyuncs.com",
			"cn-hangzhou-test-306":        "adb.aliyuncs.com",
			"cn-shanghai-et2-b01":         "adb.aliyuncs.com",
			"cn-hangzhou-finance":         "adb.aliyuncs.com",
			"ap-southeast-1":              "adb.aliyuncs.com",
			"cn-beijing-nu16-b01":         "adb.aliyuncs.com",
			"cn-edge-1":                   "adb.aliyuncs.com",
			"us-east-1":                   "adb.aliyuncs.com",
			"cn-fujian":                   "adb.aliyuncs.com",
			"ap-northeast-2-pop":          "adb.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "adb.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "adb.aliyuncs.com",
			"cn-hangzhou":                 "adb.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
