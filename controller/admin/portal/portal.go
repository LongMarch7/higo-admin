package portal

import "github.com/LongMarch7/higo/controller/base"

type adminPortalController struct {
}
var Controller = &adminPortalController{}
var name = "admin/portal"
func Init(){
    base.AddController(name, Controller)
}

