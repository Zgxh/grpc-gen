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
    protoc --proto_path=proto --go_out=proto/gen/go --java_out=proto/gen/java/src/main/java --go-grpc_out=proto/gen/go --grpc-java_out=proto/gen/java/src/main/java $file
    protoc $file
done

echo "deploy the java code to Nexus..."
cd ./proto/gen/java
mvn clean deploy

# echo "service is starting..."
# go mod tidy
# go run ./server/main.go