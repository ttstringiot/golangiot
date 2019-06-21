package rebalance

import (
	"github.com/ttstringiot/golangiot/stgnet/netm"
)

type ConsumerIdsChangeListener interface {
	ConsumerIdsChanged(group string, channels []netm.Context)
}
