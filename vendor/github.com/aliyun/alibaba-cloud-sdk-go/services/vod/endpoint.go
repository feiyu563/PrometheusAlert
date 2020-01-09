package vod

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "vod.aliyuncs.com",
			"cn-beijing-gov-1":            "vod.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "vod.aliyuncs.com",
			"cn-beijing":                  "vod.cn-shanghai.aliyuncs.com",
			"cn-shanghai-inner":           "vod.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "vod.aliyuncs.com",
			"cn-haidian-cm12-c01":         "vod.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "vod.aliyuncs.com",
			"cn-yushanfang":               "vod.aliyuncs.com",
			"cn-qingdao":                  "vod.aliyuncs.com",
			"cn-hongkong-finance-pop":     "vod.aliyuncs.com",
			"cn-qingdao-nebula":           "vod.aliyuncs.com",
			"cn-shanghai-finance-1":       "vod.aliyuncs.com",
			"cn-beijing-finance-pop":      "vod.aliyuncs.com",
			"cn-wuhan":                    "vod.aliyuncs.com",
			"cn-shenzhen":                 "vod.cn-shanghai.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "vod.aliyuncs.com",
			"rus-west-1-pop":              "vod.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "vod.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "vod.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "vod.aliyuncs.com",
			"eu-west-1-oxs":               "vod.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "vod.aliyuncs.com",
			"cn-beijing-finance-1":        "vod.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "vod.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "vod.aliyuncs.com",
			"cn-shenzhen-finance-1":       "vod.aliyuncs.com",
			"me-east-1":                   "vod.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":                  "vod.aliyuncs.com",
			"cn-hangzhou-test-306":        "vod.aliyuncs.com",
			"cn-shanghai-et2-b01":         "vod.aliyuncs.com",
			"cn-hangzhou-finance":         "vod.aliyuncs.com",
			"cn-beijing-nu16-b01":         "vod.aliyuncs.com",
			"cn-edge-1":                   "vod.aliyuncs.com",
			"ap-southeast-2":              "vod.ap-northeast-1.aliyuncs.com",
			"ap-southeast-3":              "vod.ap-northeast-1.aliyuncs.com",
			"cn-huhehaote":                "vod.aliyuncs.com",
			"cn-fujian":                   "vod.aliyuncs.com",
			"us-east-1":                   "vod.ap-northeast-1.aliyuncs.com",
			"ap-northeast-2-pop":          "vod.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "vod.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "vod.aliyuncs.com",
			"cn-hangzhou":                 "vod.cn-shanghai.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
