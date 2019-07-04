package user

import (
    "context"
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/controller/base"
)

func (u* adminUserController)Update(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil {
        if ok, _ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok {
            auth.NewCasbin().ReloadPolicy()
            utils.UpdateRoleStatus()
            return base.NewLayuiRet(0, "更新成功",0, nil), nil
        }
    }
    return base.NewLayuiRet(-2, "更新失败",0, nil), nil
}
