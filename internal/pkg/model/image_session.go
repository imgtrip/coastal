package model

type ImageSession struct {
	Session string
	ImageId uint64
}


func (ImageSession) ImageIds(session string) ([]uint64, error) {
	var ids []uint64
	var sessions []ImageSession
	err := Connect.Select("image_id").Where(&ImageSession{Session: session}).Find(&sessions).Pluck("image_id", &ids).Error
	return ids, err
}

func (ImageSession) BatchSave(session string, ids []uint64) {
	for _, id := range ids {
		user := ImageSession{Session: session, ImageId: id}
		Connect.Create(&user)
	}
}
