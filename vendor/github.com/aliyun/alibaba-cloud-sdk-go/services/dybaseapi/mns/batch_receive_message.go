package mns

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)


func (client *Queue) BatchReceiveMessage(request *BatchReceiveMessageRequest) (response *BatchReceiveMessageResponse, err error) {
	response = CreateBatchReceiveMessageResponse()
	err = client.DoActionWithSigner(request, response)
	return
}

type BatchReceiveMessageRequest struct {
	*requests.RoaRequest
	QueueName			string 			 `position:"Path" name:"QueueName"`
	NumOfMessages		requests.Integer `position:"Query" name:"numOfMessages"`
	WaitSeconds			requests.Integer `position:"Query" name:"waitseconds"`
}

type BatchReceiveMessageResponse struct {
	*responses.BaseResponse
	RequestId       	string          `json:"RequestId" xml:"RequestId"`
	Code            	string          `json:"Code" xml:"Code"`
	Message         	[]Message 		`json:"Message" xml:"Message"`
}

func CreateBatchReceiveMessageRequest() (request *BatchReceiveMessageRequest) {
	request = &BatchReceiveMessageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("MNS", "2015-06-06","BatchReceiveMessage","/queues/[QueueName]/messages", "", "")
	request.Method = "GET"
	request.Headers["x-mns-version"] = "2015-06-06"
	request.AcceptFormat = "XML"
	return
}

func CreateBatchReceiveMessageResponse() (response *BatchReceiveMessageResponse) {
	response = &BatchReceiveMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
