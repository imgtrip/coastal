package validator

import (
	"coastal/config/constant"
	"coastal/internal/pkg/errors"
	"coastal/internal/pkg/pb"
)

func CreateZoomLog(req *pb.CreateZoomLogReq) error {
	if len(req.Fingerprint) != constant.FingerprintLen || req.ImageId == 0 {
		return errors.InvalidArgument()
	}
	return nil
}

func CreateDownloadLog(req *pb.CreateDownloadLogReq) error {
	if len(req.Fingerprint) != constant.FingerprintLen || req.ImageId == 0 {
		return errors.InvalidArgument()
	}
	return nil
}
