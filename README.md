# CodeChallenge1 - Calculator

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
