package BaseController

import (
	"GrpcTest/Handler/Utills"
	"GrpcTest/MyProto"
	"time"
)

type Applications struct {
	Name       string    `json:"Name" binding:"required" `
	UserId     int       `json:"UserId" binding:"required" `
	Message    string    `json:"Message"`
	StartTime  time.Time `json:"StartTime" binding:"required"`
	EndTime    time.Time `json:"EndTime" binding:"required"`
	Department int       `json:"Department" binding:"required"`
	Leave_type int       `json:"Leave_type" binding:"required"`
}

// Base_Controller Base 控制器
type Base_Controller struct {
	MyProto.UnsafeViewServer
}

// IsTokenValid  判断token是否有效
func (Base_Controller *Base_Controller) IsTokenValid(tokenString string) (bool, *Utills.JwtCustomClaims) {
	info, err := Utills.ParseJwt(tokenString)
	if err != nil {
		return false, nil
	}

	return true, info
}

func (Base_Controller *Base_Controller) Hello() {}
