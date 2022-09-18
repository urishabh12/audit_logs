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
		},
	},
	{
		name: "Get By Entity",
		run: func(client audit_proto.AuditServiceClient) {
			in := &audit_proto.LogByEntityRequest{
				TenantID:  1001,
				Timestamp: 1662974410 - 10000,
				Entity:    "Audio",
			}
			expected_output := &audit_proto.LogsResponse{
				Logs: []*audit_proto.AuditLog{
					{
						TenantID:  1001,
						UserID:    301,
						EntityID:  101,
						Timestamp: 1663324099,
						Entity:    "Audio",
						Action:    "Edit",
						Data:      "{\"title\":\"First\",\"tags\":\"political,new,latest\"}",
					},
					{
						TenantID:  1001,
						UserID:    301,
						EntityID:  101,
						Timestamp: 1663154410,
						Entity:    "Audio",
						Action:    "Edit",
						Data:      "{\"title\":\"First\",\"tags\":\"political\"}",
					},
					{
						TenantID:  1001,
						UserID:    301,
						EntityID:  101,
						Timestamp: 1662974410,
						Entity:    "Audio",
						Action:    "Upload",
						Data:      "{\"title\":\"First\",\"tags\":\"political,new,latest\"}",
					},
				},
			}
			out, err := client.GetLogByEntity(context.Background(), in)
			if err != nil {
				log.Panic("error while getting log by entity", err)
			}

			if len(expected_output.Logs) != len(out.Logs) {
				log.Panicf("length of get log by entity not same %d!=%d\n", len(expected_output.Logs), len(out.Logs))
			}

			for i := 0; i < len(expected_output.Logs); i++ {
				if expected_output.Logs[i].Entity != out.Logs[i].Entity ||
					expected_output.Logs[i].Action != out.Logs[i].Action ||
					expected_output.Logs[i].UserID != out.Logs[i].UserID ||
					expected_output.Logs[i].EntityID != out.Logs[i].EntityID ||
					expected_output.Logs[i].TenantID != out.Logs[i].TenantID ||
					expected_output.Logs[i].Timestamp != out.Logs[i].Timestamp {

					log.Print(expected_output.Logs[i])
					log.Print(out.Logs[i])
					log.Panic("values of get log by entity not same")
				}
			}
		},
	},
}

//starts cassandra docker
//setups cassandra
//starts the server on another go routine
//runs all the tests
func main() {
	inter_test.StartDocker()

	log.Println("[Test] Sleeping to let docker start")
	time.Sleep(30 * time.Second)
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
