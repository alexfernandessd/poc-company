package company

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Database map methods from database
type Database interface {
	get(companyTable string, companyID string, out interface{}) error
}

// DynamoDatabase map requirements to connection
type DynamoDatabase struct {
	db *dynamodb.DynamoDB
}

// NewDynamoDatabase will connect to an DynamoDB in a certain AWS Region.
func NewDynamoDatabase(region, endpoint, metadataURL string) (*DynamoDatabase, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	},
	)
	if err != nil {
		return nil, err
	}
	return &DynamoDatabase{db: dynamodb.New(sess)}, nil
}

// func createAwsConfig(region, endpoint, metadataURL string) *aws.Config {
// 	config := aws.NewConfig()
// 	if metadataURL != "" {
// 		config.WithCredentials(createAwsCredentials(metadataURL))
// 	}
// 	return config
// }

// func createAwsCredentials(metadataURL string) *credentials.Credentials {
// 	return credentials.NewChainCredentials(
// 		[]credentials.Provider{
// 			&credentials.EnvProvider{},
// 			&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
// 			&ec2rolecreds.EC2RoleProvider{
// 				Client: ec2metadata.New(session.New(&aws.Config{
// 					Endpoint: aws.String(metadataURL),
// 				})),
// 			},
// 		},
// 	)
// }

func (d DynamoDatabase) get(companyTable string, companyID string, out interface{}) error {
	result, err := d.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(companyTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(companyID),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if result.Item == nil {
		fmt.Println("Could not find the item")
		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, out)

	return err
}
