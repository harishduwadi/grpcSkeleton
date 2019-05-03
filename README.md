
Skeleton Code was created using protoFile/protoFile.proto as basis.

With command (ran from protoFile directory)

```protoc -I ../protoFile --go_out=plugins=grpc:../protoFile ../protoFile/protoFile.proto```

protobuf creates the methods required to start grpc server in the file protoFile/protoFile.pb.go

Generating a temporary password for the Secure TCP connection

```openssl genrsa -out privatekey.pem 1024 ```
```openssl req -new -x509 -key privatekey.pem -out publickey.cer ```

Issue:
    Currenlty works with insecure tcp (without ssl) and secure (with ssl) but only localhost (agent trying to connect to localhost:PORT). With any/all IP (127.0.0.1) the agent just hangs. 