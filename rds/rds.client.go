package rds

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	rdsClient "github.com/alibabacloud-go/rds-20140815/v12/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cnlesscode/aliYunClient/configs"
)

type AliYunRDSClient struct {
	Client *rdsClient.Client
	Config configs.AliYunRDSConfig
}

// 初始化阿里云客户端
func NewAliYunRDSClient(config configs.AliYunRDSConfig) (*AliYunRDSClient, error) {
	configIN := &openapi.Config{
		AccessKeyId:     tea.String(config.AccessKeyId),
		AccessKeySecret: tea.String(config.AccessKeySecret),
	}
	configIN.Endpoint = tea.String(config.Endpoint)
	_client := &rdsClient.Client{}
	var _err error
	_client, _err = rdsClient.NewClient(configIN)
	if _err != nil {
		return nil, _err
	}
	return &AliYunRDSClient{Client: _client, Config: config}, nil
}
