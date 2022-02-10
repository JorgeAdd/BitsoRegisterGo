package controller

import (
	"fmt"

	"github.com/JorgeAdd/BitsoRegisterGo/BitsoRegisterRutine/internal/database"
	"github.com/JorgeAdd/BitsoRegisterGo/BitsoRegisterRutine/internal/service"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamo *dynamodb.DynamoDB

func GetBitsoInfoController() {
	dynamo = database.DynamoDbConnection()
	fmt.Println("dynamo running")

	service.StartingGoRutine(dynamo)

}
