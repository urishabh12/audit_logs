package inter_test

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

const (
	cassandra_host     = "127.0.0.1:9042"
	keyspace_query_loc = "migration/create_keyspace.cql"
	table_query_loc    = "migration/create_table.cql"
	seed_data_loc      = "migration/seed_data.cql"
)

func SetupCassandra() {
	log.Println("[Cassandra] Setting up cassandra")
	var session *gocql.Session
	var err error
	//retry 50 times
	for i := 0; i < 50; i++ {
		time.Sleep(2 * time.Second)
		session, err = GetCassandraSession()
	}
	defer session.Close()

	if err != nil {
		log.Panic("[Cassandra] error while connecting ", err)
	}

	createKeyspace(session)
	createTable(session)
	addSeedData(session)
}

func createKeyspace(session *gocql.Session) {
	log.Println("[Cassandra] Creating Keyspace")
	data, err := os.ReadFile(keyspace_query_loc)
	if err != nil {
		log.Panic("[Cassandra] error while reading query", err)
	}
	err = session.Query(string(data)).Exec()
	if err != nil {
		log.Panic("[Cassandra] error while creating keyspace ", err)
	}
}

func createTable(session *gocql.Session) {
	log.Println("[Cassandra] Creating Table")
	data, err := os.ReadFile(table_query_loc)
	if err != nil {
		log.Panic("[Cassandra] error while reading query", err)
	}
	err = session.Query(string(data)).Exec()
	if err != nil {
		log.Panic("[Cassandra] error while creating table ", err)
	}
}

func addSeedData(session *gocql.Session) {
	log.Println("[Cassandra] Adding Seed Data")
	data, err := os.ReadFile(seed_data_loc)
	if err != nil {
		log.Panic("[Cassandra] error while reading query", err)
	}
	queries := strings.Split(string(data), "\n")
	for _, query := range queries {
		err = session.Query(query).Exec()
		log.Println("[Query] ", query)
		if err != nil {
			log.Println("[Cassandra] error preparing seed data ", err)
		}
	}
}

func GetCassandraSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(cassandra_host)
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Timeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	return cluster.CreateSession()
}
