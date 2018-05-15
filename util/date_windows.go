// +build windows

package util

import (
	"time"

	"4d63.com/tz"
)

// Eastern returns the eastern timezone.
func (d date) Eastern() *time.Location {
	if _eastern == nil {
		_easternLock.Lock()
		defer _easternLock.Unlock()
		if _eastern == nil {
			var err error
			_eastern, err = tz.LoadLocation("EST")
			if err != nil {
				panic(err)
			}
		}
	}
	return _eastern
}
