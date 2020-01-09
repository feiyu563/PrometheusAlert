package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cr"

	"github.com/stretchr/testify/assert"

	"fmt"
	"os"
	"testing"
)

var crTestKey = "crtestkey" + travisValue[len(travisValue)-1]

func Test_CR_CreateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.SetContentType("JSON")
	content := fmt.Sprintf(
		`{
			"Namespace":{
				"Namespace":"%s"
			}
		}`, crTestKey,
	)
	request.SetContent([]byte(content))

	response, err := client.CreateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_UpdateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateUpdateNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = crTestKey
	request.SetContentType("JSON")
	content := fmt.Sprintf(
		`{
			"Namespace":{
				"AutoCreate":%v,
				"DefaultVisibility":"%s"
			}
		}`, false, "PUBLIC",
	)
	request.SetContent([]byte(content))

	response, err := client.UpdateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_GetNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = crTestKey

	response, err := client.GetNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_GetNamespaceList(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceListRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)

	response, err := client.GetNamespaceList(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CR_DeleteNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateDeleteNamespaceRequest()
	domain := fmt.Sprintf("cr." + os.Getenv("REGION_ID") + ".aliyuncs.com")
	request.SetDomain(domain)
	request.Namespace = crTestKey

	response, err := client.DeleteNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
