# CodeChallenge1 - Calculator

Currently deployed [calculator](https://la2bc28voc.execute-api.us-east-1.amazonaws.com/calculator) (under construction).

## Things learned

- Using AWS [HTTP_api's](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vs-rest.html), and configuring GO handler to pick them up properly (see this [commit](https://github.com/yamori/challenge1_calculator/commit/51733e05e3ae3fe50297b320f168a9b450014944))

## Dependencies

- GO + modules
- `serverless` (`npm install -g serverless`)
- AWS credz (via `aws-azure-login`)

## Commands

```
make build

sls deploy # (will print the urls/endpoints)

sls invoke -f calculator

sls remove
```

## Links and Resources

- [serverless go/lambda template](https://www.serverless.com/framework/docs/providers/aws/examples/hello-world/go/)
