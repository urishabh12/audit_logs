package inter_test

import (
	"log"

	audit_proto "github.com/urishabh12/audit_logs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	server_host = "localhost:8000"
)

func GetClient() audit_proto.AuditServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(server_host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	return audit_proto.NewAuditServiceClient(conn)
}
