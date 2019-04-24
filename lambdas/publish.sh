#!/bin/bash

#
# thanks https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
#aws lambda create-function --function-name test --runtime go1.x --role  arn:aws:iam::933879187873:policy/service-role/AWSLambdaBasicExecutionRole-4e3bc5ea-295c-4544-a212-a079f15bf7d8 --handler logs --zip-file fileb://logs.zip


declare -A lambdas=( ["logs"]="skLog" ["children"]="GetChildren" ["family"]="skFamily")
name=$1
funcName="${lambdas[$name]}"
go build -o $name/$name $name/*.go && \
zip -r $name/$name.zip -j $name/$name && \
aws lambda update-function-code --function-name "$funcName" --zip-file fileb://$name/$name.zip && \
./publish_"$name"_test.sh "$funcName"
echo "...OK"