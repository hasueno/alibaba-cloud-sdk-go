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

func (client *Client) DescribeScalingInstances(request *DescribeScalingInstancesRequest) (response *DescribeScalingInstancesResponse, err error) {
	response = CreateDescribeScalingInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeScalingInstancesWithChan(request *DescribeScalingInstancesRequest) (<-chan *DescribeScalingInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingInstances(request)
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

func (client *Client) DescribeScalingInstancesWithCallback(request *DescribeScalingInstancesRequest, callback func(response *DescribeScalingInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingInstances(request)
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

type DescribeScalingInstancesRequest struct {
	*requests.RpcRequest
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	CreationType           string           `position:"Query" name:"CreationType"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId10           string           `position:"Query" name:"InstanceId.10"`
	HealthStatus           string           `position:"Query" name:"HealthStatus"`
	InstanceId12           string           `position:"Query" name:"InstanceId.12"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId11           string           `position:"Query" name:"InstanceId.11"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	InstanceId19           string           `position:"Query" name:"InstanceId.19"`
	InstanceId17           string           `position:"Query" name:"InstanceId.17"`
	InstanceId18           string           `position:"Query" name:"InstanceId.18"`
	InstanceId15           string           `position:"Query" name:"InstanceId.15"`
	InstanceId6            string           `position:"Query" name:"InstanceId.6"`
	InstanceId16           string           `position:"Query" name:"InstanceId.16"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId7            string           `position:"Query" name:"InstanceId.7"`
	InstanceId13           string           `position:"Query" name:"InstanceId.13"`
	ScalingConfigurationId string           `position:"Query" name:"ScalingConfigurationId"`
	InstanceId8            string           `position:"Query" name:"InstanceId.8"`
	InstanceId14           string           `position:"Query" name:"InstanceId.14"`
	InstanceId9            string           `position:"Query" name:"InstanceId.9"`
	InstanceId2            string           `position:"Query" name:"InstanceId.2"`
	InstanceId3            string           `position:"Query" name:"InstanceId.3"`
	InstanceId4            string           `position:"Query" name:"InstanceId.4"`
	InstanceId5            string           `position:"Query" name:"InstanceId.5"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ScalingGroupId         string           `position:"Query" name:"ScalingGroupId"`
	InstanceId1            string           `position:"Query" name:"InstanceId.1"`
	LifecycleState         string           `position:"Query" name:"LifecycleState"`
	InstanceId20           string           `position:"Query" name:"InstanceId.20"`
}

type DescribeScalingInstancesResponse struct {
	*responses.BaseResponse
	TotalCount       int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageSize         int    `json:"PageSize" xml:"PageSize"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ScalingInstances struct {
		ScalingInstance []struct {
			InstanceId             string `json:"InstanceId" xml:"InstanceId"`
			ScalingConfigurationId string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
			ScalingGroupId         string `json:"ScalingGroupId" xml:"ScalingGroupId"`
			HealthStatus           string `json:"HealthStatus" xml:"HealthStatus"`
			LoadBalancerWeight     int    `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
			LifecycleState         string `json:"LifecycleState" xml:"LifecycleState"`
			CreationTime           string `json:"CreationTime" xml:"CreationTime"`
			CreationType           string `json:"CreationType" xml:"CreationType"`
		} `json:"ScalingInstance" xml:"ScalingInstance"`
	} `json:"ScalingInstances" xml:"ScalingInstances"`
}

func CreateDescribeScalingInstancesRequest() (request *DescribeScalingInstancesRequest) {
	request = &DescribeScalingInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingInstances", "ess", "openAPI")
	return
}

func CreateDescribeScalingInstancesResponse() (response *DescribeScalingInstancesResponse) {
	response = &DescribeScalingInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
