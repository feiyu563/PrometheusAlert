package integration

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func Test_DescribeRegionsWithRPCrequestWithAK(t *testing.T) {
	client, err := ecs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_DescribeRegionsWithRPCrequestWithSTStoken(t *testing.T) {
	assumeresponse, err := createAssumeRole()
	assert.Nil(t, err)
	credential := assumeresponse.Credentials
	client, err := ecs.NewClientWithStsToken(os.Getenv("REGION_ID"), credential.AccessKeyId, credential.AccessKeySecret, credential.SecurityToken)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 36, len(response.RequestId))
	assert.True(t, len(response.Regions.Region) > 0)
}

func Test_DescribeClusterDetailWithROArequestWithAK(t *testing.T) {
	client, err := cs.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)
	request := cs.CreateDescribeClusterDetailRequest()
	request.SetDomain("cs.aliyuncs.com")
	request.QueryParams["RegionId"] = os.Getenv("REGION_ID")
	request.Method = "GET"
	response, err := client.DescribeClusterDetail(request)
	assert.NotNil(t, err)
	assert.Equal(t, 400, response.GetHttpStatus())
	assert.Contains(t, err.Error(), "Request url is invalid")
}

func Test_DescribeRegionsWithRPCrequestWithArn(t *testing.T) {
	subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	assert.Nil(t, err)
	client, err := ecs.NewClientWithRamRoleArn(os.Getenv("REGION_ID"), subaccesskeyid, subaccesskeysecret, rolearn, "alice_test")
	assert.Nil(t, err)

	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	request.Domain = "ecs.aliyuncs.com"
	response, err := client.DescribeRegions(request)
	assert.Nil(t, err)
	assert.Equal(t, 36, len(response.RequestId))
}

func TestDescribeRegionsWithProviderAndAk(t *testing.T) {
	os.Setenv(provider.ENVAccessKeyID, os.Getenv("ACCESS_KEY_ID"))
	os.Setenv(provider.ENVAccessKeySecret, os.Getenv("ACCESS_KEY_SECRET"))
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithProvider(os.Getenv("REGION_ID"))
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func TestDescribeRegionsWithProviderAndRsaKeyPair(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2014-05-26"
	request.Product = "Ecs"
	request.ApiName = "DescribeRegions"
	request.SetDomain("ecs.ap-northeast-1.aliyuncs.com")
	request.TransToAcsRequest()

	key := os.Getenv("RSA_FILE_AES_KEY")

	srcfile, err := os.Open("./encyptfile")
	assert.Nil(t, err)
	defer srcfile.Close()

	buf := new(bytes.Buffer)
	read := bufio.NewReader(srcfile)
	read.WriteTo(buf)

	block, err := aes.NewCipher([]byte(key))
	assert.Nil(t, err)

	origData := buf.Bytes()
	blockdec := cipher.NewCBCDecrypter(block, []byte(key)[:block.BlockSize()])
	orig := make([]byte, len(origData))
	blockdec.CryptBlocks(orig, origData)
	orig = PKCS7UnPadding(orig)

	cyphbuf := bytes.NewBuffer(orig)
	scan := bufio.NewScanner(cyphbuf)
	var data string
	for scan.Scan() {
		if strings.HasPrefix(scan.Text(), "----") {
			continue
		}
		data += scan.Text() + "\n"
	}

	client, err := sdk.NewClientWithRsaKeyPair("ap-northeast-1", os.Getenv("PUBLIC_KEY_ID"), data, 3600)
	assert.Nil(t, err)

	response, err := client.ProcessCommonRequest(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func TestDescribeRegionsWithBearToken(t *testing.T) {
	request := requests.NewCommonRequest()
	request.Version = "2017-07-05"
	request.Product = "CCC"
	request.ApiName = "ListRoles "
	request.SetDomain("ccc.cn-shanghai.aliyuncs.com")
	request.TransToAcsRequest()
	client, err := sdk.NewClientWithBearerToken("cn-shanghai", "eyJhbGciOiJSUzI1NiIsImsyaWQiOiJlNE92NnVOUDhsMEY2RmVUMVhvek5wb1NBcVZLblNGRyIsImtpZCI6IkpDOXd4enJocUowZ3RhQ0V0MlFMVWZldkVVSXdsdEZodWk0TzFiaDY3dFUifQ.TjU2UldwZzFzRE1oVEN5UStjYlZLV1dzNW45cFBOSWdNRDhzQmVXYmVpLytWY012MEJqYjdTdnB3SE9LcHBiZkorUGdvclAxRy9GTjdHeldmaWZFVndoa05ueUNTem80dU0rUVFKdDFSY2V0bmFQcml5WFljTDhmNUZ2c1pFd3BhTDFOajVvRW9QVG83S1NVU3JpTFdKQmNnVHB1U094cUd4cGpCeFdXS0pDVnN0L3lzRkp4RTVlSFNzUm1Qa1FBVTVwS1lmaXE0QVFSd3lPQjdYSk1uUGFKU1BiSWhyWVFVS21WOVd5K2d3PT0.jxdCiNimyes3swDRBSxdsgaL4IlOD2Kz49Gf5w0VZ0Xap9ozUyxvSSywGzMrKvCTIoeh9QMCMjCpnt9A-nQxENj3YGAeBk8Wy19uHiT-4OVo-CiCKmKxILpzxcpOptNO-LER1swVLbt0NiTuTH4KB5CUaRwJKIFJuUwa57HcsWbvWQyZa1ms0NNOccNfGJl4177eY2LTUyyXWi4wYNA_L0YMTkZz4sOFM_Mdzks8bHXiSbGkkjfWQy0QblkLz6Bboh1OYlg3_RCLSWby_FMNoxU_eG2lGAsDnYxZDmCAq2jedY0x1RzZodo9HYRQN7DujlBhfzqm4hOBNvA3LiJfzw")
	assert.Nil(t, err)
	response, err := client.ProcessCommonRequest(request)
	assert.True(t, strings.Contains(err.Error(), "Bearertoken has expired"))
	assert.False(t, response.IsSuccess())
}
