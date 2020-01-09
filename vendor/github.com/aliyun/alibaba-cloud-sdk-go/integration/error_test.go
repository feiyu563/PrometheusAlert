package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"

	"os"
	"strings"
	"testing"
)

func Test_DescribeRegionsWithParameterError(t *testing.T) {
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

func Test_DescribeRegionsWithUnreachableError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("www.aliyun-hangzhou.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Equal(t, 0, response.GetHttpStatus())
	realerr := err.(errors.Error)
	assert.True(t, strings.Contains(realerr.OriginError().Error(), "www.aliyun-hangzhou.com"))
}

func Test_DescribeRegionsWithTimeout(t *testing.T) {
	credentail := credentials.NewAccessKeyCredential(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	config := &sdk.Config{
		Timeout: 100,
	}
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	request.SetDomain("ecs.aliyuncs.com")
	client, err := ecs.NewClientWithOptions(os.Getenv("REGION_ID"), config, credentail)
	response, err := client.DescribeRegions(request)
	assert.Equal(t, 0, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "https://ecs.aliyuncs.com")
	assert.Contains(t, err.Error(), "Client.Timeout exceeded while awaiting headers")
}

func Test_DescribeRegionsWithNilbody(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	ts := mockServer(400, ``)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "", 1)
	request.Domain = domain
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.NotNil(t, err)
}

func Test_DescribeRegionsWithFormatError(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	ts := mockServer(400, `bad json`)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "", 1)
	request.Domain = domain
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "bad json")
}
