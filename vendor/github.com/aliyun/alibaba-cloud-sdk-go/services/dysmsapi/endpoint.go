package dysmsapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"rus-west-1-pop":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"cn-beijing":         "dysmsapi-proxy.cn-beijing.aliyuncs.com",
			"ap-south-1":         "dysmsapi.ap-southeast-1.aliyuncs.com",
			"eu-west-1":          "dysmsapi.ap-southeast-1.aliyuncs.com",
			"eu-west-1-oxs":      "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"me-east-1":          "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-1":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"eu-central-1":       "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-5":     "dysmsapi.ap-southeast-1.aliyuncs.com",
			"us-east-1":          "dysmsapi.ap-southeast-1.aliyuncs.com",
			"ap-northeast-2-pop": "dysmsapi.ap-southeast-1.aliyuncs.com",
			"us-west-1":          "dysmsapi.ap-southeast-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
