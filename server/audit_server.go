package audit_server

import (
	"context"

	ap "github.com/urishabh12/audit_logs/proto"
	audit_service "github.com/urishabh12/audit_logs/service"
)

type AuditServer struct {
	auditService *audit_service.AuditService
}

func NewAuditServer(as *audit_service.AuditService) *AuditServer {
	return &AuditServer{
		auditService: as,
	}
}

func (a *AuditServer) Log(ctx context.Context, in *ap.LogRequest) (*ap.LogResponse, error) {
	out, err := a.auditService.Log(in)
	return out, err
}

func (a *AuditServer) GetLogByEntity(ctx context.Context, in *ap.LogByEntityRequest) (*ap.LogsResponse, error) {
	out, err := a.auditService.GetLogByEntity(in)
	return out, err
}

func (a *AuditServer) GetLogByEntityPaginated(ctx context.Context, in *ap.LogByEntityPagedRequest) (*ap.LogsResponse, error) {
	out, err := a.auditService.GetLogByEntityPaginated(in)
	return out, err
}

func (a *AuditServer) GetLogByEntityID(ctx context.Context, in *ap.LogByEntityIDRequest) (*ap.LogsResponse, error) {
	out, err := a.auditService.GetLogByEntityID(in)
	return out, err
}
