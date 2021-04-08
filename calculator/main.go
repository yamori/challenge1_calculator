package main

import (
	"bytes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"html/template"
)

// Edit this file with your html template.
// Add the variables you want to insert like this: {{ .Variable }}
var HtmlTemplate = `
<head>
<style>
body {
    background-color: linen;
}
h1 {
    color: maroon;
    margin-left: 40px;
} 
</style>
</head>
<html>
  <body>
    <h1>AAPL</h1>
    <h2>BANANA</h2>
	<h3>BERRY</h3>
  </body>
</html>
`

func BuildPage(htmlTemplate string) *bytes.Buffer {
    var bodyBuffer bytes.Buffer
    t := template.New("template")
    var templates = template.Must(t.Parse(htmlTemplate))
	var s struct{}
    templates.Execute(&bodyBuffer, s)
    return &bodyBuffer
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       BuildPage(HtmlTemplate).String(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
