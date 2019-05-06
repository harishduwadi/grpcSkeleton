

With command (ran from protoFile directory)

    ```protoc -I ../protoFile --go_out=plugins=grpc:../protoFile ../protoFile/protoFile.proto```

protobuf creates the file protoFile/protoFile.pb.go that enables the server and agent rpc

Generating a temporary password for the Secure TCP connection

    ```openssl genrsa -out privatekey.pem 1024 ```

    ```openssl req -new -x509 -key privatekey.pem -out publickey.cer ```
