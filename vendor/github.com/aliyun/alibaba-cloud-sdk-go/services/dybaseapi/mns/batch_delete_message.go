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
	"encoding/xml"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)


func (client *Queue) BatchDeleteMessage(request *BatchDeleteMessageRequest) (response *BatchDeleteMessageResponse, err error) {
	response = CreateBatchDeleteMessageResponse()
	err = client.DoActionWithSigner(request, response)
	return
}

type BatchDeleteMessageRequest struct {
	*requests.RoaRequest
	QueueName		string 			`position:"Path" name:"QueueName"`
}

type BatchDeleteMessageResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Code            string          `json:"Code" xml:"Code"`
}

type ReceiptHandles struct {
	XMLName xml.Name 				`xml:"ReceiptHandles"`
	Xmlns 	string   				`xml:"xmlns,attr"`
	Handles []string				`xml:"ReceiptHandle"`
}

func (request *BatchDeleteMessageRequest) SetReceiptHandles(receiptHandles []string) {
	receiptHandlesObj := &ReceiptHandles{
		Xmlns: "http://mns.aliyuncs.com/doc/v1/",
		Handles: receiptHandles,
	}
	content, err := xml.Marshal(receiptHandlesObj)
	if err != nil {
		panic(err)
	}
	request.SetContent([]byte(xml.Header + string(content)))
}

func CreateBatchDeleteMessageRequest() (request *BatchDeleteMessageRequest) {
	request = &BatchDeleteMessageRequest{
		RoaRequest: &requests.RoaRequest{
		},
	}
	request.InitWithApiInfo("MNS", "2015-06-06","BatchDeleteMessage","/queues/[QueueName]/messages", "", "")
	request.Method = "DELETE"
	request.Headers["x-mns-version"] = "2015-06-06"
	request.AcceptFormat = "XML"
	return
}

func CreateBatchDeleteMessageResponse() (response *BatchDeleteMessageResponse) {
	response = &BatchDeleteMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
