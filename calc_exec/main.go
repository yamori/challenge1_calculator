package main

import (
	"fmt"
    "strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

    "github.com/Knetic/govaluate"
)

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get the query parameters that was sent
    newOps := request.QueryStringParameters["newOps"]
    // '+' is reserved from query string
    newOps = strings.ReplaceAll(newOps, "plus", "+") 
    previousVal := request.QueryStringParameters["previousVal"]

    // blank idenfitifiers used as the 'err' goes unused, currenly
    expression, _ := govaluate.NewEvaluableExpression(previousVal+newOps);
    result, _ := expression.Evaluate(nil);

	//Generate message that want to be sent as body
	message := fmt.Sprintf(" { \"returnVal\" : %.3f } ", result)

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{
        Headers:    map[string]string{"content-type": "application/json"},
        Body: message, 
        StatusCode: 200,
    }, nil
}


func main() {
	lambda.Start(Handler)
}
