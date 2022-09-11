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

func (a *AuditServer) GetLogForTenantByEntity(ctx context.Context, in *ap.LogByEntityRequest) (*ap.LogsResponse, error) {
	out, err := a.auditService.GetLogForTenantByEntity(in)
	return out, err
}
