package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	securityGroupId = ""
	flag            = false
)

func Test_DescribeClusteWithROArequestWithXMLWithGet(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDescribeClusterDetailRequest()
	request.SetContentType("XML")
	request.SetScheme("HTTPS")
	response, err := client.DescribeClusterDetail(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_ScaleClusterWithROArequestWithXMLWithPUT(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateScaleClusterRequest()
	request.SetContentType("XML")
	request.SetScheme("HTTPS")
	response, err := client.ScaleCluster(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_CreateClusterTokenWithROArequestWithXMLWithPOST(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateCreateClusterRequest()
	request.SetContentType("XML")
	request.SetScheme("HTTPS")
	response, err := client.CreateCluster(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request body can't be empty")
}

func Test_DeleteClusterWithROArequestWithXMLWithDelete(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDeleteClusterRequest()
	request.SetContentType("XML")
	request.SetScheme("HTTPS")
	response, err := client.DeleteCluster(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_DeleteClusterWithROArequestWithJSONWithDelete(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDeleteClusterRequest()
	request.SetContentType("JSON")
	request.SetScheme("HTTPS")
	response, err := client.DeleteCluster(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_ScaleClusterWithROArequestWithJSONWithPUT(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateScaleClusterRequest()
	request.SetContentType("JSON")
	request.SetScheme("HTTPS")
	response, err := client.ScaleCluster(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_CreateSecurityGroupWithRPCrequestWithJSONWithNestingparametersWithPOST(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateCreateSecurityGroupRequest()
	request.SetContentType("JSON")
	tag := ecs.CreateSecurityGroupTag{
		Key:   "test",
		Value: "test",
	}
	request.Tag = &[]ecs.CreateSecurityGroupTag{tag}
	response, err := client.CreateSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.SecurityGroupId) > 0)
	securityGroupId = response.SecurityGroupId
}

func Test_ECS_DescribeSecurityGroupsWithRPCrequestWithJSONWithNestingparametersWithGET(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.SetContentType("JSON")
	request.Method = requests.GET
	response, err := client.DescribeSecurityGroups(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	for _, securitygroup := range response.SecurityGroups.SecurityGroup {
		if securitygroup.SecurityGroupId == securityGroupId {
			flag = true
			break
		}
	}
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, flag)
	flag = false

}

func Test_ECS_DeleteSecurityGroupWithRPCrequestWithJSONWithPOST(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDeleteSecurityGroupRequest()
	request.SetContentType("JSON")
	request.SecurityGroupId = securityGroupId
	response, err := client.DeleteSecurityGroup(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	securityGroupId = ""
}

func Test_RDS_DescribeDBInstancesWithRPCrequest(t *testing.T) {
	client, err := rds.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := rds.CreateDescribeDBInstancesRequest()
	request.SetContentType("JSON")
	response, err := client.DescribeDBInstances(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_CDN_DescribeCdnDomainDetailWithRPCrequest(t *testing.T) {
	client, err := cdn.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := cdn.CreateDescribeRefreshTasksRequest()
	response, err := client.DescribeRefreshTasks(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_RAM_ListRolesWithRPCrequest(t *testing.T) {
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ram.CreateListRolesRequest()
	request.Scheme = "HTTPS"
	response, err := client.ListRoles(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
}

func Test_SLB_DescribeRegionsWithRPCrequest(t *testing.T) {
	client, err := slb.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := slb.CreateDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_VPC_DescribeRegionsWithRPCrequest(t *testing.T) {
	client, err := vpc.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := vpc.CreateDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func mockServer(status int, json string) (server *httptest.Server) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(json))
		return
	}))
	return ts
}

func Test_DescribeRegionsWithRPCrequestWithunicode(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "HTTP"
	ts := mockServer(400, `{"Code": "&&&&杭州&&&"}`)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "", 1)
	request.Domain = domain
	response, err := client.DescribeRegions(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Equal(t, "{\"Code\": \"&&&&杭州&&&\"}", response.GetHttpContentString())
}

func Test_DescribeRegionsWithRPCrequestWithescape(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "HTTP"
	ts := mockServer(400, `{"Code": "\t"}`)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "", 1)
	request.Domain = domain
	response, err := client.DescribeRegions(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Equal(t, "{\"Code\": \"\\t\"}", response.GetHttpContentString())
}

func Test_DescribeRegionsWithRPCrequestWith3XX(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "HTTP"
	ts := mockServer(307, `{"error"}`)
	defer ts.Close()
	domain := strings.Replace(ts.URL, "http://", "", 1)
	request.Domain = domain
	response, err := client.DescribeRegions(request)
	assert.NotNil(t, err)
	assert.Equal(t, 307, response.GetHttpStatus())
	assert.Equal(t, "{\"error\"}", response.GetHttpContentString())
}

func Test_QueryAvaliableInstances(t *testing.T) {
	client, err := bssopenapi.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := bssopenapi.CreateQueryAvailableInstancesRequest()
	endpoints.AddEndpointMapping(os.Getenv("REGION_ID"), "BssOpenApi", "business.aliyuncs.com")
	response, err := client.QueryAvailableInstances(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
	assert.Equal(t, 36, len(response.RequestId))
}
