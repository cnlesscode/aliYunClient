package aliYunClient

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/cnlesscode/aliYunClient/configs"
	"github.com/cnlesscode/aliYunClient/rds"
	"github.com/cnlesscode/gotool"
)

// 客户创建测试
var rdsConfig = configs.AliYunRDSConfig{
	Endpoint:        "rds.aliyuncs.com",
	AccessKeyId:     "***",
	AccessKeySecret: "***",
	RegionId:        "cn-beijing",
}

// 获取白名单列表
// go test -v -run TestGetWhitelistTemplates
func TestGetWhitelistTemplates(t *testing.T) {
	client, err := rds.NewAliYunRDSClient(rdsConfig)
	if err != nil {
		panic(err)
	}
	templates, err := client.GetWhitelistTemplates()
	if err != nil {
		panic(err)
	}
	fmt.Printf("templates: %v\n", templates)
}

// 获取白名单详情
// go test -v -run TestGetWhitelistTemplateInfo
func TestGetWhitelistTemplateInfo(t *testing.T) {
	client, err := rds.NewAliYunRDSClient(rdsConfig)
	if err != nil {
		panic(err)
	}
	templates, err := client.GetWhitelistTemplates()
	if err != nil {
		panic(err)
	}
	// 获取第一个模板的详情
	template, err := client.GetWhitelistTemplateInfo(templates[0].TemplateId)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("template: %v\n", template)
}

// 修改白名单模板
// go test -v -run TestModifyWhitelistTemplate
func TestModifyWhitelistTemplate(t *testing.T) {
	client, err := rds.NewAliYunRDSClient(rdsConfig)
	if err != nil {
		panic(err)
	}
	templates, err := client.GetWhitelistTemplates()
	if err != nil {
		panic(err)
	}
	// 修改第一个白名单模板
	template := templates[0]
	// 由小写字母、数字、下划线（_）组成，字母开头，字母或数字结尾，长度为2-64个字符
	template.TemplateName = tea.String("white_list_template")
	// 规划 ip
	ips := strings.Split(*template.Ips, ",")
	if len(ips) < 5 {
		ips = append(ips, gotool.GetNetworkIP())
	} else {
		ips[4] = gotool.GetNetworkIP()
	}
	template.Ips = tea.String(strings.Join(ips, ","))
	err = client.ModifyWhitelistTemplate(template)
	fmt.Printf("err: %v\n", err)
}
