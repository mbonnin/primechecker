package main

import (
    "fmt"
    "strconv"
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
)

//func main() {
//	n := os.Args[1]
//	number, err := strconv.ParseInt(n, 10, 64)
//	if err != nil {
//        panic(err)
//    }
//	fmt.Println(isPrime(number))
//} 

func main() {
	lambda.Start(HandleRequest)
}

type MyEvent struct {
        Number int `json:"number"`
}

func HandleRequest(ctx context.Context, number MyEvent) (events.APIGatewayProxyResponse, error) {

	returnVal := fmt.Sprintf("%s!", strconv.FormatBool(isPrime(number.Number)))

	return events.APIGatewayProxyResponse {
  		Body:       returnVal,
  		StatusCode: 200,
 	}, nil
}

func isPrime(num int) bool {

	n := int64(num)

	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n % 2 == 0 || n % 3 == 0 {
		return false
	}

	var i int64 = 5
	for i * i <= n {
		if n % i == 0 || n % i + 2 == 0 {
			return false
		}
		i = i + 6
	}
	return true
}