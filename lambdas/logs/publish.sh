#!/bin/bash

#
# thanks https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
#aws lambda create-function --function-name test --runtime go1.x --role  arn:aws:iam::933879187873:policy/service-role/AWSLambdaBasicExecutionRole-4e3bc5ea-295c-4544-a212-a079f15bf7d8 --handler logs --zip-file fileb://logs.zip
go build
zip -r logs.zip logs
aws lambda update-function-code --function-name skLog --zip-file fileb://logs.zip
aws lambda invoke --function-name skLog  --payload '{"pathParameters":{"childId":"teste"},"queryStringParameters":{"token":"ianianso290801"},"httpMethod":"GET"}' /tmp/skLogResult.json
cat /tmp/skLogResult.json
echo "...OK"