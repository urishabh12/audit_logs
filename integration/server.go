package inter_test

import (
	"log"
	"net"

	audit_model "github.com/urishabh12/audit_logs/model"
	audit_proto "github.com/urishabh12/audit_logs/proto"
	audit_server "github.com/urishabh12/audit_logs/server"
	audit_service "github.com/urishabh12/audit_logs/service"
	"google.golang.org/grpc"
)

const (
	server_port = ":8000"
)

func StartServer(grpcServer *grpc.Server) {
	//Create Cassandra Session
	session, err := GetCassandraSession()
	if err != nil {
		log.Panic(err)
		return
	}
	defer session.Close()

	//Initiate Audit Model
	auditModel := audit_model.NewAuditModel(session)

	//Initiate Audit Service
	auditService := audit_service.NewAuditService(auditModel)

	//Initiate Audit Server
	auditServer := audit_server.NewAuditServer(auditService)

	//Start Server
	lis, err := net.Listen("tcp", server_port)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	audit_proto.RegisterAuditServiceServer(grpcServer, auditServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to serve: %s", err)
	}
}
