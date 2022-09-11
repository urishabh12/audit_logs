# Audit Logs

## Getting Started
Clone the repo

### Start Cassandra 
```
docker run --name cassandra-1 -p 9042:9042 -d cassandra:3.7
```

### Start Server
```
go run main.go
```

## Example
View ```example_client``` directory for example client

## Data Model
```Tenant Id``` Tenant on which action was made

```Entity``` The entity(feature) on which action was made example

```Action``` The action is one word describing action on entity 

```Entity Id``` Unique Id for the entity

```User Id``` User that made those actions

```Data``` Any extra metadata, can be the new values that were updated.
