package util

import (
    "coastal/config/constant"
    "coastal/internal/env"
    "coastal/internal/pkg/model"
    "google.golang.org/grpc/metadata"
    "net/http"
    "net/http/httptest"
    "time"
)

const (
    Query = "start=0&limit=10"
)

// =====================TOURIST CONTEXT===================================
type TouristContext struct {
    token string
}

type TestServer func(*http.Request) *httptest.ResponseRecorder

func (ctx TouristContext) SetToken(token string) TouristContext {
    ctx.token = token
    return ctx
}

func (ctx TouristContext) GetToken() string {
    return ctx.token
}

func (TouristContext) Deadline() (deadline time.Time, ok bool) {
    return time.Time{}, true
}

func (TouristContext) Done() <-chan struct{} {
    return nil
}

func (TouristContext) Err() error {
    return nil
}

func (ctx TouristContext) Value(key interface{}) interface{} {
    return metadata.MD{constant.TokenName: []string{ctx.token}}
}

// =====================AUTH CONTEXT===================================

type AuthContext struct {
    token string
    user  model.User
    album model.Album
}

func (AuthContext) Deadline() (deadline time.Time, ok bool) {
    return time.Time{}, true
}
func (AuthContext) Done() <-chan struct{} {
    return nil
}

func (AuthContext) Err() error {
    return nil
}

func (ctx AuthContext) SetToken(token string) AuthContext {
    ctx.token = token
    return ctx
}

func (ctx AuthContext) GetToken() string {
    return ctx.token
}

func (ctx AuthContext) SetAlbum(album model.Album) AuthContext {
    ctx.album = album
    return ctx
}

func (ctx AuthContext) GetAlbum() model.Album {
    return ctx.album
}

func (ctx AuthContext) SetUser(user model.User) AuthContext {
    ctx.user = user
    return ctx
}

func (ctx AuthContext) GetUser() model.User {
    return ctx.user
}

func (ctx AuthContext) GetUserId() uint64 {
    return ctx.GetUser().ID
}

func (ctx AuthContext) Value(key interface{}) interface{} {
    return metadata.MD{constant.TokenName: []string{ctx.token}}
}

// ========================CERTIFIED SERVER CONTEXT================================

type CertifiedServerContext struct {
    token string
}

func (CertifiedServerContext) Deadline() (deadline time.Time, ok bool) {
    return time.Time{}, true
}
func (CertifiedServerContext) Done() <-chan struct{} {
    return nil
}

func (CertifiedServerContext) Err() error {
    return nil
}

func (ctx CertifiedServerContext) Value(key interface{}) interface{} {
    return metadata.MD{constant.TokenName: []string{env.Process.NodeServerToken}}
}
