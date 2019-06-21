package remoting

import (
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// RequestProcessor request processor
type RequestProcessor interface {
	ProcessRequest(ctx netm.Context, request *protocol.RemotingCommand) (*protocol.RemotingCommand, error)
}
