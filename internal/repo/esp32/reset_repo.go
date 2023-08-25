package repo

import (
	"fmt"
)

func (p esp32Service) ResetAccessPoint() error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.SetupMode = true
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetStm32() error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.Reboot = true
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetRadsens() error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.RadsensHVMode = false
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetRadsensHV(val uint8) error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.RadsensHVState = val != 0
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetRadsensSens(val uint8) error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.RadsensSensitivity = int(val)
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetJournal() error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.ClearJournalEsp32 = false
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}

func (p esp32Service) ResetAvr(val bool) error {
	s, err := p.GetSettings()
	if err != nil {
		return fmt.Errorf("error read settings: %w", err)
	}

	s.DigisparkReboot = val
	err = p.db.Save(s).Error
	if err != nil {
		return fmt.Errorf("error save settings: %w", err)
	}
	return nil
}
