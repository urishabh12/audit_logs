CREATE KEYSPACE audit WITH REPLICATION = { 
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
};