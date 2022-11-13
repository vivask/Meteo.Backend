package entities

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateRadAccountView(db *gorm.DB) error {

	query := `CREATE OR REPLACE VIEW accounting AS
	SELECT radacct.radacctid, radacct.username, radacct.nasipaddress,
	radacct.nasportid, radacct.acctstarttime, radacct.acctupdatetime,
	radacct.acctstoptime, radacct.calledstationid, radacct.callingstationid,
	radverified.username AS valid, radverified.callingstationid AS verified
	FROM radacct
	LEFT JOIN radverified on radverified.callingstationid = radacct.callingstationid;`
	err := db.Exec(query).Error
	if err != nil {
		return fmt.Errorf("create view error: %w", err)
	}
	return nil
}
