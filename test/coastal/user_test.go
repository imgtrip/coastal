package main

import (
    "coastal/config/constant"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/pkg/hash"
    "coastal/pkg/magic"
    "coastal/test/coastal/util"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCreateUser(t *testing.T) {
    fakeName := fake.FullName()
    fakeEmail := fake.EmailAddress()
    fakePwd := fake.SimplePassword()

    // TESTING: missing required params
    _, e := testServer.CreateUser(touristCtx(), &pb.CreateUserReq{})
    assert.NotNil(t, e)

    _, e = testServer.CreateUser(touristCtx(), &pb.CreateUserReq{Name: fakeName})
    assert.NotNil(t, e)

    _, e = testServer.CreateUser(touristCtx(), &pb.CreateUserReq{Name: fakeName, Email: fakeEmail})
    assert.NotNil(t, e)

    _, e = testServer.CreateUser(touristCtx(), &pb.CreateUserReq{Name: fakeName, Email: fakeEmail, Password: fakePwd})
    assert.NotNil(t, e)

    _, e = testServer.CreateUser(touristCtx(), &pb.CreateUserReq{Name: fakeName, Email: fakeEmail, Password: fakePwd, VerificationCode: util.VerificationCode()})
    assert.NotNil(t, e)

    // TESTING: duplicate user
    fakeName = fake.FullName()
    fakeEmail = fake.EmailAddress()
    user := model.User{}
    dbConnect.First(&user)
    ctx := touristCtx()
    code := util.ValidVerificationCode(fakeEmail, ctx.GetToken())

    _, e = testServer.CreateUser(ctx, &pb.CreateUserReq{Name: user.Name, Email: fake.EmailAddress(), Password: fakePwd, VerificationCode: code})
    assert.NotNil(t, e)

    ctx = touristCtx()
    code = util.ValidVerificationCode(fakeEmail, ctx.GetToken())
    _, e = testServer.CreateUser(ctx, &pb.CreateUserReq{Name: fake.FullName(), Email: user.Email, Password: fakePwd, VerificationCode: code})
    assert.NotNil(t, e)

    ctx = touristCtx()
    code = util.ValidVerificationCode(fakeEmail, ctx.GetToken())
    _, e = testServer.CreateUser(ctx, &pb.CreateUserReq{Name: user.Name, Email: user.Email, Password: fakePwd, VerificationCode: code})
    assert.NotNil(t, e)

    // create success
    fakeName = fake.FullName()
    fakeEmail = fake.EmailAddress()
    ctx = touristCtx()
    code = util.ValidVerificationCode(fakeEmail, ctx.GetToken())
    req := pb.CreateUserReq{Name: fakeName, Email: fakeEmail, Password: fakePwd, VerificationCode: code}
    res, e := testServer.CreateUser(ctx, &req)
    assert.Nil(t, e)
    assert.Equal(t, req.Name, res.User.Name)
    assert.NotEmpty(t, res.User.AlbumId)
    user, e = model.User{}.ById(res.User.Id)
    assert.Nil(t, e)
    assert.Equal(t, int(user.DailyFreeDownloadNumber), constant.InitDailyFreeDownloadImage)

    token, err := model.Token{}.ByHash(ctx.GetToken())
    if err != nil {
        panic(err)
    }
    assert.NotEmpty(t, token.UserId)

    // referer user success
    fakeName = fake.FullName()
    fakeEmail = fake.EmailAddress()
    user = model.User{}
    dbConnect.First(&user)
    ctx = touristCtx()
    code = util.ValidVerificationCode(fakeEmail, ctx.GetToken())
    e = dbConnect.Model(&model.User{}).Where("referer_user_id=?", user.ID).Update("referer_user_id", 0).Error
    assert.Nil(t, e)
    e = dbConnect.Model(&model.User{}).Where("id=?", user.ID).Update("daily_free_download_number", 0).Error
    assert.Nil(t, e)
    e = dbConnect.Where("id=?", user.ID).First(&user).Error
    assert.Nil(t, e)
    req = pb.CreateUserReq{Name: fakeName, Email: fakeEmail, Password: fakePwd, VerificationCode: code, ReferrerEmail: user.Email}
    res, e = testServer.CreateUser(ctx, &req)
    assert.Nil(t, e)
    afterReferUser := model.User{}
    dbConnect.Where("id=?", user.ID).First(&afterReferUser)
    assert.Equal(t, int(user.DailyFreeDownloadNumber)+constant.RefererUserFreeDownloadIncrementStep, int(afterReferUser.DailyFreeDownloadNumber))

    // referer user maxed
    fakeName = fake.FullName()
    fakeEmail = fake.EmailAddress()
    user = model.User{}
    dbConnect.First(&user)
    e = dbConnect.Model(&model.User{}).Where("id=?", user.ID).Update("daily_free_download_number", constant.MaxRefererNumber*constant.RefererUserFreeDownloadIncrementStep).Error
    assert.Nil(t, e)
    e = dbConnect.Model(&model.User{}).Where("referer_user_id=?", user.ID).Update("referer_user_id", 0).Error
    assert.Nil(t, e)
    e = dbConnect.Model(&model.User{}).Where("referer_user_id=?", 0).Where("id <>?", user.ID).Limit(constant.MaxRefererNumber).Update("referer_user_id", user.ID).Error
    assert.Nil(t, e)
    e = dbConnect.Where("id=?", user.ID).First(&user).Error
    assert.Nil(t, e)

    ctx = touristCtx()
    code = util.ValidVerificationCode(fakeEmail, ctx.GetToken())
    req = pb.CreateUserReq{Name: fakeName, Email: fakeEmail, Password: fakePwd, VerificationCode: code, ReferrerEmail: user.Email}
    _, e = testServer.CreateUser(ctx, &req)
    assert.Nil(t, e)
    afterReferUser = model.User{}
    e = dbConnect.Where("id=?", user.ID).First(&afterReferUser).Error
    assert.Nil(t, e)
    assert.Equal(t, user.DailyFreeDownloadNumber, afterReferUser.DailyFreeDownloadNumber)
}

func TestAuthUser(t *testing.T) {
    var user model.User
    name := fake.LastName()
    email := fake.EmailAddress()
    password := fake.SimplePassword()
    u, e := user.Create(name, email, password)
    assert.Nil(t, e)

    _, e = testServer.AuthUser(touristCtx(), &pb.AuthUserReq{})
    assert.NotNil(t, e)

    _, e = testServer.AuthUser(touristCtx(), &pb.AuthUserReq{Email: u.Email, Password: password + magic.Str.Random(1)})
    assert.NotNil(t, e)

    _, e = testServer.AuthUser(touristCtx(), &pb.AuthUserReq{Email: u.Email + magic.Str.Random(1), Password: password})
    assert.NotNil(t, e)

    _, e = testServer.AuthUser(touristCtx(), &pb.AuthUserReq{Email: magic.Str.Random(10), Password: password})
    assert.NotNil(t, e)

    ctx := touristCtx()
    _, e = testServer.AuthUser(ctx, &pb.AuthUserReq{Email: email, Password: password})
    assert.Nil(t, e)
    token, e := model.Token{}.ByHash(ctx.GetToken())
    assert.Nil(t, e, e)
    assert.Equal(t, u.ID, token.UserId)
}

func TestUpdatePassword(t *testing.T) {
    var user model.User
    name := fake.LastName()
    email := fake.EmailAddress()
    password := fake.SimplePassword()
    _, e := user.Create(name, email, password)
    assert.Nil(t, e)

    _, e = testServer.UpdatePassword(touristCtx(), &pb.UpdatePasswordReq{})
    assert.NotNil(t, e)

    freshPassword := fake.SimplePassword()
    fakeVerificationCode := util.VerificationCode()
    _, e = testServer.UpdatePassword(touristCtx(), &pb.UpdatePasswordReq{Email: email, Password: freshPassword, VerificationCode: fakeVerificationCode})
    assert.NotNil(t, e)

    freshPassword = fake.SimplePassword()
    ctx := touristCtx()
    verificationCode := util.VerificationCode()
    _, e = model.VerificationCode{}.SimpleCreate(email, verificationCode, ctx.GetToken())
    assert.Nil(t, e)
    _, e = testServer.UpdatePassword(ctx, &pb.UpdatePasswordReq{Email: email, Password: freshPassword, VerificationCode: verificationCode})
    assert.Nil(t, e)

    user = model.User{}
    e = dbConnect.Where("email=?", email).First(&user).Error
    assert.Nil(t, e)

    assert.True(t, hash.Compare(freshPassword, user.Password))
}

func TestCreateVerificationCode(t *testing.T) {
    var user model.User
    name := fake.LastName()
    email := fake.EmailAddress()
    password := fake.SimplePassword()
    _, e := user.Create(name, email, password)

    e = dbConnect.Where("email=?", email).Delete(&model.VerificationCode{}).Error
    assert.Nil(t, e)

    _, e = testServer.CreateVerificationEmail(touristCtx(), &pb.CreateVerificationEmailReq{})
    assert.NotNil(t, e)

    // FIXME How to test send email?
    // _, e = testServer.CreateVerificationEmail(touristCtx(), &pb.CreateVerificationEmailReq{Email: email})
    // assert.Nil(t, e)

    // count := 0
    // e = dbConnect.Model(&model.VerificationCode{}).Where("email=?", email).Count(&count).Error
    // assert.Nil(t, e)
    // assert.True(t, count == 1)
    //
    // ctx := touristCtx()
    // _, e = testServer.CreateVerificationEmail(ctx, &pb.CreateVerificationEmailReq{Email: fake.EmailAddress()})
    // assert.Nil(t, e)
    // _, e = testServer.CreateVerificationEmail(ctx, &pb.CreateVerificationEmailReq{Email: fake.EmailAddress()})
    // assert.NotNil(t, e)
}

func TestImageVotes(t *testing.T) {
    _, err := testServer.ImageVotes(touristCtx(), &pb.ImageVotesReq{})
    assert.NotNil(t, err)

    _, err = testServer.ImageVotes(authCtx(), &pb.ImageVotesReq{})
    assert.Nil(t, err)
}

func TestScoreLogs(t *testing.T) {
    // FIXME test score logs
}
