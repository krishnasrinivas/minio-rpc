package main

import "time"

type RPCGenericReply struct{}
type RPCGenericRequest struct{}

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
