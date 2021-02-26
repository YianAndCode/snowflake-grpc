package main

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"snowflake/proto"

	"google.golang.org/grpc"
)

// SnowflakeService Snowflake gRPC 服务
type SnowflakeService struct{}

// Next 获取下一个唯一 id
func (s SnowflakeService) Next(ctx context.Context, req *proto.NextRequest) (*proto.NextReply, error) {
	res := new(proto.NextReply)
	res.Id = next(req.ServiceId)
	return res, nil
}

// Parse 解析 snowflake id
func (s SnowflakeService) Parse(ctx context.Context, req *proto.ParseRequest) (*proto.ParseReply, error) {
	res := new(proto.ParseReply)
	_, timestamp, nodeID, serviceID, seq := parse(req.Id)
	res.Timestamp = timestamp
	res.NodeId = nodeID
	res.ServiceId = serviceID
	res.Seq = seq
	return res, nil
}

func main() {
	config := getConfig()
	var Addr string
	Addr = config.Server + ":" + strconv.Itoa(config.Port)
	nodeID = int64(config.NodeId)
	listen, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Printf("Failed to listen: %v, program exited\n", err)
		return
	}
	server := grpc.NewServer()
	var snowflakeService SnowflakeService
	proto.RegisterSnowflakeServer(server, snowflakeService)
	fmt.Println("Listen on: " + Addr)
	server.Serve(listen)
}
