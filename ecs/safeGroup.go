package ecs

import (
	"errors"

	_client "github.com/alibabacloud-go/ecs-20140526/v7/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 获取安全组列表
// https://api.aliyun.com/api/Ecs/2014-05-26/DescribeSecurityGroups?lang=GO&params={%22RegionId%22:%22cn-beijing%22}&RegionId=cn-beijing
func (aliYunECSClient *AliYunECSClient) GetSecurityGroups() ([]*_client.DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup, error) {
	describeSecurityGroupsRequest := &_client.DescribeSecurityGroupsRequest{
		RegionId: tea.String(aliYunECSClient.Config.RegionId),
	}
	runtime := &util.RuntimeOptions{}
	res, tryErr := func() (_res *_client.DescribeSecurityGroupsResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_res, _err := aliYunECSClient.Client.DescribeSecurityGroupsWithOptions(describeSecurityGroupsRequest, runtime)
		if _err != nil {
			return nil, _err
		}
		return _res, nil
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return nil, error
	}
	if *res.StatusCode != 200 {
		return nil, errors.New("获取安全组列表失败")
	}
	groups := res.Body.SecurityGroups.SecurityGroup
	return groups, nil
}

// 查询某个安全组的具体规则
// https://api.aliyun.com/api/Ecs/2014-05-26/DescribeSecurityGroupAttribute
func (aliYunECSClient *AliYunECSClient) GetSecurityGroupRules(securityGroupId *string) ([]*_client.DescribeSecurityGroupAttributeResponseBodyPermissionsPermission, error) {
	describeSecurityGroupAttributeRequest := &_client.DescribeSecurityGroupAttributeRequest{
		RegionId:        tea.String(aliYunECSClient.Config.RegionId),
		SecurityGroupId: securityGroupId,
	}
	runtime := &util.RuntimeOptions{}
	permission := make([]*_client.DescribeSecurityGroupAttributeResponseBodyPermissionsPermission, 0)
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		res, _err := aliYunECSClient.Client.DescribeSecurityGroupAttributeWithOptions(describeSecurityGroupAttributeRequest, runtime)
		if _err != nil {
			return _err
		}
		permission = res.Body.Permissions.Permission
		return nil
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return nil, error
	}
	return permission, nil
}

// 修改安全组的具体规则
// https://api.aliyun.com/api/Ecs/2014-05-26/ModifySecurityGroupRule
func (aliYunECSClient *AliYunECSClient) ModifySecurityGroupRule(
	securityGroupId *string,
	rule *_client.DescribeSecurityGroupAttributeResponseBodyPermissionsPermission,
) (_err error) {
	modifySecurityGroupRuleRequest := &_client.ModifySecurityGroupRuleRequest{
		RegionId:            tea.String(aliYunECSClient.Config.RegionId),
		SourceCidrIp:        rule.SourceCidrIp,
		SecurityGroupRuleId: rule.SecurityGroupRuleId,
		SecurityGroupId:     securityGroupId,
		Description:         rule.Description,
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, err := aliYunECSClient.Client.ModifySecurityGroupRuleWithOptions(modifySecurityGroupRuleRequest, runtime)
		return err
	}()
	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		return error
	}
	return nil
}

// 添加安全组入方向规则
// https://api.aliyun.com/api/Ecs/2014-05-26/AuthorizeSecurityGroup?RegionId=cn-beijing&tab=DEMO&lang=GO
func (aliYunECSClient *AliYunECSClient) AddRule(securityGroupId *string, item *_client.AuthorizeSecurityGroupRequest) error {
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err := aliYunECSClient.Client.AuthorizeSecurityGroupWithOptions(item, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()
	if tryErr != nil {
		return tryErr
	}
	return nil
}

// 删除安全组入方向规则
// https://api.aliyun.com/api/Ecs/2014-05-26/RevokeSecurityGroup
func (aliYunECSClient *AliYunECSClient) RemoveSecurityGroupRule(securityGroupId *string, securityGroupRuleId *string) (_err error) {
	revokeSecurityGroupRequest := &_client.RevokeSecurityGroupRequest{
		RegionId:            tea.String(aliYunECSClient.Config.RegionId),
		SecurityGroupId:     securityGroupId,
		SecurityGroupRuleId: []*string{securityGroupRuleId},
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = aliYunECSClient.Client.RevokeSecurityGroupWithOptions(revokeSecurityGroupRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()
	return tryErr
}
