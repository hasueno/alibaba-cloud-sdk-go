package rds

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

func (client *Client) DescribeSQLLogFiles(request *DescribeSQLLogFilesRequest) (response *DescribeSQLLogFilesResponse, err error) {
	response = CreateDescribeSQLLogFilesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSQLLogFilesWithChan(request *DescribeSQLLogFilesRequest) (<-chan *DescribeSQLLogFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogFiles(request)
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

func (client *Client) DescribeSQLLogFilesWithCallback(request *DescribeSQLLogFilesRequest, callback func(response *DescribeSQLLogFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogFiles(request)
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

type DescribeSQLLogFilesRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	FileName             string           `position:"Query" name:"FileName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeSQLLogFilesResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		LogFile []struct {
			FileID         string `json:"FileID" xml:"FileID"`
			LogStatus      string `json:"LogStatus" xml:"LogStatus"`
			LogDownloadURL string `json:"LogDownloadURL" xml:"LogDownloadURL"`
			LogSize        string `json:"LogSize" xml:"LogSize"`
			LogStartTime   string `json:"LogStartTime" xml:"LogStartTime"`
			LogEndTime     string `json:"LogEndTime" xml:"LogEndTime"`
		} `json:"LogFile" xml:"LogFile"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeSQLLogFilesRequest() (request *DescribeSQLLogFilesRequest) {
	request = &DescribeSQLLogFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogFiles", "rds", "openAPI")
	return
}

func CreateDescribeSQLLogFilesResponse() (response *DescribeSQLLogFilesResponse) {
	response = &DescribeSQLLogFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}