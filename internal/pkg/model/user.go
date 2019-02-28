package model

import (
    "coastal/config/constant"
    "coastal/pkg/hash"
    "github.com/jinzhu/gorm"
    "time"
)

type User struct {
    CommonFields
    Name                    string
    Email                   string
    Password                string
    Avatar                  string
    AlbumId                 uint64
    RefererUserId           uint64
    DailyFreeDownloadNumber uint64
    IsOfficial              bool
    Score                   uint64
}

func (u User) Existed(name string, email string) (bool, error) {
    nameExisted, err := u.IsNameExisted(name)
    if err != nil {
        return true, err
    }

    emailExisted, err := u.IsEmailExisted(email)
    if err != nil {
        return true, err
    }

    return nameExisted || emailExisted, nil
}

func (User) IsNameExisted(name string) (bool, error) {
    count := 0
    if e := Connect.Model(&User{}).Where("name=?", name).Count(&count).Error; e != nil {
        return true, e
    }
    return count > 0, nil
}

func (User) IsEmailExisted(email string) (bool, error) {
    count := 0
    if e := Connect.Model(&User{}).Where("email=?", email).Count(&count).Error; e != nil {
        return true, e
    }
    return count > 0, nil
}

func (u User) Create(name string, email string, password string) (User, error) {
    hashedPassword, err := hash.In(password)
    if err != nil {
        return u, err
    }

    user := User{
        Name:                    name,
        Email:                   email,
        Password:                hashedPassword,
        DailyFreeDownloadNumber: constant.InitDailyFreeDownloadImage,
    }
    err = DB.Create(&user).Error
    return user, err
}

func (User) ById(id uint64) (User, error) {
    var user User
    err := Connect.Where("id=?", id).Find(&user).Error
    return user, err
}

func (User) ByEmail(email string) (User, error) {
    var user User
    err := Connect.Where("email=?", email).Find(&user).Error
    return user, err
}

func (User) ByEmailAndPassword(email string, password string) (User, error) {
    var u User
    user, err := u.ByEmail(email)
    if err != nil {
        return u, err
    }
    return user, nil
}

func (u User) UpdateReferer(id uint64, refererEmail string) error {
    referer, err := u.ByEmail(refererEmail)
    if err != nil {
        return err
    }

    err = Connect.Model(&User{}).Where("id=?", id).Update("referer_user_id", referer.ID).Error
    if err != nil {
        return err
    }

    referCount, err := u.CountByReferer(referer.ID)
    if err != nil {
        return err
    }
    if referCount < constant.MaxRefererNumber {
        field := "daily_free_download_number"
        err = Connect.Model(&User{}).Where("id=?", referer.ID).Update(field, gorm.Expr(field+" +?", constant.RefererUserFreeDownloadIncrementStep)).Error
        if err != nil {
            return err
        }
    }

    return nil
}

func (User) CountByReferer(id uint64) (int, error) {
    count := 0
    err := Connect.Model(&User{}).Where("referer_user_id=?", id).Count(&count).Error
    return count, err
}

func (User) OfficialUsers() ([]User, error) {
    var user []User
    err := Connect.Where("is_official=?", true).Find(&user).Error
    if err != nil {
        return []User{}, err
    }

    return user, nil
}

func (u User) IsOfficialUser(id uint64) (bool, error) {
    user, err := u.ById(id)
    if err != nil {
        return false, err
    }

    return user.IsOfficial, nil
}

func (u User) UpdateScore(id uint64, score int64, scoreCategoryId uint64) error {
    expr := ""
    if score > 0 {
        expr = "score + ?"
    } else if score < 0 {
        expr = "score - ?"
        score = -score
    } else {
        return nil
    }

    // FIXME 使用事务
    err := ScoreLog{}.Create(id, score, scoreCategoryId)
    if err != nil {
        return err
    }

    return Connect.Model(&User{}).Where("id=?", id).Update("score", gorm.Expr(expr, score)).Error
}

func (u User) IsScoreGreater(id uint64, score uint64) (bool, error) {
    count := 0
    err := Connect.Model(&User{}).Where("id = ?", id).Where("score >= ?", score).Count(&count).Error

    return count > 0, err
}

func (u User) CanFreeDownload(id uint64) (bool, error) {
    logCount, err := FreeDownloadLog{}.CountBy(id, time.Now().Local())
    if err != nil && err != gorm.ErrRecordNotFound {
        return false, err
    }

    status, err := u.IsFreeDownloadNumberGreater(id, logCount)
    if err != nil {
        return false, err
    }

    return status, nil
}

func (User) IsFreeDownloadNumberGreater(id uint64, number int) (bool, error) {
    count := 0
    err := Connect.Model(&User{}).Where("id=?", id).Where("daily_free_download_number > ?", number).Count(&count).Error
    return count > 0, err
}
