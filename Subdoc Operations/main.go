package main

import (
	"doctype/database"
	"doctype/utils"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

type Airline struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	CallSign string `json:"callsign"`
	Country  string `json:"country"`
}

func main() {

	var (
		endpoint       = "localhost"
		bucketName     = "travel-sample"
		scopeName      = "inventory"
		collectionName = "airline"
		userName       = "Kira"
		password       = "Kira@123"
	)

	/*
	 TODO: Need to connect with the deisred cluster and collections.

	*/
	clusterObj := database.ConnectWithCluster(endpoint, userName, password)

	travelSampleBucket := clusterObj.Bucket(bucketName)
	airlineCollectionObj := travelSampleBucket.Scope(scopeName).Collection(collectionName)

	fmt.Println("airlineCollectionObj: ", airlineCollectionObj)

	// ? Lookup subdoc api

	/*
	* fetches only a part of the document. Faster as data transfered is less.
	 */

	// * Lookup for a subdoc
	resultObjList := make([]interface{}, 2)
	lookupOps := []gocb.LookupInSpec{
		gocb.GetSpec("city", &gocb.GetSpecOptions{}),
		gocb.ExistsSpec("country", &gocb.ExistsSpecOptions{}),
	}
	lookupResult, err := airlineCollectionObj.LookupIn("airline_10", lookupOps, &gocb.LookupInOptions{})
	utils.CheckErrorNil(err)

	// * Getting the content
	lookupResult.ContentAt(0, &resultObjList[0])
	lookupResult.ContentAt(1, &resultObjList[1])

	//  * Printing the result content
	fmt.Println("result list: ", resultObjList)
	fmt.Println("is result exists: ", lookupResult.Exists(0))
	fmt.Println("is result exists: ", lookupResult.Exists(1))

	// ? Mutate subdoc api

	/*
	* Multiple opeartions on the same doc.
	* The mutation is atomic so if one of the opeartion fails it is rolled back.
	 */
	mutateops := []gocb.MutateInSpec{
		gocb.UpsertSpec("city", "LA", &gocb.UpsertSpecOptions{}),
		gocb.RemoveSpec("country", &gocb.RemoveSpecOptions{}),
	}
	mutateResult, err := airlineCollectionObj.MutateIn("airline_10", mutateops, &gocb.MutateInOptions{})
	utils.CheckErrorNil(err)
	fmt.Printf("mutate result: %#v\n", mutateResult.MutationResult.Result)
}
