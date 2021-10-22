#!/bin/bash

echo "start compiling..."

echo "create and clear the dir..."
if [ ! -d proto/gen/go ]; then 
    mkdir proto/gen/go/grpc
fi
if [ ! -d proto/gen/java ]; then 
    mkdir proto/gen/java/src/main/java
fi

rm -rf proto/gen/go/grpc/*
rm -rf proto/gen/java/src/main/java/*

echo "code generated..."
for file in `find . -name "*.proto"`;
do
    protoc --proto_path=proto \
    --go_out=proto/gen/go \
    --java_out=proto/gen/java/src/main/java \
    --go-grpc_out=proto/gen/go \
    --grpc-java_out=proto/gen/java/src/main/java \
    --grpc-gateway_out=proto/gen/go \
    $file
    protoc $file
done

echo "deploy the java code to Nexus..."
cd ./proto/gen/java
mvn clean deploy

echo "commit and push the src to github..."
cd ../../../
git add .
git commit -m "update"
git push

# echo "service is starting..."
# go mod tidy
# go run ./server/main.go