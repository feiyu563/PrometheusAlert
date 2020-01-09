package baas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "baas.aliyuncs.com",
			"cn-beijing":     "baas.aliyuncs.com",
			"ap-south-1":     "baas.ap-southeast-1.aliyuncs.com",
			"eu-west-1":      "baas.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1": "baas.ap-southeast-1.aliyuncs.com",
			"cn-qingdao":     "baas.aliyuncs.com",
			"cn-shanghai":    "baas.aliyuncs.com",
			"cn-hongkong":    "baas.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2": "baas.ap-southeast-1.aliyuncs.com",
			"eu-central-1":   "baas.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":   "baas.aliyuncs.com",
			"us-east-1":      "baas.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou": "baas.aliyuncs.com",
			"us-west-1":      "baas.ap-southeast-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
