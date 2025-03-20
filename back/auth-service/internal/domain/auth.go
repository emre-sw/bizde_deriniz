package domain

type Auth struct {
	ID       string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Email    string `json:"email" gorm:"uniqueIndex;type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

func (a *Auth) TableName() string {
	return "auths"
}
