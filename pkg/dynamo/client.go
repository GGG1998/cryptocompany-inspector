package dynamo

import (
	"cryptocompany-inspector/pkg/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Client struct {
	Config  *config.AwsConfigurator
	session *session.Session
	svc     *dynamodb.DynamoDB
}

func NewClient(awsConfig *config.AwsConfigurator) (*Client, error) {
	sess, err := session.NewSession(awsConfig.Config)
	if err != nil {
		return nil, err
	}

	db := dynamodb.New(sess)

	return &Client{
		Config:  awsConfig,
		session: sess,
		svc:     db,
	}, nil
}
