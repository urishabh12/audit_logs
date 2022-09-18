package inter_test

import (
	"log"
	"os/exec"
)

func StartDocker() {
	log.Println("[Docker] Starting docker")
	err := exec.Command("/bin/sh", "scripts/cassandra_docker.sh").Run()
	if err != nil {
		StopDocker()
		log.Panic("[Docker] error while starting", err)
	}
}

func StopDocker() {
	log.Println("[Docker] Stopping docker")
	err := exec.Command("/bin/sh", "scripts/cassandra_docker_stop.sh").Run()
	if err != nil {
		log.Panic("[Docker] error while stopping", err)
	}
	log.Println("[Docker] Stopped docker")
}
