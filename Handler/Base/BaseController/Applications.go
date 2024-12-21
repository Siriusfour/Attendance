package BaseController

import (
	"GrpcTest/MyConfig"
	"GrpcTest/MyProto"
	"context"
)

func (controller *Base_Controller) Applications(ctx context.Context, applications_Info *MyProto.Application) (*MyProto.Succee, error) {

	result := MyConfig.DB.Create(applications_Info)
	if result.Error != nil {
		return &MyProto.Succee{
			Flag:    false,
			ErrInfo: result.Error.Error(),
		}, nil
	}

	return &MyProto.Succee{
		Flag:    true,
		ErrInfo: "",
	}, nil
}
