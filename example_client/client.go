package main

import (
	"context"
	"log"
	"time"

	audit_proto "github.com/urishabh12/audit_logs/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	auditClient := audit_proto.NewAuditServiceClient(conn)
	addAuditLog(auditClient)
	log.Println("<<<<<<<<<<<<<<<<Get Log by Entity Paged>>>>>>>>>>>>>>>>>")
	getByEntityPaged(auditClient)
	log.Println("<<<<<<<<<<<<<<<<Get Log by Entity>>>>>>>>>>>>>>>>>")
	getByEntity(auditClient)
	log.Println("<<<<<<<<<<<<<<<<Get Log by Entity Id>>>>>>>>>>>>>>>>>")
	getByEntityId(auditClient)
}

func addAuditLog(c audit_proto.AuditServiceClient) {
	req := &audit_proto.LogRequest{
		TenantID:  211,
		UserID:    1,
		Timestamp: time.Now().Unix(),
		EntityID:  11,
		Entity:    "Sample",
		Action:    "Create",
		Data:      "New Sample",
	}

	res, err := c.Log(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}

func getByEntity(c audit_proto.AuditServiceClient) {
	req := &audit_proto.LogByEntityRequest{
		TenantID:  211,
		Entity:    "Sample",
		Timestamp: time.Now().Unix() - 10000,
	}
	res, err := c.GetLogByEntity(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.Logs)
}

func getByEntityPaged(c audit_proto.AuditServiceClient) {
	req := &audit_proto.LogByEntityPagedRequest{
		TenantID:       211,
		Entity:         "Sample",
		StartTimestamp: time.Now().Unix(),
		EndTimestamp:   time.Now().Unix() - 10000,
		PageSize:       2,
	}
	res, err := c.GetLogByEntityPaginated(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.Logs)
}

func getByEntityId(c audit_proto.AuditServiceClient) {
	req := &audit_proto.LogByEntityIDRequest{
		TenantID:  211,
		Entity:    "Sample",
		Timestamp: time.Now().Unix() - 10000,
		EntityID:  11,
	}
	res, err := c.GetLogByEntityID(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.Logs)
}
