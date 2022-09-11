package main

import (
	"log"
	"net"
	"time"

	"github.com/gocql/gocql"
	audit_model "github.com/urishabh12/audit_logs/model"
	audit_proto "github.com/urishabh12/audit_logs/proto"
	audit_server "github.com/urishabh12/audit_logs/server"
	audit_service "github.com/urishabh12/audit_logs/service"
	"google.golang.org/grpc"
)

func main() {
	//Create Cassandra Session
	session, err := GetCassandraSession()
	if err != nil {
		log.Fatal(err)
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
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	audit_proto.RegisterAuditServiceServer(grpcServer, auditServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func GetCassandraSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	return cluster.CreateSession()
}
