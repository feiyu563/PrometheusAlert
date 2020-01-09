package green

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"ap-south-1":            "green.ap-southeast-1.aliyuncs.com",
			"eu-west-1":             "green.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1":        "green.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "green.aliyuncs.com",
			"me-east-1":             "green.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":            "green.aliyuncs.com",
			"cn-north-2-gov-1":      "green.aliyuncs.com",
			"cn-qingdao":            "green.aliyuncs.com",
			"cn-shanghai-finance-1": "green.aliyuncs.com",
			"cn-hongkong":           "green.aliyuncs.com",
			"cn-hangzhou-finance":   "green.aliyuncs.com",
			"ap-southeast-2":        "green.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":        "green.ap-southeast-1.aliyuncs.com",
			"eu-central-1":          "green.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":          "green.aliyuncs.com",
			"ap-southeast-5":        "green.ap-southeast-1.aliyuncs.com",
			"us-east-1":             "green.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou":        "green.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
