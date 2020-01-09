package requests

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_AcsRequest(t *testing.T) {
	r := defaultBaseRequest()
	assert.NotNil(t, r)

	// query params
	query := r.GetQueryParams()
	assert.Equal(t, 0, len(query))
	r.addQueryParam("key", "value")
	assert.Equal(t, 1, len(query))
	assert.Equal(t, "value", query["key"])

	// form params
	form := r.GetFormParams()
	assert.Equal(t, 0, len(form))
	r.addFormParam("key", "value")
	assert.Equal(t, 1, len(form))
	assert.Equal(t, "value", form["key"])

	// getter/setter for stringtosign
	assert.Equal(t, "", r.GetStringToSign())
	r.SetStringToSign("s2s")
	assert.Equal(t, "s2s", r.GetStringToSign())

	// content type
	_, contains := r.GetContentType()
	assert.False(t, contains)
	r.SetContentType("application/json")
	ct, contains := r.GetContentType()
	assert.Equal(t, "application/json", ct)
	assert.True(t, contains)

	// default 3 headers & content-type
	headers := r.GetHeaders()
	assert.Equal(t, 4, len(headers))
	r.addHeaderParam("x-key", "x-key-value")
	assert.Equal(t, 5, len(headers))
	assert.Equal(t, "x-key-value", headers["x-key"])

	// Version
	r.SetVersion("2017-06-06")
	assert.Equal(t, "2017-06-06", r.GetVersion())
	// GetActionName
	assert.Equal(t, "", r.GetActionName())

	// GetMethod
	assert.Equal(t, "GET", r.GetMethod())
	r.Method = "POST"
	assert.Equal(t, "POST", r.GetMethod())

	// Domain
	assert.Equal(t, "", r.GetDomain())
	r.SetDomain("ecs.aliyuncs.com")
	assert.Equal(t, "ecs.aliyuncs.com", r.GetDomain())

	// Region
	assert.Equal(t, "", r.GetRegionId())
	r.RegionId = "cn-hangzhou"
	assert.Equal(t, "cn-hangzhou", r.GetRegionId())

	// AcceptFormat
	assert.Equal(t, "JSON", r.GetAcceptFormat())
	r.AcceptFormat = "XML"
	assert.Equal(t, "XML", r.GetAcceptFormat())

	// GetLocationServiceCode
	assert.Equal(t, "", r.GetLocationServiceCode())

	// GetLocationEndpointType
	assert.Equal(t, "", r.GetLocationEndpointType())

	// GetProduct
	assert.Equal(t, "", r.GetProduct())

	// GetScheme
	assert.Equal(t, "", r.GetScheme())
	r.SetScheme("HTTPS")
	assert.Equal(t, "HTTPS", r.GetScheme())

	// GetReadTimeout
	assert.Equal(t, 0*time.Second, r.GetReadTimeout())
	r.SetReadTimeout(5 * time.Second)
	assert.Equal(t, 5*time.Second, r.GetReadTimeout())

	// GetConnectTimeout
	assert.Equal(t, 0*time.Second, r.GetConnectTimeout())
	r.SetConnectTimeout(5 * time.Second)
	assert.Equal(t, 5*time.Second, r.GetConnectTimeout())

	// GetHTTPSInsecure
	assert.True(t, r.GetHTTPSInsecure() == nil)
	r.SetHTTPSInsecure(true)
	assert.Equal(t, true, *r.GetHTTPSInsecure())

	// GetPort
	assert.Equal(t, "", r.GetPort())

	// GetUserAgent
	r.AppendUserAgent("cli", "1.01")
	assert.Equal(t, "1.01", r.GetUserAgent()["cli"])
	// GetUserAgent
	r.AppendUserAgent("cli", "2.02")
	assert.Equal(t, "2.02", r.GetUserAgent()["cli"])
	// Content
	assert.Equal(t, []byte(nil), r.GetContent())
	r.SetContent([]byte("The Content"))
	assert.True(t, bytes.Equal([]byte("The Content"), r.GetContent()))
}

type AcsRequestTest struct {
	*baseRequest
	Ontology AcsRequest
	Query    string                 `position:"Query" name:"Query"`
	Header   string                 `position:"Header" name:"Header"`
	Path     string                 `position:"Path" name:"Path"`
	Body     string                 `position:"Body" name:"Body"`
	Target   map[string]interface{} `position:"Query" name:"Target"`
	TypeAcs  *[]string              `position:"type" name:"type" type:"Repeated"`
}

func (r AcsRequestTest) BuildQueries() string {
	return ""
}

func (r AcsRequestTest) BuildUrl() string {
	return ""
}

func (r AcsRequestTest) GetBodyReader() io.Reader {
	return nil
}

func (r AcsRequestTest) GetStyle() string {
	return ""
}

func (r AcsRequestTest) addPathParam(key, value string) {
	return
}

func Test_AcsRequest_InitParams(t *testing.T) {
	r := &AcsRequestTest{
		baseRequest: defaultBaseRequest(),
		Query:       "query value",
		Header:      "header value",
		Path:        "path value",
		Body:        "body value",
		Target: map[string]interface{}{
			"key":   "test",
			"value": 1234,
		},
	}
	tmp := []string{r.Query, r.Header}
	r.TypeAcs = &tmp
	r.addQueryParam("qkey", "qvalue")
	InitParams(r)

	queries := r.GetQueryParams()
	assert.Equal(t, "query value", queries["Query"])
	assert.Equal(t, "{\"key\":\"test\",\"value\":1234}", queries["Target"])
	headers := r.GetHeaders()
	assert.Equal(t, "header value", headers["Header"])
	// TODO: check the body & path
}

// CreateContainerGroupRequest is the request struct for api CreateContainerGroup
type CreateContainerGroupRequest struct {
	*RpcRequest
	OwnerId              Integer                       `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	Volume               *[]CreateContainerGroupVolume `position:"Query" name:"Volume" type:"Repeated"`
	Memory               Float                         `position:"Query" name:"Memory"`
	SlsEnable            Boolean                       `position:"Query" name:"SlsEnable"`
	DnsConfig            CreateContainerGroupDnsConfig `position:"Query" name:"DnsConfig" type:"Struct"`
	OptionJson           map[string]interface{}        `position:"Query" name:"OptionJson"`
}

type CreateContainerGroupVolume struct {
	Name      string                        `name:"Name"`
	Type      string                        `name:"Type"`
	NFSVolume CreateContainerGroupNFSVolume `name:"NFSVolume" type:"Struct"`
}

type CreateContainerGroupDnsConfig struct {
	NameServer []string                      `name:"NameServer"`
	Search     []string                      `name:"Search"`
	Option     *[]CreateContainerGroupOption `name:"Option" type:"Repeated"`
}

type CreateContainerGroupNFSVolume struct {
	Server   string                     `name:"Server"`
	Path     string                     `name:"Path"`
	ReadOnly Boolean                    `name:"ReadOnly"`
	Option   CreateContainerGroupOption `name:"Option" type:"Struct"`
}

type CreateContainerGroupOption struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

func GetQueryString(r *CreateContainerGroupRequest) string {
	queries := r.GetQueryParams()
	// To store the keys in slice in sorted order
	sortedKeys := make([]string, 0)
	for k := range queries {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	// To perform the opertion you want
	resultBuilder := bytes.Buffer{}
	for _, key := range sortedKeys {
		resultBuilder.WriteString(key + "=" + queries[key] + "&")
	}
	return resultBuilder.String()
}

func InitRequest() (r *CreateContainerGroupRequest) {
	r = &CreateContainerGroupRequest{
		RpcRequest: &RpcRequest{},
	}
	r.InitWithApiInfo("Eci", "2018-08-08", "CreateContainerGroup", "eci", "openAPI")
	return
}

func Test_AcsRequest_InitParams2(t *testing.T) {
	r := InitRequest()
	InitParams(r)
	assert.Equal(t, "", GetQueryString(r))
}

func Test_AcsRequest_InitParams3(t *testing.T) {
	r := InitRequest()
	r.OwnerId = "OwnerId"
	InitParams(r)
	assert.Equal(t, "OwnerId=OwnerId&", GetQueryString(r))
}

func Test_AcsRequest_InitParams4(t *testing.T) {
	r := InitRequest()
	r.RegionId = "regionid"
	r.DnsConfig = CreateContainerGroupDnsConfig{
		NameServer: []string{"nameserver1", "nameserver2"},
	}
	InitParams(r)
	assert.Equal(t, "DnsConfig.NameServer.1=nameserver1&DnsConfig.NameServer.2=nameserver2&", GetQueryString(r))
}

func Test_AcsRequest_InitParams5(t *testing.T) {
	r := InitRequest()
	r.RegionId = "regionid"
	r.DnsConfig = CreateContainerGroupDnsConfig{
		Option: &[]CreateContainerGroupOption{
			{
				Name:  "sdk",
				Value: "test",
			},
		},
	}
	InitParams(r)
	assert.Equal(t, "DnsConfig.Option.1.Name=sdk&DnsConfig.Option.1.Value=test&", GetQueryString(r))
}

func Test_AcsRequest_InitParams6(t *testing.T) {
	r := InitRequest()
	r.Volume = &[]CreateContainerGroupVolume{
		{
			Name: "nginx",
			Type: "1",
			NFSVolume: CreateContainerGroupNFSVolume{
				Path: "/",
				Option: CreateContainerGroupOption{
					Name: "sdk",
				},
			},
		},
	}
	InitParams(r)
	assert.Equal(t, "Volume.1.NFSVolume.Option.Name=sdk&Volume.1.NFSVolume.Path=/&Volume.1.Name=nginx&Volume.1.Type=1&", GetQueryString(r))
}

type StartMPUTaskRequest struct {
	*RpcRequest
	OwnerId         Integer   `position:"Query" name:"OwnerId"`
	AppId           string    `position:"Query" name:"AppId"`
	ChannelId       string    `position:"Query" name:"ChannelId"`
	TaskId          string    `position:"Query" name:"TaskId"`
	MediaEncode     Integer   `position:"Query" name:"MediaEncode"`
	BackgroundColor Integer   `position:"Query" name:"BackgroundColor"`
	LayoutIds       []Integer `position:"Query" name:"LayoutIds" type:"Repeated"`
	StreamURL       string    `position:"Query" name:"StreamURL"`
}

func Test_RPCRequest_InitParams(t *testing.T) {
	channelID := "id"
	r := &StartMPUTaskRequest{
		RpcRequest: &RpcRequest{},
	}
	r.init()
	r.Domain = "rtc.aliyuncs.com"
	r.AppId = "app ID"
	r.ChannelId = channelID
	r.TaskId = channelID
	r.MediaEncode = NewInteger(2)
	r.BackgroundColor = NewInteger(0)
	r.StreamURL = fmt.Sprintf("rtmp://video-center.alivecdn.com/galaxy/%s_%s?vhost=fast-live.chinalivestream.top", channelID, channelID)
	var out []Integer
	out = append(out, NewInteger(2))
	r.LayoutIds = out

	InitParams(r)

	queries := r.GetQueryParams()

	assert.Equal(t, "2", queries["LayoutIds.1"])
	assert.Len(t, queries, 7)
}

type TestAcsRequestWithErrorMap struct {
	*RpcRequest
	OptionJson  map[string]interface{} `position:"Host" name:"OptionJson"`
	OptionJsonN map[string]interface{} `position:"Host" name:"OptionJsonN"`
}

func Test_InitParamsWithErrorMap(t *testing.T) {
	r := &TestAcsRequestWithErrorMap{
		RpcRequest: &RpcRequest{},
	}
	r.init()
	r.OptionJsonN = map[string]interface{}{
		"sdk": "test",
	}

	err := InitParams(r)
	assert.Equal(t, "[SDK.UnsupportedParamPosition] Specified param position (Host) is not supported, please upgrade sdk and retry", err.Error())
}

type TestAcsRequestWithErrorRepeated struct {
	*RpcRequest
	OptionJsonN []OptionJson `position:"Host" name:"OptionJsonN" type:"Repeated"`
}

type OptionJson struct {
	Option map[string]interface{} `name:"Option"`
}

func Test_InitParamsWithErrorRepeated(t *testing.T) {
	r := &TestAcsRequestWithErrorRepeated{
		RpcRequest: &RpcRequest{},
	}
	r.init()
	r.OptionJsonN = []OptionJson{
		{
			map[string]interface{}{
				"sdk": "test",
			},
		},
	}

	err := InitParams(r)
	assert.Equal(t, "[SDK.UnsupportedParamPosition] Specified param position (Host) is not supported, please upgrade sdk and retry", err.Error())
}
