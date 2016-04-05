package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type RPCServer struct{}

// Volume operations

func (r *RPCServer) MakeVol(arg *string, reply *RPCGenericReply) error {
	log.Println("MakeVol", *arg)
	return nil
}

func (r *RPCServer) ListVols(arg *string, reply *ListVolsReply) error {
	log.Println("StatVol", *arg)
	vols := []VolInfo{{"aaaa", time.Now()}, {"bbbb", time.Now()}, {"cccc", time.Now()}}
	reply.Vols = vols
	return nil
}

func (r *RPCServer) StatVol(arg *string, reply *VolInfo) error {
	log.Println("StatVol", *arg)
	reply.Name = "aaaaa"
	reply.Created = time.Now()
	return nil
}

func (r *RPCServer) DeleteVol(arg *string, reply *RPCGenericReply) error {
	log.Println("DeleteVol", *arg)
	return nil
}

// File operations

func (r *RPCServer) ListDir(arg *ListDirArgs, reply *ListDirReply) error {
	log.Println("ListDir")
	return nil
}

func (r *RPCServer) ListDirRecursive(arg *ListDirArgs, reply *ListDirReply) error {
	log.Println("ListDirRecursive")
	return nil
}

func (r *RPCServer) OpenFile(arg *OpenFileArgs, reply *OpenFileReply) error {
	log.Println("OpenFile")
	return nil
}

func (r *RPCServer) ReadFile(arg *ReadFileArgs, reply *ReadFileReply) error {
	log.Println("ReadFile")
	return nil
}

func (r *RPCServer) CloseFile(arg *CloseFileArgs, reply *RPCGenericReply) error {
	log.Println("CloseFile")
	return nil
}

func (r *RPCServer) CreateFile(arg *CreateFileArgs, reply *CreateFileReply) error {
	log.Println("CreateFile")
	return nil
}

func (r *RPCServer) AppendFile(arg *AppendFileArgs, reply *RPCGenericReply) error {
	log.Println("AppendFile")
	return nil
}

func (r *RPCServer) StatFile(arg *StatFileArgs, reply *FileInfo) error {
	log.Println("StatFile")
	return nil
}

func (r *RPCServer) DeleteFile(arg *DeleteFileArgs, reply *RPCGenericReply) error {
	log.Println("DeleteFile")
	return nil
}

func main() {
	server := new(RPCServer)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)

	}
	done := make(chan bool)
	go func() {
		http.Serve(l, nil)
		done <- true
	}()
	<-done
}
