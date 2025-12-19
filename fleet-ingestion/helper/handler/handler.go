package handler

import (
	"fmt"
)

func PanicError(serviceName string, funcName string) (err error) {

	if r := recover(); r != nil {
		err = fmt.Errorf(fmt.Sprintf("error function %s on service %s : ", funcName, serviceName, r))
	}
	return
}
