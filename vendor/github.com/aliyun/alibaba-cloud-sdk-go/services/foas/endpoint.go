package foas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"ap-south-1":            "foas.ap-northeast-1.aliyuncs.com",
			"cn-hongkong":           "foas.aliyuncs.com",
			"eu-west-1":             "foas.ap-northeast-1.aliyuncs.com",
			"ap-southeast-2":        "foas.ap-northeast-1.aliyuncs.com",
			"eu-central-1":          "foas.ap-northeast-1.aliyuncs.com",
			"cn-huhehaote":          "foas.aliyuncs.com",
			"cn-shenzhen-finance-1": "foas.aliyuncs.com",
			"ap-southeast-5":        "foas.ap-northeast-1.aliyuncs.com",
			"us-east-1":             "foas.ap-northeast-1.aliyuncs.com",
			"me-east-1":             "foas.ap-northeast-1.aliyuncs.com",
			"us-west-1":             "foas.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":            "foas.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
