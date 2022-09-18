CREATE TABLE audit.audit_logs (
    tenant_id bigint,
    user_id bigint,
    time_stamp bigint,
    entity_id bigint,
    entity text,
    action text,
    data text,
    PRIMARY KEY ((tenant_id), entity, time_stamp)
)
WITH CLUSTERING ORDER BY (entity ASC, time_stamp DESC);