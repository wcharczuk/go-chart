// +build windows

package util

import "time"

// Eastern returns the eastern timezone.
func (d date) Eastern() *time.Location {
	if _eastern == nil {
		_easternLock.Lock()
		defer _easternLock.Unlock()
		if _eastern == nil {
			_eastern, _ = time.LoadLocation("EST")
		}
	}
	return _eastern
}
