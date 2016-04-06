package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost"+":8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	genericReply := RPCGenericReply{}
	// err = client.Call("RPCServer.MakeVol", "testvol", &genericReply)
	// if err != nil {
	// 	log.Fatal("MakeVol", err)
	// }
	// log.Println("MakeVol success")

	// listVols := ListVolsReply{}
	// err = client.Call("RPCServer.ListVols", "testvol", &listVols)
	// if err != nil {
	// 	log.Fatal("ListVols", err)
	// }
	// log.Println("ListVols success", listVols)

	// statVol := VolInfo{}
	// err = client.Call("RPCServer.StatVol", "testvol", &statVol)
	// if err != nil {
	// 	log.Fatal("StatVol", err)
	// }
	// log.Println("StatVol success", statVol)

	err = client.Call("RPCServer.DeleteVol", "testvol", &genericReply)
	if err != nil {
		log.Fatal("DeleteVol", err)
	}
	log.Println("DeleteVol success")
}
