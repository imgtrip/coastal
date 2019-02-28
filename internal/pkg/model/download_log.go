package model

import "time"

type DownloadLog struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	Fingerprint string
	ImageId     uint64
}

func (*DownloadLog) Create(ImageId uint64, fp string) error {
	return Connect.Create(&DownloadLog{ImageId: ImageId, Fingerprint: fp}).Error
}
