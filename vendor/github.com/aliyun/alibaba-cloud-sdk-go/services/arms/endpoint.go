package arms

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"ap-south-1":     "arms.ap-southeast-1.aliyuncs.com",
			"eu-west-1":      "arms.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2": "arms.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3": "arms.ap-southeast-1.aliyuncs.com",
			"eu-central-1":   "arms.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1": "arms.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":   "arms.aliyuncs.com",
			"ap-southeast-5": "arms.ap-southeast-1.aliyuncs.com",
			"us-east-1":      "arms.ap-southeast-1.aliyuncs.com",
			"me-east-1":      "arms.ap-southeast-1.aliyuncs.com",
			"us-west-1":      "arms.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":     "arms.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
