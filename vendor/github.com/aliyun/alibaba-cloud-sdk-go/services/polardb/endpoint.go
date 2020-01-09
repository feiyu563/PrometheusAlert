package polardb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "polardb.aliyuncs.com",
			"cn-beijing":            "polardb.aliyuncs.com",
			"cn-shenzhen-finance-1": "polardb.aliyuncs.com",
			"cn-north-2-gov-1":      "polardb.aliyuncs.com",
			"cn-qingdao":            "polardb.aliyuncs.com",
			"cn-shanghai":           "polardb.aliyuncs.com",
			"cn-shanghai-finance-1": "polardb.aliyuncs.com",
			"cn-hongkong":           "polardb.aliyuncs.com",
			"cn-hangzhou-finance":   "polardb.aliyuncs.com",
			"ap-southeast-1":        "polardb.aliyuncs.com",
			"us-east-1":             "polardb.ap-northeast-1.aliyuncs.com",
			"us-west-1":             "polardb.aliyuncs.com",
			"cn-hangzhou":           "polardb.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
