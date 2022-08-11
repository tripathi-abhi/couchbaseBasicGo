package database

import (
	"doctype/utils"
	"time"

	"github.com/couchbase/gocb/v2"
)

func ConnectWithCluster(clusterEndpoint string, username string, password string) *gocb.Cluster {
	clusterObj, connError := gocb.Connect("couchbase://"+clusterEndpoint, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		}})

	utils.CheckErrorNil(connError)
	connError = clusterObj.WaitUntilReady(3*time.Second, nil)
	utils.CheckErrorNil(connError)
	return clusterObj
}
