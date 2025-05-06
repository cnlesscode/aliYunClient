package ecs

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	_client "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cnlesscode/aliYunClient/configs"
)

type AliYunECSClient struct {
	Client *_client.Client
	Config configs.AliYunECSConfig
}

// 初始化阿里云客户端
func NewAliYunECSClient(config configs.AliYunECSConfig) (*AliYunECSClient, error) {
	configIN := &openapi.Config{
		AccessKeyId:     tea.String(config.AccessKeyId),
		AccessKeySecret: tea.String(config.AccessKeySecret),
	}
	configIN.Endpoint = tea.String(config.Endpoint)
	_result := &_client.Client{}
	var _err error
	_result, _err = _client.NewClient(configIN)
	if _err != nil {
		return nil, _err
	}
	return &AliYunECSClient{Client: _result, Config: config}, nil
}
