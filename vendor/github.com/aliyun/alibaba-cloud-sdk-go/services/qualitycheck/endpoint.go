package qualitycheck

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "qualitycheck.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "qualitycheck.aliyuncs.com",
			"cn-beijing":                  "qualitycheck.aliyuncs.com",
			"cn-shanghai-inner":           "qualitycheck.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "qualitycheck.aliyuncs.com",
			"cn-north-2-gov-1":            "qualitycheck.aliyuncs.com",
			"cn-yushanfang":               "qualitycheck.aliyuncs.com",
			"cn-qingdao-nebula":           "qualitycheck.aliyuncs.com",
			"cn-beijing-finance-pop":      "qualitycheck.aliyuncs.com",
			"cn-wuhan":                    "qualitycheck.aliyuncs.com",
			"cn-zhangjiakou":              "qualitycheck.aliyuncs.com",
			"us-west-1":                   "qualitycheck.aliyuncs.com",
			"rus-west-1-pop":              "qualitycheck.aliyuncs.com",
			"cn-shanghai-et15-b01":        "qualitycheck.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "qualitycheck.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "qualitycheck.aliyuncs.com",
			"ap-northeast-1":              "qualitycheck.aliyuncs.com",
			"cn-shanghai-et2-b01":         "qualitycheck.aliyuncs.com",
			"ap-southeast-1":              "qualitycheck.aliyuncs.com",
			"ap-southeast-2":              "qualitycheck.aliyuncs.com",
			"ap-southeast-3":              "qualitycheck.aliyuncs.com",
			"ap-southeast-5":              "qualitycheck.aliyuncs.com",
			"us-east-1":                   "qualitycheck.aliyuncs.com",
			"cn-shenzhen-inner":           "qualitycheck.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "qualitycheck.aliyuncs.com",
			"cn-beijing-gov-1":            "qualitycheck.aliyuncs.com",
			"ap-south-1":                  "qualitycheck.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "qualitycheck.aliyuncs.com",
			"cn-haidian-cm12-c01":         "qualitycheck.aliyuncs.com",
			"cn-qingdao":                  "qualitycheck.aliyuncs.com",
			"cn-hongkong-finance-pop":     "qualitycheck.aliyuncs.com",
			"cn-shanghai":                 "qualitycheck.aliyuncs.com",
			"cn-shanghai-finance-1":       "qualitycheck.aliyuncs.com",
			"cn-hongkong":                 "qualitycheck.aliyuncs.com",
			"eu-central-1":                "qualitycheck.aliyuncs.com",
			"cn-shenzhen":                 "qualitycheck.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "qualitycheck.aliyuncs.com",
			"eu-west-1":                   "qualitycheck.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "qualitycheck.aliyuncs.com",
			"eu-west-1-oxs":               "qualitycheck.aliyuncs.com",
			"cn-beijing-finance-1":        "qualitycheck.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "qualitycheck.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "qualitycheck.aliyuncs.com",
			"cn-shenzhen-finance-1":       "qualitycheck.aliyuncs.com",
			"me-east-1":                   "qualitycheck.aliyuncs.com",
			"cn-chengdu":                  "qualitycheck.aliyuncs.com",
			"cn-hangzhou-test-306":        "qualitycheck.aliyuncs.com",
			"cn-hangzhou-finance":         "qualitycheck.aliyuncs.com",
			"cn-beijing-nu16-b01":         "qualitycheck.aliyuncs.com",
			"cn-edge-1":                   "qualitycheck.aliyuncs.com",
			"cn-huhehaote":                "qualitycheck.aliyuncs.com",
			"cn-fujian":                   "qualitycheck.aliyuncs.com",
			"ap-northeast-2-pop":          "qualitycheck.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
