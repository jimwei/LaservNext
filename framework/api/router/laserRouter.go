package Framework

import (
	. "framework/api/handlers"
)

type LaserRouter struct {
	AbstractRouter
}

func (this *LaserRouter) Execute() (bool, error) {
	return true, nil
}
func (this *LaserRouter) GetHandler() *ILaserHandler {
	return nil
}
