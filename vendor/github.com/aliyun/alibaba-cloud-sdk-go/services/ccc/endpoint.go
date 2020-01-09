package ccc

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "ccc.aliyuncs.com",
			"cn-beijing":            "ccc.aliyuncs.com",
			"ap-south-1":            "ccc.aliyuncs.com",
			"eu-west-1":             "ccc.aliyuncs.com",
			"ap-northeast-1":        "ccc.aliyuncs.com",
			"cn-shenzhen-finance-1": "ccc.aliyuncs.com",
			"me-east-1":             "ccc.aliyuncs.com",
			"cn-chengdu":            "ccc.aliyuncs.com",
			"cn-north-2-gov-1":      "ccc.aliyuncs.com",
			"cn-qingdao":            "ccc.aliyuncs.com",
			"cn-shanghai-finance-1": "ccc.aliyuncs.com",
			"cn-hongkong":           "ccc.aliyuncs.com",
			"cn-hangzhou-finance":   "ccc.aliyuncs.com",
			"ap-southeast-1":        "ccc.aliyuncs.com",
			"ap-southeast-2":        "ccc.aliyuncs.com",
			"ap-southeast-3":        "ccc.aliyuncs.com",
			"eu-central-1":          "ccc.aliyuncs.com",
			"cn-huhehaote":          "ccc.aliyuncs.com",
			"ap-southeast-5":        "ccc.aliyuncs.com",
			"us-east-1":             "ccc.aliyuncs.com",
			"cn-zhangjiakou":        "ccc.aliyuncs.com",
			"us-west-1":             "ccc.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
