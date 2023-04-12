package model

type Node struct {
	ID                      int `gorm:"primaryKey; autoIncrement"`
	CreatedAt               int
	UpdatedAt               int
	UserGroupID             int
	DNSTemplateID           int
	InboundTemplateID       int
	OutboundTemplateID      int
	RoutingID               int
	StreamSettingTemplateID int
	Name                    string  `gorm:"unique; not null"`
	Multiplier              float64 `gorm:"not null"`
	Address                 string  `gorm:"index:url" `
	Port                    uint16  `gorm:"index:url" `
	OverrideAddress         string  `gorm:"index:override_url"`
	OverridePort            uint16  `gorm:"index:override_url"`
}
