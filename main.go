package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const MASKED_STR = "********"

type user struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

// zapを使った構造体のマスク方法
func (u user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("id", u.ID)
	enc.AddString("name", MASKED_STR)
	enc.AddString("email", MASKED_STR)
	enc.AddInt64("created_at", u.CreatedAt.UnixNano())
	return nil
}

func main() {
	logger, _ := zap.NewDevelopment()
	user := &user{
		ID:        "9d3d893d9d3daofa",
		Name:      "Zap",
		Email:     "zap@sample.com",
		CreatedAt: time.Now(),
	}
	logger.Info("object log sample", zap.Object("user", user))
}
