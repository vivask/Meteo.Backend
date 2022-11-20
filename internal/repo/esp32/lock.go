package repo

import (
	"fmt"
	"meteo/internal/kit"
	"sync"
)

var mux sync.Mutex

var (
	lockBmx280   = false
	lockDs18b20  = false
	lockZe08ch2o = false
	lockRadsens  = false
	lockMics6814 = false
)

func UnlockBmx280(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/unlock/bmx280", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockBmx280 = false
	mux.Unlock()
	return nil
}

func LockBmx280(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/lock/bmx280", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockBmx280 = true
	mux.Unlock()
	return nil
}

func isLockedBmx280() bool {
	mux.Lock()
	defer mux.Unlock()
	return lockBmx280
}

func UnlockDs18b20(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/unlock/ds18b20", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockDs18b20 = false
	mux.Unlock()
	return nil
}

func LockDs18b20(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/lock/ds18b20", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockDs18b20 = true
	mux.Unlock()
	return nil
}

func isLockedDs18b20() bool {
	mux.Lock()
	defer mux.Unlock()
	return lockDs18b20
}

func UnlockZe08ch2o(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/unlock/ze08ch2o", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockZe08ch2o = false
	mux.Unlock()
	return nil
}

func LockZe08ch2o(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/lock/ze08ch2o", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockZe08ch2o = true
	mux.Unlock()
	return nil
}

func isLockedZe08ch2o() bool {
	mux.Lock()
	defer mux.Unlock()
	return lockZe08ch2o
}

func UnlockRadsens(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/unlock/radsens", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockRadsens = false
	mux.Unlock()
	return nil
}

func LockRadsens(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/lock/radsens", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockRadsens = true
	mux.Unlock()
	return nil
}

func isLockedRadsens() bool {
	mux.Lock()
	defer mux.Unlock()
	return lockRadsens
}

func UnlockMics6814(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/unlock/mics6814", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockMics6814 = false
	mux.Unlock()
	return nil
}

func LockMics6814(ext bool) error {
	if ext {
		_, err := kit.PutExt("/esp32/database/lock/mics6814", nil)
		if err != nil {
			return fmt.Errorf("error PUT: %w", err)
		}
	}
	mux.Lock()
	lockMics6814 = true
	mux.Unlock()
	return nil
}

func isLockedMics6814() bool {
	mux.Lock()
	defer mux.Unlock()
	return lockMics6814
}
