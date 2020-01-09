package cusanalytic_sc_online

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "cusanalytic.aliyuncs.com",
			"cn-beijing":            "cusanalytic.aliyuncs.com",
			"ap-south-1":            "cusanalytic.aliyuncs.com",
			"eu-west-1":             "cusanalytic.aliyuncs.com",
			"ap-northeast-1":        "cusanalytic.aliyuncs.com",
			"cn-shenzhen-finance-1": "cusanalytic.aliyuncs.com",
			"me-east-1":             "cusanalytic.aliyuncs.com",
			"cn-chengdu":            "cusanalytic.aliyuncs.com",
			"cn-north-2-gov-1":      "cusanalytic.aliyuncs.com",
			"cn-qingdao":            "cusanalytic.aliyuncs.com",
			"cn-shanghai":           "cusanalytic.aliyuncs.com",
			"cn-shanghai-finance-1": "cusanalytic.aliyuncs.com",
			"cn-hongkong":           "cusanalytic.aliyuncs.com",
			"cn-hangzhou-finance":   "cusanalytic.aliyuncs.com",
			"ap-southeast-1":        "cusanalytic.aliyuncs.com",
			"ap-southeast-2":        "cusanalytic.aliyuncs.com",
			"ap-southeast-3":        "cusanalytic.aliyuncs.com",
			"eu-central-1":          "cusanalytic.aliyuncs.com",
			"cn-huhehaote":          "cusanalytic.aliyuncs.com",
			"ap-southeast-5":        "cusanalytic.aliyuncs.com",
			"us-east-1":             "cusanalytic.aliyuncs.com",
			"cn-zhangjiakou":        "cusanalytic.aliyuncs.com",
			"us-west-1":             "cusanalytic.aliyuncs.com",
			"cn-hangzhou":           "cusanalytic.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
