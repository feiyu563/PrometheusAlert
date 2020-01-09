package actiontrail

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "actiontrail.aliyuncs.com",
			"cn-beijing-gov-1":            "actiontrail.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "actiontrail.aliyuncs.com",
			"cn-shanghai-inner":           "actiontrail.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "actiontrail.aliyuncs.com",
			"cn-haidian-cm12-c01":         "actiontrail.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "actiontrail.aliyuncs.com",
			"cn-yushanfang":               "actiontrail.aliyuncs.com",
			"cn-hongkong-finance-pop":     "actiontrail.aliyuncs.com",
			"cn-qingdao-nebula":           "actiontrail.aliyuncs.com",
			"cn-beijing-finance-pop":      "actiontrail.aliyuncs.com",
			"cn-wuhan":                    "actiontrail.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "actiontrail.aliyuncs.com",
			"rus-west-1-pop":              "actiontrail.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "actiontrail.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "actiontrail.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "actiontrail.aliyuncs.com",
			"eu-west-1-oxs":               "actiontrail.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "actiontrail.aliyuncs.com",
			"cn-beijing-finance-1":        "actiontrail.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "actiontrail.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "actiontrail.aliyuncs.com",
			"cn-shenzhen-finance-1":       "actiontrail.aliyuncs.com",
			"cn-hangzhou-test-306":        "actiontrail.aliyuncs.com",
			"cn-shanghai-et2-b01":         "actiontrail.aliyuncs.com",
			"cn-hangzhou-finance":         "actiontrail.aliyuncs.com",
			"cn-beijing-nu16-b01":         "actiontrail.aliyuncs.com",
			"cn-edge-1":                   "actiontrail.aliyuncs.com",
			"cn-fujian":                   "actiontrail.aliyuncs.com",
			"ap-northeast-2-pop":          "actiontrail.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "actiontrail.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "actiontrail.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
