package domain_intl

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

func (client *Client) VerifyEmail(request *VerifyEmailRequest) (response *VerifyEmailResponse, err error) {
	response = CreateVerifyEmailResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) VerifyEmailWithChan(request *VerifyEmailRequest) (<-chan *VerifyEmailResponse, <-chan error) {
	responseChan := make(chan *VerifyEmailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyEmail(request)
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

func (client *Client) VerifyEmailWithCallback(request *VerifyEmailRequest, callback func(response *VerifyEmailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyEmailResponse
		var err error
		defer close(result)
		response, err = client.VerifyEmail(request)
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

type VerifyEmailRequest struct {
	*requests.RpcRequest
	Token        string `position:"Query" name:"Token"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type VerifyEmailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateVerifyEmailRequest() (request *VerifyEmailRequest) {
	request = &VerifyEmailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "VerifyEmail", "", "")
	return
}

func CreateVerifyEmailResponse() (response *VerifyEmailResponse) {
	response = &VerifyEmailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}