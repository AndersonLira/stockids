#test family lambda
echo Testing $1 lambda

echo GET 
aws lambda invoke --function-name "$1"  --payload '{"pathParameters":{"childId":"teste"},"queryStringParameters":{"token":"ianianso290801"},"httpMethod":"GET"}' /tmp/$1_get.json && \
cat /tmp/$1_get.json
echo

echo POST
aws lambda invoke --function-name "$1"  --payload '{"pathParameters":{"childId":"teste"},"queryStringParameters":{"token":"ianianso290801"},"httpMethod":"POST","body":"{\"id\":\"\",\"name\":\"Family A\",\"description\":\"A test family\",\"avatar\":\"none\"}"}' /tmp/$1_post.json && \
created=$(</tmp/$1_post.json)
echo $created
echo

echo DELETE
aws lambda invoke --function-name "$1"  --payload '{"pathParameters":{"id":"0b74c95f-7357-43e0-a1b7-90b6d0e9978a"},"queryStringParameters":{"token":"ianianso290801"},"httpMethod":"DELETE"}' /tmp/$1_delete.json && \
cat /tmp/$1_delete.json

echo
