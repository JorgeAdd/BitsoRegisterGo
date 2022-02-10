package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

type Bitso struct {
	Success bool    `json:"success"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Id         string `json:"id"`
	High       string `json:"high"`
	Last       string `json:"last"`
	Book       string `json:"book"`
	Created_at string `json:"created_at"`
	Volume     string `json:"volume"`
	Vwap       string `json:"vwap"`
	Low        string `json:"low"`
	Ask        string `json:"ask"`
	Bid        string `json:"bid"`
	Change_24  string `json:"change_24"`
}

var lastLastChange = "0"

const TABLE_NAME = "BitsoRegisterDB"

func StartingGoRutine(dynamo *dynamodb.DynamoDB) {
	output1 := make(chan string)
	go GetBitsoInfo(dynamo, output1)
loop:
	for {
		select {
		case s1, ok := <-output1:
			if !ok {
				break loop
			}
			fmt.Println(s1)
		}
	}
}

func GetBitsoInfo(dynamo *dynamodb.DynamoDB, ch chan string) {

	for {
		response, err := http.Get("https://api.bitso.com/v3/ticker/?book=btc_mxn")
		if err != nil {
			panic(err)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		var responseObject Bitso
		json.Unmarshal(responseData, &responseObject)

		myuuid := uuid.NewV1().String()
		responseObject.Payload.Id = myuuid

		if responseObject.Success && lastLastChange != responseObject.Payload.Last {
			lastLastChange = responseObject.Payload.Last
			insertItemDynamo(dynamo, responseObject)
		}

		fmt.Println("-----------------------")
		time.Sleep(10 * time.Second)
		ch <- responseObject.Payload.Last
	}
}

func insertItemDynamo(dynamo *dynamodb.DynamoDB, reg Bitso) error {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"registerId": {
				S: aws.String(reg.Payload.Id),
			},
			"ask": {
				S: aws.String(reg.Payload.Ask),
			},
			"bid": {
				S: aws.String(reg.Payload.Bid),
			},
			"book": {
				S: aws.String(reg.Payload.Book),
			},
			"change_24": {
				S: aws.String(reg.Payload.Change_24),
			},
			"created_at": {
				S: aws.String(reg.Payload.Created_at),
			},
			"high": {
				S: aws.String(reg.Payload.High),
			},
			"last": {
				S: aws.String(reg.Payload.Last),
			},
			"low": {
				S: aws.String(reg.Payload.Low),
			},
			"success": {
				S: aws.String("true"),
			},
			"volume": {
				S: aws.String(reg.Payload.Volume),
			},
			"vwap": {
				S: aws.String(reg.Payload.Vwap),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})
	return err
}
