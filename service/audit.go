package service

import (
	"errors"

	"github.com/urishabh12/audit_logs/model"
	audit_proto "github.com/urishabh12/audit_logs/proto"
)

type AuditService struct {
	auditModel *model.AuditModel
}

func NewAuditService(am *model.AuditModel) *AuditService {
	return &AuditService{
		auditModel: am,
	}
}

func (a *AuditService) Log(req *audit_proto.LogRequest) (*audit_proto.LogResponse, error) {
	err := verifyLogRequest(req)
	if err != nil {
		return nil, err
	}

	err = a.auditModel.InsertAuditLog(req)
	if err != nil {
		return nil, err
	}

	return &audit_proto.LogResponse{
		Status: true,
	}, nil
}

func (a *AuditService) GetLogForTenantByEntity(req *audit_proto.LogByEntityRequest) (*audit_proto.LogsResponse, error) {
	if req.TenantID == 0 || req.Entity == "" || req.Timestamp == 0 {
		return nil, errors.New("every request field is required can't be empty")
	}

	resp, err := a.auditModel.GetTenantLogByEntity(req.TenantID, req.Entity, req.Timestamp)
	if err != nil {
		return nil, err
	}

	out := &audit_proto.LogsResponse{
		Logs: resp,
	}
	return out, nil
}

func verifyLogRequest(req *audit_proto.LogRequest) error {
	if req.TenantID == 0 || req.UserID == 0 || req.EntityID == 0 || req.Entity == "" || req.Action == "" || req.Timestamp == 0 {
		return errors.New("every request field is required can't be empty")
	}
	return nil
}
