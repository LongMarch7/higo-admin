package admin

import (
    "github.com/LongMarch7/higo-admin/models/admin"
    "github.com/LongMarch7/higo/controller/base"
    "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (a* adminController)Upload(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil{
        if ok,_ := admin.IsLogin(param.Cookie.T, param.Cookie.U, param.Pattern, param.Method); ok{
            return "success",nil
        }else{
            return "", status.New(codes.PermissionDenied,"not permission").Err()
        }
    }
    return "", status.New(codes.InvalidArgument,"not permission").Err()
}