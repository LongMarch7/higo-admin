package admin

import (
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/define"
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)



func (a* adminController)LogoutPost(ctx context.Context) (rs string , err error){
    if param := base.GetParamByCtx(ctx); param != nil{
        cookie := param.Cookie
        cookie.T = ""
        header := metadata.Pairs(define.ResCookieName, cookie.Marshal())
        grpc.SetHeader(ctx, header)
        return base.NewHtmlRet(0, "退出成功","/admin/login"), nil
    }else{
        return base.NewHtmlRet(-1, "参数错误",""), nil
    }
}
