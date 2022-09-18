INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1001, 'Audio', 101, 'Upload', 301, 1662974410, '{"title":"First","tags":"political,new,latest"}');
INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1001, 'Audio', 101, 'Edit', 301, 1663154410, '{"title":"First","tags":"political"}');
INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1001, 'Audio', 101, 'Edit', 301, 1663324099, '{"title":"First Audio","tags":"political"}');

INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1002, 'Channel', 101, 'Create', 310, 1662974410, '{"title":"Car"}');
INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1002, 'Channel', 101, 'Edit', 314, 1663154410, '{"title":"Car Attack"}');
INSERT INTO audit.audit_logs (tenant_id, entity, entity_id, action, user_id, time_stamp, data) VALUES (1002, 'Channel', 102, 'Create', 310, 1663324099, '{"title":"Bike Attack"}');