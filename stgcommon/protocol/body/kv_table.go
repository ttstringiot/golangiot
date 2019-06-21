package body

import (
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

type KVTable struct {
	Table map[string]string `json:"table"`
	*protocol.RemotingSerializable
}

func NewKVTable() *KVTable {
	kvTable := &KVTable{
		Table: make(map[string]string),
	}
	return kvTable
}
