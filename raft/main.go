package main

import (
	"fmt"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"path/filepath"
)

func main() {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID("127.0.0.1:2000")
	logStore, err := raftboltdb.NewBoltStore(filepath.Join("./",
		"raft-log.bolt"))
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join("./",
		"raft-stable.bolt"))

	fmt.Println(logStore, err)
	fmt.Println(stableStore)
}
