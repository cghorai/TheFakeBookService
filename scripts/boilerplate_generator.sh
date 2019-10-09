#!/usr/bin/env bash
project_name=fakebookservice #Make sure this is same as your proto file name
grpc=./api/proto/v1 #Change this if you have placed your protobuf file in a different directory
protoFile=${grpc}/${project_name}.proto #Your proto file
pbFile=${grpc}/${project_name}.pb.go #Your generated pb.go file
pbDest=./pkg/server #Where your pb and gateway go file be copied to

echo "Generating ${project_name}.pb.go"
protoc -I/usr/local/include -I. \
                    -I$GOPATH/src \
                    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
                    --go_out=plugins=grpc:. \
                    $protoFile

echo "Adding bson tag. This is required because we will use MongoDb"
my_array=(`structnames --input ${pbFile} --output ./output.txt`)
my_array_length=${#my_array[@]}
for element in "${my_array[@]}"
do
   gomodifytags -file ${pbFile} -struct "${element}" -add-tags bson -transform camelcase > temp.txt && mv temp.txt ${pbFile}
done
rm ./output.txt

echo "Generating ${project_name}.pb.gw.go"
protoc -I/usr/local/include -I. \
                   -I$GOPATH/src \
                   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
                   --grpc-gateway_out=logtostderr=true:. \
                   $protoFile

echo "Copying the generated files in correct location"
mv $grpc/${project_name}.pb.go $pbDest
mv $grpc/${project_name}.pb.gw.go $pbDest