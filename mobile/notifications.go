package mobile

import (
	"github.com/textileio/textile-go/core"
)

// Notifications call core Notifications
func (m *Mobile) Notifications(offset string, limit int) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	notes := m.node.Notifications(offset, limit)
	if len(notes) == 0 {
		notes = make([]core.NotificationInfo, 0)
	}
	return toJSON(notes)
}

// CountUnreadNotifications calls core CountUnreadNotifications
func (m *Mobile) CountUnreadNotifications() int {
	if !m.node.Started() {
		return 0
	}

	return m.node.CountUnreadNotifications()
}

// ReadNotification calls core ReadNotification
func (m *Mobile) ReadNotification(id string) error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.ReadNotification(id)
}

// ReadAllNotifications calls core ReadAllNotifications
func (m *Mobile) ReadAllNotifications() error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.ReadAllNotifications()
}

// AcceptThreadInviteViaNotification call core AcceptThreadInviteViaNotification
func (m *Mobile) AcceptThreadInviteViaNotification(id string) (string, error) {
	if !m.node.Online() {
		return "", core.ErrOffline
	}

	addr, err := m.node.AcceptThreadInviteViaNotification(id)
	if err != nil {
		return "", err
	}
	return addr.B58String(), nil
}

// IgnoreThreadInviteViaNotification call core IgnoreThreadInviteViaNotification
func (m *Mobile) IgnoreThreadInviteViaNotification(id string) error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.IgnoreThreadInviteViaNotification(id)
}
