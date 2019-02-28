package model

type Fingerprint struct {
	CommonFields
	Hash   string
	UserId uint64
}

func (Fingerprint) Create(hash string, userId uint64) (Fingerprint, error) {
	fp := Fingerprint{Hash: hash, UserId: userId}
	err := Connect.Create(&fp).Error
	return fp, err
}

func (Fingerprint) Update(hash string, userId uint64) (Fingerprint, error) {
	fp := Fingerprint{}
	err := Connect.Model(&fp).Where("hash=?", hash).Updates(Fingerprint{UserId: userId}).Error
	return fp, err
}

func (Fingerprint) ByHash(hash string) (Fingerprint, error) {
	fp := Fingerprint{}
	err := Connect.Model(&fp).Where("hash=?", hash).First(&fp).Error
	return fp, err
}
