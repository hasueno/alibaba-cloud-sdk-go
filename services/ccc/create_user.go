package ccc

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

func (client *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	response = CreateCreateUserResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateUserWithChan(request *CreateUserRequest) (<-chan *CreateUserResponse, <-chan error) {
	responseChan := make(chan *CreateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUser(request)
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

func (client *Client) CreateUserWithCallback(request *CreateUserRequest, callback func(response *CreateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserResponse
		var err error
		defer close(result)
		response, err = client.CreateUser(request)
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

type CreateUserRequest struct {
	*requests.RpcRequest
	LoginName    string    `position:"Query" name:"LoginName"`
	Phone        string    `position:"Query" name:"Phone"`
	SkillLevel   *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
	RoleId       *[]string `position:"Query" name:"RoleId"  type:"Repeated"`
	Email        string    `position:"Query" name:"Email"`
	SkillGroupId *[]string `position:"Query" name:"SkillGroupId"  type:"Repeated"`
	InstanceId   string    `position:"Query" name:"InstanceId"`
	DisplayName  string    `position:"Query" name:"DisplayName"`
}

type CreateUserResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	UserId         string `json:"UserId" xml:"UserId"`
}

func CreateCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateUser", "", "")
	return
}

func CreateCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
