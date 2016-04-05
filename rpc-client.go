package main

import (
	"log"
	"net/rpc"
	"time"
)

type RPCGenericReply struct{}

// VolInfo - volume info
type VolInfo struct {
	Name    string
	Created time.Time
}

// FileInfo - file info
type FileInfo struct {
	Vol     string
	Name    string
	ModTime time.Time
	Size    int64
	IsDir   bool
	IsEOF   bool
}

type ListVols struct {
	Vols []VolInfo
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost"+":8080")
	if err != nil {
		log.Fatal("dialing:", err)

	}

	genericReply := RPCGenericReply{}
	err = client.Call("RPCServer.MakeVol", "testvol", &genericReply)
	if err != nil {
		log.Fatal("MakeVol", err)
	}
	log.Println("MakeVol success")

	listVols := ListVols{}
	err = client.Call("RPCServer.ListVols", "testvol", &listVols)
	if err != nil {
		log.Fatal("ListVols", err)
	}
	log.Println("ListVols success", listVols)

	statVol := VolInfo{}
	err = client.Call("RPCServer.StatVol", "testvol", &statVol)
	if err != nil {
		log.Fatal("StatVol", err)
	}
	log.Println("StatVol success", statVol)

	err = client.Call("RPCServer.DeleteVol", "testvol", &genericReply)
	if err != nil {
		log.Fatal("MakeVol", err)
	}
	log.Println("DeleteVol success")

	time.Sleep(10 * time.Second)
}
