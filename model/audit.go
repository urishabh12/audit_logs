package model

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	audit_proto "github.com/urishabh12/audit_logs/proto"
)

type AuditModel struct {
	cassandraSession *gocql.Session
}

func NewAuditModel(cs *gocql.Session) *AuditModel {
	return &AuditModel{
		cassandraSession: cs,
	}
}

func (a *AuditModel) InsertAuditLog(req *audit_proto.LogRequest) error {
	query := mapLogRequestToInsertQuery(req)
	log.Println("[Cassandra]", query)
	return a.cassandraSession.Query(query).Exec()
}

func (a *AuditModel) GetTenantLogByEntity(tenant_id int64, entity string, timestamp int64) ([]*audit_proto.AuditLog, error) {
	query := tenantLogByEntityQuery(tenant_id, entity, timestamp)
	log.Println("[Cassandra]", query)
	resp := a.cassandraSession.Query(query).Iter().Scanner()
	return scanAuditLogResponse(resp)
}

func scanAuditLogResponse(scanner gocql.Scanner) ([]*audit_proto.AuditLog, error) {
	out := []*audit_proto.AuditLog{}

	for scanner.Next() {
		r := &audit_proto.AuditLog{}
		err := scanner.Scan(&r.TenantID, &r.UserID, &r.Timestamp, &r.EntityID, &r.Entity, &r.Action, &r.Data)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}

	return out, nil
}

func mapLogRequestToInsertQuery(req *audit_proto.LogRequest) string {
	return fmt.Sprintf("INSERT INTO audit.audit_logs (tenant_id, user_id, time_stamp, entity_id, entity, action, data) VALUES (%d, %d, %d, %d, '%s', '%s', '%s');", req.TenantID, req.UserID, req.Timestamp, req.EntityID, req.Entity, req.Action, req.Data)
}

func tenantLogByEntityQuery(tenant_id int64, entity string, timestamp int64) string {
	return fmt.Sprintf("SELECT tenant_id, user_id, time_stamp, entity_id, entity, action, data FROM audit.audit_logs WHERE tenant_id=%d and entity='%s' and time_stamp >= %d;", tenant_id, entity, timestamp)
}
