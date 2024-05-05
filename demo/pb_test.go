syntax = "proto3";

package demo

service Greeter {
rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
string name = 1;
}

message HelloReply {
string message = 1;
}