package model

type (
	Setting struct {
		ID                int `gorm:"primaryKey; autoIncrement"`
		CreatedAt         int
		UpdatedAt         int
		Brand             string `gorm:"not null"`
		Slogan            string
		APIKey            string
		SubscribePath     string `gorm:"not null"`
		InvitePath        string `gorm:"not null"`
		TimeZone          string `gorm:"not null"`
		PauseRegistration bool
		*EmailSetting
		*TelegramSetting
	}
	EmailSetting struct {
		Host     string
		Port     uint16
		StartTLS bool
		Username string
		Password string
	}
	TelegramSetting struct {
		SendNodeStatus         bool
		SendTicketNotification bool
		SendOrderNotification  bool
	}
)
