package ess

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

func (client *Client) CreateScalingRule(request *CreateScalingRuleRequest) (response *CreateScalingRuleResponse, err error) {
	response = CreateCreateScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateScalingRuleWithChan(request *CreateScalingRuleRequest) (<-chan *CreateScalingRuleResponse, <-chan error) {
	responseChan := make(chan *CreateScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScalingRule(request)
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

func (client *Client) CreateScalingRuleWithCallback(request *CreateScalingRuleRequest, callback func(response *CreateScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateScalingRule(request)
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

type CreateScalingRuleRequest struct {
	*requests.RpcRequest
	ScalingRuleName      string           `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentType       string           `position:"Query" name:"AdjustmentType"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Cooldown             requests.Integer `position:"Query" name:"Cooldown"`
	AdjustmentValue      requests.Integer `position:"Query" name:"AdjustmentValue"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
}

type CreateScalingRuleResponse struct {
	*responses.BaseResponse
	ScalingRuleId  string `json:"ScalingRuleId" xml:"ScalingRuleId"`
	ScalingRuleAri string `json:"ScalingRuleAri" xml:"ScalingRuleAri"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateScalingRuleRequest() (request *CreateScalingRuleRequest) {
	request = &CreateScalingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingRule", "ess", "openAPI")
	return
}

func CreateCreateScalingRuleResponse() (response *CreateScalingRuleResponse) {
	response = &CreateScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
