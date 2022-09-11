CREATE KEYSPACE audit WITH REPLICATION = { 
   'class' : 'NetworkTopologyStrategy',
   'datacenter1' : 3 
};

CREATE TABLE audit.audit_logs (
    tenant_id bigint,
    user_id bigint,
    time_stamp bigint,
    entity_id bigint,
    entity text,
    action text,
    data text,
    PRIMARY KEY ((tenant_id), entity, time_stamp)
);