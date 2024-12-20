package Model

import "time"

type User struct {
	UserID       int    `gorm:"type:int;not null"`
	Name         string `gorm:"type:varchar(20);not null"`
	Sex          int    `gorm:"type:INT;not null"`
	departmental string `gorm:"type:varchar(255);not null"`
	Leader       int    `gorm:"type:int;not null"`
}

type Application struct {
	Name       string    `gorm:"type:varchar(20);not null"`
	UserID     int       `gorm:"type:int;not null"`
	Message    string    `gorm:"type:varchar(255);not null"`
	StartTime  time.Time `gorm:"type:datetime;not null"`
	EndTime    time.Time `gorm:"type:datetime;not null"`
	Department int       `gorm:"type:varchar(20);not null"`
	Leave_type int       `gorm:"type:varchar(20);not null"` //0事假 1出差 2外派
	Status     int       `gorm:"type:int;not null"`
}

func (U *User) TableName() string {
	return "user" // 指定表名为 "user"
}

func (A *Application) TableName() string {
	return "Applications" // 指定表名为 "pending_applications"
}
