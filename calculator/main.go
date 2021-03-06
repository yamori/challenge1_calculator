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
<html>
    <head>
        <meta charset="UTF-8">
        <script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
        <script>
            function removeBadChars() {
                var inputString = $("#calc_inputs").val();
                // TBD, figure out how to regex restrict input
                //$("#calc_inputs").val("");
                //$("#calc_inputs").val( "h");
            }

            var previousVal = "";
            $(document).on('keypress',function(e) {
                if(e.which != 13) { return; } // Only on 'enter'

                var newOps = $("#calc_inputs").val();
                $("#calc_tape").val($("#calc_tape").val() + "\n  " + newOps);
                newOps = newOps.replace(/\+/g,"plus"); // because '+' is reserved from URL query strings

                $.ajax({
                    url: "/calc_exec?previousVal=" + previousVal + "&newOps=" + newOps,
                    type: "GET",
                    dataType: "json",
                    contentType: "charset=utf-8",
                    success: function (response, status, http) {
                        // (success, nothing to do))
                        previousVal = response.returnVal;
                        $("#calc_tape").val($("#calc_tape").val() + "\n\n            " + response.returnVal);
                    },
                    error: function (error) {
                        previousVal = 0;
                        $("#calc_tape").val($("#calc_tape").val() + "\nINPUT ERROR\n\n            0");
                    },
                    complete: function(data) {
                        $("#calc_inputs").val("");

                        // Scroll to bottom of 'tape'
                        var psconsole = $('#calc_tape');
                        if(psconsole.length) {
                            psconsole.scrollTop(psconsole[0].scrollHeight - psconsole.height());
                        }
                    }
                });
            });

            $( document ).ready(function() {
                $("#calc_inputs").val("");
                $("#calc_inputs").focus();
                $("#calc_tape").val("\n".repeat(32));
            });
        </script>
        <style>
        .bg-image {
            /* The image used */
            background-image: url("https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Background_3-8_01.43702ae9.jpg/1280px-Background_3-8_01.43702ae9.jpg");

            /* Add the blur effect */
            filter: blur(8px);
            -webkit-filter: blur(8px);

            /* Full height */
            height: 100%;

            /* Center and scale the image nicely */
            background-position: center;
            background-repeat: no-repeat;
            background-size: cover;
            z-index: -1;
        }
        #calc_tape {
			font-family: 'Courier New', Courier, monospace;
            position: fixed;
            left: 50%;
            bottom: 80px;
            transform: translate(-50%, 0);
            resize: none;
        }
        input {
            font-size: larger;
            font-family: 'Courier New', Courier, monospace;
            text-align: center; 
            width: 450px;
            position: fixed;
            left: 50%;
            bottom: 10px;
            transform: translate(-50%, -50%);
            margin: 0 auto;
        }
        #h4_label {
            position: fixed;
            left: 20px;
            top: 20px;
            color: white;
            -webkit-transform:translate3d(0,0,1px);
        }
        </style>
    </head>
    <body>
        <h4 id="h4_label">GoLang Server-side Calculator</h4>
        <div class="bg-image">
            </div>
            <textarea id="calc_tape" name="w3review" rows="32" cols="40"></textarea>
            <input id="calc_inputs" placeholder="calc commands here, then 'enter'" type="text" onkeyup="removeBadChars();">
        
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
