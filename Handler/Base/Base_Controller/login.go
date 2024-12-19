package BaseController

import (
	"errors"
)

func (controller *Base_Controller) Login(token string) (error, []ExcusedApplicationsDTO) {

	ok, _ := controller.IsTokenValid(token)
	if !ok {
		return errors.New("token is failed"), nil
	}

}
