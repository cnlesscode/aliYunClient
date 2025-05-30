package aliYunClient

import (
	"fmt"
	"log"
	"testing"

	_client "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cnlesscode/aliYunClient/configs"
	"github.com/cnlesscode/aliYunClient/ecs"
	"github.com/cnlesscode/gotool"
)

// 客户创建测试
var config = configs.AliYunECSConfig{
	Endpoint:        "ecs.cn-beijing.aliyuncs.com",
	AccessKeyId:     "***",
	AccessKeySecret: "***",
	RegionId:        "cn-beijing",
}

// 测试创建客户端
// go test -v -run NewAliYunECSClient
func TestNewAliYunECSClient(t *testing.T) {
	client, err := ecs.NewAliYunECSClient(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("aliYunECSClient: %v\n", client)
}

// 获取所有的安全组
// go test -v -run TestGetSecurityGroups
func TestGetSecurityGroups(t *testing.T) {
	client, err := ecs.NewAliYunECSClient(config)
	if err != nil {
		panic(err)
	}
	groups, err := client.GetSecurityGroups()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("groups: %v\n", groups)
}

// 添加一条安全组规则
// go test -v -run TestAddRule
func TestAddRule(t *testing.T) {
	client, err := ecs.NewAliYunECSClient(config)
	if err != nil {
		panic(err)
	}
	// 查询所有安全组
	groups, err := client.GetSecurityGroups()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 在第一个安全组内添加入规则
	rule := &_client.AuthorizeSecurityGroupRequest{
		RegionId:        tea.String(config.RegionId),
		SecurityGroupId: groups[0].SecurityGroupId,
		PortRange:       tea.String("3306/3306"),
		SourceCidrIp:    tea.String(gotool.GetNetworkIP()),
		IpProtocol:      tea.String("tcp"),
		Policy:          tea.String("accept"),
		Description:     tea.String("mySQL"),
	}
	err = client.AddRule(groups[0].SecurityGroupId, rule)
	fmt.Printf("err: %v\n", err)
}

// 查询第一个安全组下的所有规则并修改指定端口
// go test -v -run TestModifySecurityIpByPortRange
func TestModifySecurityIpByPortRange(t *testing.T) {
	client, err := ecs.NewAliYunECSClient(config)
	if err != nil {
		panic(err)
	}
	groups, err := client.GetSecurityGroups()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	rules, err := client.GetSecurityGroupRules(groups[0].SecurityGroupId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 遍历规则，检查是否存在 3306 端口的规则
	done := false
	for _, rule := range rules {
		if *rule.PortRange == "3306/3306" {
			// 修改 3306 端口的访问权限
			rule.SourceCidrIp = tea.String(gotool.GetNetworkIP())
			rule.Description = tea.String("mySQL ...")
			err = client.ModifySecurityGroupRule(groups[0].SecurityGroupId, rule)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			done = true
			return
		}
	}
	// 如果不存在 3306 端口的规则，则添加
	if !done {
		log.Println("不存在 3306 端口的规则，执行添加")
		rule := &_client.AuthorizeSecurityGroupRequest{
			RegionId:        tea.String(config.RegionId),
			SecurityGroupId: groups[0].SecurityGroupId,
			PortRange:       tea.String("3306/3306"),
			SourceCidrIp:    tea.String(gotool.GetNetworkIP()),
			IpProtocol:      tea.String("tcp"),
			Policy:          tea.String("accept"),
			Description:     tea.String("mySQL"),
		}
		err = client.AddRule(groups[0].SecurityGroupId, rule)
		fmt.Printf("err: %v\n", err)
	}
}

// 删除一个安全组规则
// go test -v -run TestRemoveSecurityRule
func TestRemoveSecurityRule(t *testing.T) {
	client, err := ecs.NewAliYunECSClient(config)
	if err != nil {
		panic(err)
	}
	groups, err := client.GetSecurityGroups()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	rules, err := client.GetSecurityGroupRules(groups[0].SecurityGroupId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 遍历规则，检查是否存在 3306 端口的规则
	for _, rule := range rules {
		if *rule.PortRange == "3306/3306" {
			err = client.RemoveSecurityGroupRule(groups[0].SecurityGroupId, rule.SecurityGroupRuleId)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			return
		}
	}
}
