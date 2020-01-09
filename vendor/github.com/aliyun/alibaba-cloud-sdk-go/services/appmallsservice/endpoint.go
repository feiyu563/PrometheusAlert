package appmallsservice

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "appms.aliyuncs.com",
			"cn-beijing":     "appms.aliyuncs.com",
			"ap-south-1":     "appms.aliyuncs.com",
			"eu-west-1":      "appms.aliyuncs.com",
			"ap-northeast-1": "appms.aliyuncs.com",
			"me-east-1":      "appms.aliyuncs.com",
			"cn-chengdu":     "appms.aliyuncs.com",
			"cn-qingdao":     "appms.aliyuncs.com",
			"cn-shanghai":    "appms.aliyuncs.com",
			"cn-hongkong":    "appms.aliyuncs.com",
			"ap-southeast-1": "appms.aliyuncs.com",
			"ap-southeast-2": "appms.aliyuncs.com",
			"ap-southeast-3": "appms.aliyuncs.com",
			"eu-central-1":   "appms.aliyuncs.com",
			"cn-huhehaote":   "appms.aliyuncs.com",
			"ap-southeast-5": "appms.aliyuncs.com",
			"us-east-1":      "appms.aliyuncs.com",
			"cn-zhangjiakou": "appms.aliyuncs.com",
			"us-west-1":      "appms.aliyuncs.com",
			"cn-hangzhou":    "appms.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
