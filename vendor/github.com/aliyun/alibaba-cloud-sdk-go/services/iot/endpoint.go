package iot

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "iot.aliyuncs.com",
			"cn-beijing-gov-1":            "iot.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "iot.aliyuncs.com",
			"cn-beijing":                  "iot.aliyuncs.com",
			"ap-south-1":                  "iot.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-inner":           "iot.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "iot.aliyuncs.com",
			"cn-haidian-cm12-c01":         "iot.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "iot.aliyuncs.com",
			"cn-north-2-gov-1":            "iot.aliyuncs.com",
			"cn-yushanfang":               "iot.aliyuncs.com",
			"cn-qingdao":                  "iot.aliyuncs.com",
			"cn-hongkong-finance-pop":     "iot.aliyuncs.com",
			"cn-qingdao-nebula":           "iot.aliyuncs.com",
			"cn-shanghai-finance-1":       "iot.aliyuncs.com",
			"cn-hongkong":                 "iot.aliyuncs.com",
			"cn-beijing-finance-pop":      "iot.aliyuncs.com",
			"cn-wuhan":                    "iot.aliyuncs.com",
			"cn-zhangjiakou":              "iot.aliyuncs.com",
			"cn-shenzhen":                 "iot.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "iot.aliyuncs.com",
			"rus-west-1-pop":              "iot.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "iot.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "iot.aliyuncs.com",
			"eu-west-1":                   "iot.ap-northeast-1.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "iot.aliyuncs.com",
			"eu-west-1-oxs":               "iot.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "iot.aliyuncs.com",
			"cn-beijing-finance-1":        "iot.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "iot.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "iot.aliyuncs.com",
			"cn-shenzhen-finance-1":       "iot.aliyuncs.com",
			"me-east-1":                   "iot.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":                  "iot.aliyuncs.com",
			"cn-hangzhou-test-306":        "iot.aliyuncs.com",
			"cn-shanghai-et2-b01":         "iot.aliyuncs.com",
			"cn-hangzhou-finance":         "iot.aliyuncs.com",
			"cn-beijing-nu16-b01":         "iot.aliyuncs.com",
			"cn-edge-1":                   "iot.aliyuncs.com",
			"ap-southeast-2":              "iot.ap-northeast-1.aliyuncs.com",
			"ap-southeast-3":              "iot.ap-northeast-1.aliyuncs.com",
			"cn-huhehaote":                "iot.aliyuncs.com",
			"ap-southeast-5":              "iot.ap-northeast-1.aliyuncs.com",
			"cn-fujian":                   "iot.aliyuncs.com",
			"ap-northeast-2-pop":          "iot.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "iot.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "iot.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
