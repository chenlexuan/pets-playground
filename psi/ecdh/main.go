package main

import (
	"fmt"
	psi_client "github.com/openmined/psi/client"
	"github.com/openmined/psi/datastructure"
	"github.com/openmined/psi/pb"
	psi_server "github.com/openmined/psi/server"
	"google.golang.org/protobuf/proto"
)

const fpr = 0.01
const numClientInputs = 10
const numServerInputs = 100

func main() {
	// Whether to reveal the intersection or just the size
	revealIntersection := true
	client, err := psi_client.CreateWithNewKey(revealIntersection)
	if err != nil {
		fmt.Printf("Failed to create PSI client %v\n", err)
	}
	server, err := psi_server.CreateWithNewKey(revealIntersection)
	if err != nil {
		fmt.Printf("Failed to create PSI server %v\n", err)
	}
	clientInputs := []string{}
	for i := 0; i < numClientInputs; i++ {
		clientInputs = append(clientInputs, "Element "+string(i))
	}
	serverInputs := []string{}
	for i := 0; i < numServerInputs; i++ {
		serverInputs = append(serverInputs, "Element "+string(i*2))
	}

	// Step 1: Create the server setup message
	serverSetup, err := server.CreateSetupMessage(fpr, int64(len(clientInputs)), serverInputs, datastructure.Raw) // use datastructure.BloomFilter or datastructure.Gcs to decrease the communication cost but will have false positives
	if err != nil {
		fmt.Printf("Failed to create server setup message %v\n", err)
	}
	serializedServerSetup, err := proto.Marshal(serverSetup)
	if err != nil {
		fmt.Printf("Failed to serialize the server setup message %v\n", err)
	}
	serverSetup = &pb.ServerSetup{}
	err = proto.Unmarshal(serializedServerSetup, serverSetup)
	if err != nil {
		fmt.Printf("Failed to deserialize the server setup message %v\n", err)
	}

	// Step 2: Create the client request
	request, err := client.CreateRequest(clientInputs)
	if err != nil {
		fmt.Printf("Failed to create client request %v\n", err)
	}
	serializedRequest, err := proto.Marshal(request)
	if err != nil {
		fmt.Printf("Failed to serialize the client request %v\n", err)
	}
	request = &pb.Request{}
	err = proto.Unmarshal(serializedRequest, request)
	if err != nil {
		fmt.Printf("Failed to deserialize the client request %v\n", err)
	}

	// Step 3: Get the server response
	response, err := server.ProcessRequest(request)
	if err != nil {
		fmt.Printf("Failed to process the request: %v\n", err)
	}
	serializedResponse, err := proto.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to serialize the server response %v\n", err)
	}
	response = &pb.Response{}
	err = proto.Unmarshal(serializedResponse, response)
	if err != nil {
		fmt.Printf("Failed to deserialize the server response %v\n", err)
	}

	// Step 4: Compute the intersection
	if revealIntersection {
		intersection, err := client.GetIntersection(serverSetup, response)
		if err != nil {
			fmt.Printf("Failed to get the intersection %v\n", err)
		}
		fmt.Printf("Intersection: %v\n", intersection)
	} else {
		intersectionSize, err := client.GetIntersectionSize(serverSetup, response)
		if err != nil {
			fmt.Printf("Failed to get the intersection size %v\n", err)
		}
		fmt.Printf("Intersection size: %v\n", intersectionSize)
	}

	// cleanup
	server.Destroy()
	client.Destroy()
}
