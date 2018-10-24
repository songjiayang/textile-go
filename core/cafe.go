package core

import (
	"github.com/textileio/textile-go/repo"
	"gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
	"sync"
)

// RegisterCafe registers this account with another peer (the "cafe"),
// which provides a session token for the service
func (t *Textile) RegisterCafe(peerId string) error {
	if !t.Online() {
		return ErrOffline
	}

	// call up the peer, see if they're offering a cafe
	pid, err := peer.IDB58Decode(peerId)
	if err != nil {
		return err
	}
	return t.cafeService.Register(pid)
}

// DeregisterCafe removes the session associated with the given cafe
func (t *Textile) DeregisterCafe(peerId string) error {
	if err := t.touchDatastore(); err != nil {
		return err
	}
	return t.datastore.CafeSessions().Delete(peerId)
}

// ListRegisteredCafes list registered cafe sessions
func (t *Textile) ListRegisteredCafes() ([]repo.CafeSession, error) {
	if err := t.touchDatastore(); err != nil {
		return nil, err
	}
	return t.datastore.CafeSessions().List(), nil
}

// CheckCafeMessages fetches new messages from registered cafes
func (t *Textile) CheckCafeMessages() error {
	if !t.Online() {
		return ErrOffline
	}

	// get active cafe sessions
	sessions := t.datastore.CafeSessions().List()
	if len(sessions) == 0 {
		return nil
	}

	// check each concurrently
	wg := sync.WaitGroup{}
	var cerr error
	for _, session := range sessions {
		cafe, err := peer.IDB58Decode(session.CafeId)
		if err != nil {
			cerr = err
			continue
		}
		wg.Add(1)
		go func(cafe peer.ID) {
			if err := t.cafeService.CheckMessages(cafe); err != nil {
				cerr = err
			}
			wg.Done()
		}(cafe)
	}
	wg.Wait()
	return cerr
}
