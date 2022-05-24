package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/k8scat/auto-dns/utils"
)

const defaultRegion = "cn-hangzhou"

var (
	accessKey    string
	accessSecret string

	domain string
	rr     string
	ip     string
)

func initFlags() {
	flag.StringVar(&accessKey, "access-key", "", "access key")
	flag.StringVar(&accessSecret, "access-secret", "", "access secret")

	flag.StringVar(&domain, "domain", "", "domain")
	flag.StringVar(&rr, "rr", "", "rr")
	flag.StringVar(&ip, "ip", "", "ip")

	flag.Parse()
}

func checkArgs() error {
	if accessKey == "" || accessSecret == "" {
		return fmt.Errorf("access key or secret is empty")
	}
	if domain == "" || rr == "" || ip == "" {
		return fmt.Errorf("domain, rr or ip is empty")
	}
	return nil
}

func main() {
	initFlags()
	if err := checkArgs(); err != nil {
		panic(err)
	}

	client, err := alidns.NewClientWithAccessKey(defaultRegion, accessKey, accessSecret)
	if err != nil {
		panic(err)
	}

	records, err := utils.DescribeDomainRecords(client, domain, rr)
	if err != nil {
		panic(err)
	}

	if len(records) == 0 {
		err = utils.AddDomainRecord(client, domain, rr, ip)
	} else {
		fmt.Printf("exist record: %#v\n", records[0])
		err = utils.UpdateDomainRecord(client, records[0].RecordId, rr, ip)
	}
	if err != nil {
		panic(err)
	}
}
