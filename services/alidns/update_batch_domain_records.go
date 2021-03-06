package alidns

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

func (client *Client) UpdateBatchDomainRecords(request *UpdateBatchDomainRecordsRequest) (response *UpdateBatchDomainRecordsResponse, err error) {
	response = CreateUpdateBatchDomainRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateBatchDomainRecordsWithChan(request *UpdateBatchDomainRecordsRequest) (<-chan *UpdateBatchDomainRecordsResponse, <-chan error) {
	responseChan := make(chan *UpdateBatchDomainRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBatchDomainRecords(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) UpdateBatchDomainRecordsWithCallback(request *UpdateBatchDomainRecordsRequest, callback func(response *UpdateBatchDomainRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBatchDomainRecordsResponse
		var err error
		defer close(result)
		response, err = client.UpdateBatchDomainRecords(request)
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

type UpdateBatchDomainRecordsRequest struct {
	*requests.RpcRequest
	Records      string `position:"Query" name:"Records"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type UpdateBatchDomainRecordsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
}

func CreateUpdateBatchDomainRecordsRequest() (request *UpdateBatchDomainRecordsRequest) {
	request = &UpdateBatchDomainRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateBatchDomainRecords", "", "")
	return
}

func CreateUpdateBatchDomainRecordsResponse() (response *UpdateBatchDomainRecordsResponse) {
	response = &UpdateBatchDomainRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
