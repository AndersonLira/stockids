#!/bin/bash

#
# thanks https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
#aws lambda create-function --function-name test --runtime go1.x --role  arn:aws:iam::933879187873:policy/service-role/AWSLambdaBasicExecutionRole-4e3bc5ea-295c-4544-a212-a079f15bf7d8 --handler logs --zip-file fileb://logs.zip

declare -A lambdas=( ["logs"]="skLog" ["children"]="GetChildren" ["family"]="skFamily")
go build -o $1/$1 $1/*.go && \
zip -r $1/$1.zip -j $1/$1 && \
aws lambda update-function-code --function-name "${lambdas[$1]}" --zip-file fileb://$1/$1.zip && \
aws lambda invoke --function-name "${lambdas[$1]}"  --payload '{"pathParameters":{"childId":"teste"},"queryStringParameters":{"token":"ianianso290801"},"httpMethod":"GET"}' /tmp/$1result.json && \
cat /tmp/$1result.json && \
echo "...OK"