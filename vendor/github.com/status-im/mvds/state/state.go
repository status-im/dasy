// Package state contains everything related to the synchronization state for MVDS.
package state

type RecordType int

const (
	OFFER RecordType = iota
	REQUEST
	MESSAGE
)

type State struct {
	Type      RecordType
	SendCount uint64
	SendEpoch int64
}

type SyncState interface {
	Get(group GroupID, id MessageID, peer PeerID) (State, error)
	Set(group GroupID, id MessageID, peer PeerID, newState State) error
	Remove(group GroupID, id MessageID, peer PeerID) error
	Map(epoch int64, process func(GroupID, MessageID, PeerID, State) State) error
}