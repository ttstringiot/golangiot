package remoting

import (
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// RemotingServer remoting server define
type RemotingServer interface {
	InvokeSync(ctx netm.Context, request *protocol.RemotingCommand, timeoutMillis int64) (*protocol.RemotingCommand, error)
	InvokeAsync(ctx netm.Context, request *protocol.RemotingCommand, timeoutMillis int64, invokeCallback InvokeCallback) error
	InvokeOneway(ctx netm.Context, request *protocol.RemotingCommand, timeoutMillis int64) error
	RegisterProcessor(requestCode int32, processor RequestProcessor)
	RegisterDefaultProcessor(processor RequestProcessor)
	RegisterRPCHook(rpcHook RPCHook)
	RegisterContextListener(contextListener netm.ContextListener)
	Start()
	Shutdown()
}
