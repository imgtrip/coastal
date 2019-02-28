package main

import (
    "coastal/internal/app/coastal"
    "coastal/internal/env"
    "coastal/internal/pkg/model"
    "coastal/internal/runtime"
    "coastal/test/coastal/util"
    "flag"
    "fmt"
    "github.com/icrowley/fake"
    "github.com/jinzhu/gorm"
    "os"
    "testing"
)

var testServer *coastal.Server
var tServer *coastal.Server
var dbConnect *gorm.DB
var db *gorm.DB

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    shutdown()
    os.Exit(code)
}

func setup() {
    env.New()
    if !env.Process.Debug {
        os.Exit(1)
    }
    parseFlag()
    databaseUp()
    runtime.New()
    testServer = coastal.New()
    tServer = testServer
    dbConnect = runtime.DBConnect()
    db = dbConnect
}

func parseFlag() {
    dbHost := flag.String("db_host", "", "")
    dbUser := flag.String("db_user", "", "")
    dbPassword := flag.String("db_password", "", "")
    dbName := flag.String("db_name", "", "")
    dbPort := flag.String("db_port", "", "")
    flag.Parse()

    if *dbHost != "" {
        env.Process.DBHost = *dbHost
    }

    if *dbUser != "" {
        env.Process.DBUser = *dbUser
    }
    if *dbPassword != "" {
        env.Process.DBPassword = *dbPassword
    }
    if *dbName != "" {
        env.Process.DBName = *dbName
    }

    if *dbPort != "" {
        env.Process.DBPort = *dbPort
    }

    if env.Process.DBName == "" || env.Process.DBPort == "" || env.Process.DBUser == "" || env.Process.DBPassword == "" {
        fmt.Println("===========================================================================================")
        fmt.Println("| WARNING : it seems that there is no environment variable set. did you missing set env ? |")
        fmt.Println("===========================================================================================")
    }

}

func shutdown() {
    // databaseDown()
}

func touristCtx() util.TouristContext {
    token, err := model.Token{}.Create(0)
    if err != nil {
        panic(err)
    }
    return util.TouristContext{}.SetToken(token.Hash)
}

func authCtx() util.AuthContext {
    user, err := model.User{}.Create(fake.FullName(), fake.EmailAddress(), fake.SimplePassword())
    if err != nil {
        panic(err)
    }

    album, err := model.Album{}.SimpleCreate(fake.Title(), user.ID)
    if err != nil {
        panic(err)
    }

    token, err := model.Token{}.Create(user.ID)
    if err != nil {
        panic(err)
    }

    user.AlbumId = album.ID
    err = model.DB.Model(model.User{}).Where("id=?", user.ID).Update("album_id", album.ID).Error
    if err != nil {
        panic(err)
    }

    return util.AuthContext{}.SetToken(token.Hash).SetUser(user).SetAlbum(*album)
}
