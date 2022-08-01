package database

import (
	"kvOperations/utils"

	"github.com/couchbase/gocb/v2"
)

func ConnectWithCluster(clusterEndpoint string, username string, password string) *gocb.Cluster {
	cluster, connError := gocb.Connect("couchbases://"+clusterEndpoint, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		}})

	utils.CheckErrorNil(connError)
	return cluster
}
