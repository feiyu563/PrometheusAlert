package market

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "market.aliyuncs.com",
			"cn-beijing":            "market.aliyuncs.com",
			"ap-south-1":            "market.ap-southeast-1.aliyuncs.com",
			"eu-west-1":             "market.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1":        "market.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "market.aliyuncs.com",
			"me-east-1":             "market.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":            "market.aliyuncs.com",
			"cn-north-2-gov-1":      "market.aliyuncs.com",
			"cn-qingdao":            "market.aliyuncs.com",
			"cn-shanghai":           "market.aliyuncs.com",
			"cn-shanghai-finance-1": "market.aliyuncs.com",
			"cn-hongkong":           "market.aliyuncs.com",
			"cn-hangzhou-finance":   "market.aliyuncs.com",
			"ap-southeast-2":        "market.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":        "market.ap-southeast-1.aliyuncs.com",
			"eu-central-1":          "market.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":          "market.aliyuncs.com",
			"ap-southeast-5":        "market.ap-southeast-1.aliyuncs.com",
			"us-east-1":             "market.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou":        "market.aliyuncs.com",
			"us-west-1":             "market.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou":           "market.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
