package rds

import (
	"errors"

	rdsClient "github.com/alibabacloud-go/rds-20140815/v12/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 查询白名单模板列表
// https://api.aliyun.com/api/Rds/2014-08-15/DescribeAllWhitelistTemplate?params={%22RegionId%22:%22cn-beijing%22}
func (client *AliYunRDSClient) GetWhitelistTemplates() ([]*rdsClient.DescribeAllWhitelistTemplateResponseBodyDataTemplates, error) {
	describeAllWhitelistTemplateRequest := &rdsClient.DescribeAllWhitelistTemplateRequest{
		RegionId:          tea.String(client.Config.RegionId),
		MaxRecordsPerPage: tea.Int32(50),
		PageNumbers:       tea.Int32(1),
	}
	runtime := &util.RuntimeOptions{}
	_result, tryErr := func() (_result *rdsClient.DescribeAllWhitelistTemplateResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		return client.Client.DescribeAllWhitelistTemplateWithOptions(describeAllWhitelistTemplateRequest, runtime)
	}()

	if tryErr != nil {
		return nil, tryErr
	}
	if *_result.StatusCode != 200 {
		return nil, errors.New("获取白名单模板列表失败")
	} else {
		return _result.Body.Data.Templates, nil
	}
}

// 查询白名单模板信息
// https://api.aliyun.com/api/Rds/2014-08-15/DescribeWhitelistTemplate
func (client *AliYunRDSClient) GetWhitelistTemplateInfo(templateId *int32) (*rdsClient.DescribeWhitelistTemplateResponseBodyDataTemplate, error) {
	describeWhitelistTemplateRequest := &rdsClient.DescribeWhitelistTemplateRequest{
		RegionId:   client.Client.RegionId,
		TemplateId: templateId,
	}
	runtime := &util.RuntimeOptions{}
	result, tryErr := func() (_result *rdsClient.DescribeWhitelistTemplateResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		return client.Client.DescribeWhitelistTemplateWithOptions(describeWhitelistTemplateRequest, runtime)
	}()
	if tryErr != nil {
		return nil, tryErr
	}
	return result.Body.Data.Template, nil
}

// 修改白名单模板
// https://api.aliyun.com/api/Rds/2014-08-15/ModifyWhitelistTemplate?params={%22RegionId%22:%22cn-beijing%22}
func (client *AliYunRDSClient) ModifyWhitelistTemplate(template *rdsClient.DescribeAllWhitelistTemplateResponseBodyDataTemplates) error {
	modifyWhitelistTemplateRequest := &rdsClient.ModifyWhitelistTemplateRequest{
		RegionId:     client.Client.RegionId,
		TemplateId:   template.TemplateId,
		TemplateName: template.TemplateName,
		IpWhitelist:  template.Ips,
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err := client.Client.ModifyWhitelistTemplateWithOptions(modifyWhitelistTemplateRequest, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()
	return tryErr
}
