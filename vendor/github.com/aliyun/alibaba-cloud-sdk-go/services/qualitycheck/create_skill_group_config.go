package qualitycheck

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateSkillGroupConfig invokes the qualitycheck.CreateSkillGroupConfig API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/createskillgroupconfig.html
func (client *Client) CreateSkillGroupConfig(request *CreateSkillGroupConfigRequest) (response *CreateSkillGroupConfigResponse, err error) {
	response = CreateCreateSkillGroupConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSkillGroupConfigWithChan invokes the qualitycheck.CreateSkillGroupConfig API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createskillgroupconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSkillGroupConfigWithChan(request *CreateSkillGroupConfigRequest) (<-chan *CreateSkillGroupConfigResponse, <-chan error) {
	responseChan := make(chan *CreateSkillGroupConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSkillGroupConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateSkillGroupConfigWithCallback invokes the qualitycheck.CreateSkillGroupConfig API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createskillgroupconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSkillGroupConfigWithCallback(request *CreateSkillGroupConfigRequest, callback func(response *CreateSkillGroupConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSkillGroupConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateSkillGroupConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateSkillGroupConfigRequest is the request struct for api CreateSkillGroupConfig
type CreateSkillGroupConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// CreateSkillGroupConfigResponse is the response struct for api CreateSkillGroupConfig
type CreateSkillGroupConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateCreateSkillGroupConfigRequest creates a request to invoke CreateSkillGroupConfig API
func CreateCreateSkillGroupConfigRequest() (request *CreateSkillGroupConfigRequest) {
	request = &CreateSkillGroupConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "CreateSkillGroupConfig", "", "")
	return
}

// CreateCreateSkillGroupConfigResponse creates a response to parse from CreateSkillGroupConfig response
func CreateCreateSkillGroupConfigResponse() (response *CreateSkillGroupConfigResponse) {
	response = &CreateSkillGroupConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}