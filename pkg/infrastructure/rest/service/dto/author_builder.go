package dto

import (
	"basesvc_v2/pkg/domain/entity"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorkBuilder struct{}

func (rec *AuthorkBuilder) GetAuthorResponse(in interface{}) (interface{}, error) {
	var out *entity.Author
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, status.Error(codes.DataLoss, "failed build response")
	}
	return out, nil
}
