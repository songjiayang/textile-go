package mobile

import (
	"github.com/golang/protobuf/proto"
	"github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/pb"
)

// RegisterCafe calls core RegisterCafe
func (m *Mobile) RegisterCafe(host string) error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	if _, err := m.node.RegisterCafe(host); err != nil {
		return err
	}
	return nil
}

// CafeSessions calls core CafeSessions
func (m *Mobile) CafeSessions() ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	sessions, err := m.node.CafeSessions()
	if err != nil {
		return nil, err
	}
	cafeSessions := &pb.CafeSessions{
		Values: sessions,
	}
	bytes, err := proto.Marshal(cafeSessions)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// CafeSession calls core CafeSession
func (m *Mobile) CafeSession(peerId string) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	session, err := m.node.CafeSession(peerId)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}
	bytes, err := proto.Marshal(session)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// RefreshCafeSession calls core RefreshCafeSession
func (m *Mobile) RefreshCafeSession(peerId string) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	session, err := m.node.RefreshCafeSession(peerId)
	if err != nil {
		return nil, err
	}
	bytes, err := proto.Marshal(session)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// DeegisterCafe calls core DeregisterCafe
func (m *Mobile) DeregisterCafe(peerId string) error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.DeregisterCafe(peerId)
}

// CheckCafeMessages calls core CheckCafeMessages
func (m *Mobile) CheckCafeMessages() error {
	if !m.node.Started() {
		return core.ErrOffline
	}

	return m.node.CheckCafeMessages()
}
