package main

import "time"

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

type ListVolsReply struct {
	Vols []VolInfo
}

type ListDirArgs struct {
	Vol    string
	prefix string
	marker string
	count  int
}

type ListDirReply struct {
	Files []FileInfo
}

type OpenFileArgs struct {
	Vol    string
	Path   string
	Offset int64
}

type OpenFileReply struct {
	Handle string
}

type ReadFileArgs struct {
	Handle string
	Count  int
}

type ReadFileReply struct {
	Data []byte
	EOF  bool
}

type CreateFileArgs struct {
	Vol  string
	Path string
}

type CreateFileReply struct {
	Handle string
}

type AppendFileArgs struct {
	Handle string
	Data   []byte
}

type CloseFileArgs struct {
	Handle string
}

type StatFileArgs struct {
	Vol  string
	path string
}

type DeleteFileArgs struct {
	Vol  string
	path string
}
