package webplus

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "webplus.cn-hangzhou.aliyuncs.com",
			"cn-beijing":            "webplus.cn-hangzhou.aliyuncs.com",
			"ap-south-1":            "webplus.aliyuncs.com",
			"eu-west-1":             "webplus.aliyuncs.com",
			"ap-northeast-1":        "webplus.aliyuncs.com",
			"cn-shenzhen-finance-1": "webplus.aliyuncs.com",
			"me-east-1":             "webplus.aliyuncs.com",
			"cn-chengdu":            "webplus.aliyuncs.com",
			"cn-north-2-gov-1":      "webplus.aliyuncs.com",
			"cn-qingdao":            "webplus.aliyuncs.com",
			"cn-shanghai":           "webplus.cn-hangzhou.aliyuncs.com",
			"cn-shanghai-finance-1": "webplus.aliyuncs.com",
			"cn-hongkong":           "webplus-vpc.cn-hongkong.aliyuncs.com",
			"cn-hangzhou-finance":   "webplus.aliyuncs.com",
			"ap-southeast-1":        "webplus.aliyuncs.com",
			"ap-southeast-2":        "webplus.aliyuncs.com",
			"ap-southeast-3":        "webplus.aliyuncs.com",
			"eu-central-1":          "webplus.aliyuncs.com",
			"cn-huhehaote":          "webplus.aliyuncs.com",
			"ap-southeast-5":        "webplus.aliyuncs.com",
			"us-east-1":             "webplus.aliyuncs.com",
			"cn-zhangjiakou":        "webplus.cn-hangzhou.aliyuncs.com",
			"us-west-1":             "webplus.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
