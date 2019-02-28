package coastal

import (
    "coastal/internal/app/coastal/validator"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "golang.org/x/net/context"
)

func (*Server) CreateZoomLog(c context.Context, req *pb.CreateZoomLogReq) (*pb.CreateZoomLogRes, error) {
    empty := &pb.CreateZoomLogRes{}

    if err := validator.CreateZoomLog(req); err != nil {
        return empty, err
    }

    if err := ZoomLogModel.Create(req.ImageId, req.Fingerprint); err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) CreateDownloadLog(c context.Context, req *pb.CreateDownloadLogReq) (*pb.CreateDownloadLogRes, error) {
    empty := &pb.CreateDownloadLogRes{}

    if err := validator.CreateDownloadLog(req); err != nil {
        return empty, err
    }

    if err := DownloadLogModel.Create(req.ImageId, req.Fingerprint); err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) CreateSearchLog(c context.Context, req *pb.CreateSearchLogReq) (*pb.CreateSearchLogRes, error) {
    empty := &pb.CreateSearchLogRes{}

    return empty, nil
}

func (*Server) CreateErrorLog(c context.Context, req *pb.CreateErrorLogReq) (*pb.CreateErrorLogRes, error) {
    res := &pb.CreateErrorLogRes{}

    err := model.DB.Create(&model.ErrorLog{
        Code:        req.Code,
        Message:     req.Message,
        Url:         req.Url,
        Payload:     req.Payload,
        Environment: req.Environment,
        Header:      req.Header,
        Cookie:      req.Cookie,
    }).Error

    return res, err
}
