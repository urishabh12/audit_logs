package main

import (
	"context"
	"log"
	"time"

	inter_test "github.com/urishabh12/audit_logs/integration"
	audit_proto "github.com/urishabh12/audit_logs/proto"
	"google.golang.org/grpc"
)

type tests struct {
	name string
	run  func(client audit_proto.AuditServiceClient)
}

var testCases []tests = []tests{
	{
		name: "Insert Log",
		run: func(client audit_proto.AuditServiceClient) {
			in := &audit_proto.LogRequest{
				TenantID:  1005,
				UserID:    1,
				Timestamp: time.Now().Unix(),
				EntityID:  220,
				Entity:    "Sample",
				Action:    "Create",
				Data:      "New Sample",
			}
			out, err := client.Log(context.Background(), in)
			if !out.Status || err != nil {
				log.Panic("error while inserting log", err)
			}
			log.Println()
		},
	},
}

//starts cassandra docker
//setups cassandra
//starts the server on another go routine
//runs all the tests
func main() {
	inter_test.StartDocker()

	inter_test.SetupCassandra()

	//start server
	grpcServer := grpc.NewServer()
	go inter_test.StartServer(grpcServer)

	//connect client
	log.Println("[Test] Sleeping to let server start")
	time.Sleep(30 * time.Second)
	auditClient := inter_test.GetClient()

	//run tests
	for _, test := range testCases {
		log.Println("[Test] Running test case", test.name)
		test.run(auditClient)
		log.Println("[Test] Succenfully ran test case", test.name)
	}

	//stop server
	grpcServer.Stop()
}
