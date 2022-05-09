package utils

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// rr: 主机记录
func AddDomainRecord(client *alidns.Client, domain, rr, ip string) error {
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"

	request.Value = ip
	request.Type = "A"
	request.RR = rr
	request.DomainName = domain

	response, err := client.AddDomainRecord(request)
	if err != nil {
		return err
	}
	if response.BaseResponse.GetHttpStatus() != 200 {
		return fmt.Errorf("add record failed: %#v", response)
	}
	return nil
}

func UpdateDomainRecord(client *alidns.Client, recordID, rr, ip string) error {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = recordID
	request.RR = rr
	request.Type = "A"
	request.Value = ip

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		return err
	}
	if response.BaseResponse.GetHttpStatus() != 200 {
		return fmt.Errorf("add record failed: %#v", response)
	}
	return nil
}

func DescribeDomainRecords(client *alidns.Client, domain, rr string) ([]alidns.Record, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"

	request.DomainName = domain
	request.KeyWord = rr
	// 直接匹配
	request.SearchMode = "EXACT"
	request.Type = "A"

	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}
	if response.BaseResponse.GetHttpStatus() != 200 {
		return nil, fmt.Errorf("add record failed: %#v", response)
	}
	return response.DomainRecords.Record, nil
}
