package BaseController

import (
	"GrpcTest/InitConfig"
	"GrpcTest/Model"
	"GrpcTest/MyProto"
	"context"
)

type Applications_data struct {
	Applications_list    []Applications
	My_Applications_list []Applications
}

func (controller *Base_Controller) Login(ctx context.Context, UserID *MyProto.UserID) (*MyProto.ApplicationArray, error) {

	/*
		ok, UserInfo := controller.IsTokenValid(token)
		if !ok {
			return errors.New("token is failed"), nil
		}
	*/
	var ApplicationArray MyProto.ApplicationArray
	var UserInfo Model.User

	err := InitConfig.DB.Where("user_id=?", UserID.UserID).First(&UserInfo).Error
	if err != nil {
		return &ApplicationArray, err
	}

	if UserInfo.Leader != 0 {
		err = InitConfig.DB.Where("Department=?", UserInfo.Leader).First(&ApplicationArray.ApplicationsList).Error
		if err != nil {
			return &ApplicationArray, err
		}
	}

	err = InitConfig.DB.Where("user_id=?", UserID).First(&ApplicationArray.MyApplications).Error
	if err != nil {
		return &ApplicationArray, err
	}

	return &ApplicationArray, nil

}
