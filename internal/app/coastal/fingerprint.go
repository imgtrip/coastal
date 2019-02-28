package coastal

import (
	"coastal/internal/app/coastal/response"
	"coastal/internal/app/coastal/validator"
	"coastal/internal/pkg/model"
	"coastal/internal/pkg/pb"
	"golang.org/x/net/context"
)

func (*Server) CreateFingerprint(c context.Context, req *pb.CreateFingerprintReq) (*pb.CreateFingerprintRes, error) {
	empty := &pb.CreateFingerprintRes{}
	auth, err := Auth(c, cachedUser)
	if err != nil {
		return empty, err
	}

	if err := validator.CreateFingerprint(req, auth.ID); err != nil {
		return empty, err
	}

	fp, err := model.Fingerprint{}.Create(req.Hash, auth.ID)
	if err != nil {
		return empty, err
	}

	return response.CreateFingerprint(fp), nil
}

func (*Server) UpdateFingerprint(c context.Context, req *pb.UpdateFingerprintReq) (*pb.UpdateFingerprintRes, error) {
	empty := &pb.UpdateFingerprintRes{}

	auth, err := Auth(c, cachedUser)
	if err != nil {
		return empty, err
	}

	if err = validator.UpdateFingerprint(req, auth.ID); err != nil {
		return empty, err
	}

	_, err = model.Fingerprint{}.Update(req.Hash, req.UserId)
	if err != nil {
		return empty, err
	}

	return empty, nil
}
