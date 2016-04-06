package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type RPCServer struct {
	disk StorageAPI
}

// Volume operations

func (r *RPCServer) MakeVol(arg *string, reply *RPCGenericReply) error {
	return r.disk.MakeVol(*arg)
}

func (r *RPCServer) ListVols(arg *string, reply *ListVolsReply) error {
	vols, err := r.disk.ListVols()
	if err != nil {
		return err
	}
	reply.Vols = vols
	return nil
}

func (r *RPCServer) StatVol(arg *string, reply *VolInfo) error {
	var err error
	*reply, err = r.disk.StatVol(*arg)
	if err != nil {
		return err
	}
	if volInfo, err := r.disk.StatVol(*arg); err != nil {
		return err
	} else {
		*reply = volInfo
	}
	return nil
}

func (r *RPCServer) DeleteVol(arg *string, reply *RPCGenericReply) error {
	return r.disk.DeleteVol(*arg)
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
	diskPath := os.Args[1]
	if diskPath == "" {
		log.Fatal("diskPath not given")
	}
	server := new(RPCServer)
	disk := new(storageDisk)
	disk.diskPath = diskPath
	server.disk = disk
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
