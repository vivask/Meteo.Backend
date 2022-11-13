package entities

type Homezone struct {
	ID      uint32 `gorm:"column:id;primaryKey" json:"id"`
	Name    string `gorm:"column:name;not null;unique;index;size:100" json:"name"`
	Address string `gorm:"column:address;not null;size:20" json:"address"`
	Mac     string `gorm:"column:mac;size:20" json:"mac"`
	Note    string `gorm:"column:note" json:"note"`
	Active  bool   `json:"active"`
}

func (Homezone) TableName() string {
	return "homezones"
}
