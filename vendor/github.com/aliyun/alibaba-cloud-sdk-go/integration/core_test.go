package integration

import (
	"fmt"
	"github.com/goji/httpauth"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/stretchr/testify/assert"
)

func Test_DescribeRegionsWithCommonRequestWithRPC(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithSTStoken(t *testing.T) {
	assumeresponse, err := createAssumeRole()
	assert.Nil(t, err)
	credential := assumeresponse.Credentials
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithStsToken(os.Getenv("REGION_ID"), credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithHTTPS(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	request.SetScheme("HTTPS")
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithUnicodeSpecificParams(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	request.SetContent([]byte("sdk&-杭&&&州-test"))
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_DescribeRegionsWithCommonRequestWithError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "Describe"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	_, err = client.ProcessCommonRequest(request)
	realerr := err.(errors.Error)
	assert.Equal(t, "InvalidParameter", realerr.ErrorCode())
	assert.Equal(t, "The specified parameter \"Action or Version\" is not valid.", realerr.Message())
}

func Test_DescribeRegionsWithCommonRequestWithIncompleteSignature(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.AcceptFormat = "json"
	request.SetScheme("https")
	request.Method = "POST"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), strings.ToUpper(os.Getenv("ACCESS_KEY_SECRET")))
	assert.Nil(t, err)
	_, err = client.ProcessCommonRequest(request)
	realerr := err.(*errors.ServerError)
	assert.Equal(t, "IncompleteSignature", realerr.ErrorCode())
	assert.Equal(t, "InvalidAccessKeySecret: Please check you AccessKeySecret", realerr.Recommend())
}

func Test_DescribeClustersWithCommonRequestWithROA(t *testing.T) {
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters"
	request.ApiName = "DescribeClusters"
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.TransToAcsRequest()
	resp, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetHttpStatus())

}

func Test_DescribeClustersWithCommonRequestWithSignatureDostNotMatch(t *testing.T) {
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), strings.ToUpper(os.Getenv("ACCESS_KEY_SECRET")))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/[ClusterId]"
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.TransToAcsRequest()
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	real, _ := err.(*errors.ServerError)
	assert.Contains(t, real.Recommend(), "InvalidAccessKeySecret: Please check you AccessKeySecret")
	assert.Equal(t, real.ErrorCode(), "SignatureDoesNotMatch")
}

func Test_DescribeClustersWithCommonRequestWithROAWithSTStoken(t *testing.T) {
	assumeresponse, err := createAssumeRole()
	assert.Nil(t, err)
	credential := assumeresponse.Credentials
	client, err := sdk.NewClientWithStsToken(os.Getenv("REGION_ID"), credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/clusters/[ClusterId]"
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.TransToAcsRequest()
	f1, err := os.Create("test.txt")
	defer os.Remove("test.txt")
	assert.Nil(t, err)
	templete := `{version}, {host}`
	client.SetLogger("error", "Alibaba", f1, templete)
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, client.GetLoggerMsg(), `1.1, cs.aliyuncs.com`)
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_DescribeClusterDetailWithCommonRequestWithROAWithHTTPS(t *testing.T) {
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.SetScheme("HTTPS")
	request.PathPattern = "/clusters/[ClusterId]"
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.TransToAcsRequest()

	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_DescribeClusterDetailWithCommonRequestWithTimeout(t *testing.T) {
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.SetScheme("HTTPS")
	request.PathPattern = "/clusters/[ClusterId]"
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.ReadTimeout = 1 * time.Millisecond
	request.ConnectTimeout = 1 * time.Nanosecond
	request.TransToAcsRequest()
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)

	request.ConnectTimeout = 1 * time.Second
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Read timeout. Please set a valid ReadTimeout.")
}

func Test_CreateInstanceWithCommonRequestWithPolicy(t *testing.T) {
	err := createAttachPolicyToRole()
	assert.Nil(t, err)

	subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	assert.Nil(t, err)
	client, err := sdk.NewClientWithRamRoleArnAndPolicy(os.Getenv("REGION_ID"), subaccesskeyid, subaccesskeysecret, rolearn, "alice_test", "")
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "Ecs"
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.SetScheme("HTTPS")
	request.ApiName = "CreateInstance"
	request.QueryParams["ImageId"] = "win2008r2_64_ent_sp1_en-us_40G_alibase_20170915.vhd"
	request.QueryParams["InstanceType"] = "ecs.g5.large"
	request.TransToAcsRequest()
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "user order resource type [classic] not exists in [random]")

	policy := `{
    "Version": "1",
    "Statement": [
        {
            "Action": "rds:*",
            "Resource": "*",
            "Effect": "Allow"
        },
        {
            "Action": "dms:LoginDatabase",
            "Resource": "acs:rds:*:*:*",
            "Effect": "Allow"
        }
    ]
}`
	client, err = sdk.NewClientWithRamRoleArnAndPolicy(os.Getenv("REGION_ID"), subaccesskeyid, subaccesskeysecret, rolearn, "alice_test", policy)
	assert.Nil(t, err)
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "User not authorized to operate on the specified resource, or this API doesn't support RAM.")
}

func handlerTrue(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("test"))
	return
}

func handlerFake(w http.ResponseWriter, r *http.Request) {
	trueserver := handlerTrueServer()
	url, err := url.Parse(trueserver.URL)
	if err != nil {
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	w.Write([]byte("sdk"))
	proxy.ServeHTTP(w, r)

	return
}

func handlerFakeServer() (server *httptest.Server) {
	handleFunc := httpauth.SimpleBasicAuth("someuser", "somepassword")(http.HandlerFunc(handlerFake))
	server = httptest.NewServer(handleFunc)

	return server
}

func handlerTrueServer() (server *httptest.Server) {
	server = httptest.NewServer(http.HandlerFunc(handlerTrue))

	return server
}

func Test_HTTPProxy(t *testing.T) {

	ts := handlerFakeServer()
	ts1 := handlerTrueServer()
	defer func() {
		ts.Close()
		ts1.Close()
	}()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	domain := strings.Replace(ts1.URL, "http://", "", 1)
	request.Domain = domain
	request.Version = "2015-12-15"
	request.TransToAcsRequest()
	resp, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetHttpStatus())
	assert.Equal(t, "test", resp.GetHttpContentString())

	originEnv := os.Getenv("HTTP_PROXY")
	domain = strings.Replace(ts.URL, "http://", "", 1)
	os.Setenv("HTTP_PROXY", fmt.Sprintf("http://someuser:somepassword@%s", domain))
	resp, err = client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetHttpStatus())
	assert.Equal(t, "sdktest", resp.GetHttpContentString())

	os.Setenv("HTTP_PROXY", originEnv)
}

func Test_DdoscooWithServiceCode(t *testing.T) {
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Version = "2017-12-28"
	request.Product = "ddoscoo"
	request.ServiceCode = "ddoscoo"
	request.ApiName = "DescribeInstanceSpecs"
	request.RegionId = "cn-hangzhou"
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "InstanceIds is mandatory for this action.")
}

func Test_RoaRequestWithEcsRole(t *testing.T) {
	client, err := sdk.NewClientWithEcsRamRole("cn-shenzhen", "test-go-role")
	assert.Nil(t, err)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "edas.cn-hangzhou.aliyuncs.com"
	request.Version = "2017-08-01"
	request.PathPattern = "/pop/v5/resource/region_list"

	request.QueryParams["RegionId"] = "cn-shenzhen"
	_, err = client.ProcessCommonRequest(request)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "refresh Ecs sts token err")
}
