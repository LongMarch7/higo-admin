package setting

import "github.com/LongMarch7/higo/controller/base"

type adminSettingController struct {
}
var Controller = &adminSettingController{}
var name = "admin/setting"
func Init(){
    base.AddController(name, Controller)
}
