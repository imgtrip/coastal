package model

import "time"

type ZoomLog struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	ImageId     uint64
	Fingerprint string
}

func (*ZoomLog) Create(ImageId uint64, fp string) error {
	return Connect.Create(&ZoomLog{ImageId: ImageId, Fingerprint: fp}).Error
}
