package main

import (
    "github.com/LongMarch7/higo/config"
    "github.com/LongMarch7/higo/app"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/service/base"
    "github.com/LongMarch7/higo/router"
    "github.com/LongMarch7/higo/service/utils"
    "github.com/LongMarch7/higo/service/web"
    "github.com/LongMarch7/higo/util/global"
    "google.golang.org/grpc/grpclog"
    "net/http"
    "os"
    "os/signal"
    "sync"
    "time"
)

type cliConfig struct{
    cOpts        []app.COption
    mOpts        []middleware.MOption
    serviceList  []config.ServiceList
    router       *router.Router
    port         string
    domain       string
}

var c chan os.Signal
var wg sync.WaitGroup

func Producer(){
Loop:
    for{
        select {
        case s := <-c:
            grpclog.Info("Producer get", s)
            break Loop
        default:
        }
        time.Sleep(500 * time.Millisecond)
    }
    wg.Done()
}

func GateWay(config *config.Configer) {

    grpclog.SetLoggerV2(LogConfig(config.Config.Logger,config.Name + ".log").NewLogger())

    if len(config.Config.ServiceList) == 0 {
        grpclog.Error("not have service")
        return
    }
    cliConf := cliResolving(config)
    //mw := middleware.NewMiddleware(middleware.Prefix("gateway"),middleware.MethodName("request"))
    mw := middleware.NewMiddleware(cliConf.mOpts...)
    client := app.NewClient(cliConf.cOpts...)

    for _,service := range cliConf.serviceList{
        htmlHandler := func(pattern string)  func(http.ResponseWriter, *http.Request){
            return base.MakeReqDataMiddleware(
                web.MakeHtmlCallHandler(client.GetClientEndpoint(service.Name),pattern))
        }
        apiHandler := func(pattern string)  func(http.ResponseWriter, *http.Request){
            return base.MakeReqDataMiddleware(
                web.MakeApiCallHandler(client.GetClientEndpoint(service.Name),pattern))
        }
        switch service.Name {
        case "AdminServer":
            client.AddEndpoint(app.CMiddleware(mw),app.CServiceName(service.Name))
            cliConf.router.Add([]router.Routs{
                //{"post|get","/admin",base.MakeReqDataMiddleware(
                //    web.MakeHtmlCallHandler(client.GetClientEndpoint(service.Name),"admin:Index"))},
                //{"post|get","/login/{name}",base.MakeReqDataMiddleware(
                //    web.MakeApiCallHandler(client.GetClientEndpoint(service.Name),"admin:Login"))},

                /***********html***********/
                //base
                {"get","/admin",htmlHandler("admin:Index")},
                {"get","/admin/welcome",htmlHandler("admin:Welcome")},
                {"get","/admin/login",htmlHandler("admin:Login")},
                //role
                {"get","/admin/user/role_index",htmlHandler("admin/user:RoleIndex")},
                {"get","/admin/user/role_edit",htmlHandler("admin/user:RoleEdit")},
                //user
                {"get","/admin/user/user_add",htmlHandler("admin/user:UserAdd")},
                {"get","/admin/user/user_index",htmlHandler("admin/user:UserIndex")},
                {"get","/admin/user/user_edit",htmlHandler("admin/user:UserEdit")},
                //member
                {"get","/admin/member/index",htmlHandler("admin/member:Index")},
                {"get","/admin/member/edit",htmlHandler("admin/member:Edit")},
                //mail
                {"get","/admin/setting/mail",htmlHandler("admin/setting:Mail")},
                //link
                {"get","/admin/setting/link",htmlHandler("admin/setting:Link")},
                {"get","/admin/setting/link_edit",htmlHandler("admin/setting:LinkEdit")},
                //site
                {"get","/admin/setting/site",htmlHandler("admin/setting:Site")},
                //upload
                {"get","/admin/setting/upload",htmlHandler("admin/setting:Upload")},
                //tag
                {"get","/admin/portal/tag_index",htmlHandler("admin/portal:TagIndex")},
                //cate
                {"get","/admin/portal/cate_index",htmlHandler("admin/portal:CateIndex")},
                {"get","/admin/portal/cate_edit",htmlHandler("admin/portal:CateEdit")},
                {"get","/admin/portal/cate_restore",htmlHandler("admin/portal:CateRestore")},
                //article
                {"get","/admin/portal/article_index",htmlHandler("admin/portal:ArticleIndex")},
                {"get","/admin/portal/article_edit",htmlHandler("admin/portal:ArticleEdit")},

                /***********api***********/
                //base
                {"post","/admin/login_post",apiHandler("admin:LoginPost")},
                {"post","/admin/logout_post",apiHandler("admin:LogoutPost")},
                {"get","/admin/user/update",apiHandler("admin/user:Update")},
                //role
                {"get","/admin/user/role_list",apiHandler("admin/user:RoleList")},
                {"post","/admin/user/role_delete",apiHandler("admin/user:RoleDelete")},
                {"post","/admin/user/role_status_change",apiHandler("admin/user:RoleStatusChange")},
                {"post","/admin/user/role_edit_post",apiHandler("admin/user:RoleEditPost")},
                //user
                {"get","/admin/user/user_list",apiHandler("admin/user:UserList")},
                {"post","/admin/user/user_delete",apiHandler("admin/user:UserDelete")},
                {"post","/admin/user/user_status_change",apiHandler("admin/user:UserStatusChange")},
                {"post","/admin/user/user_edit_post",apiHandler("admin/user:UserEditPost")},
                //member
                {"get","/admin/member/list",apiHandler("admin/member:List")},
                {"post","/admin/member/delete",apiHandler("admin/member:Delete")},
                {"post","/admin/member/status_change",apiHandler("admin/member:StatusChange")},
                {"post","/admin/member/edit_post",apiHandler("admin/member:EditPost")},
                //mail
                {"post","/admin/setting/mail_conf",apiHandler("admin/setting:MailConf")},
                {"post","/admin/setting/mail_test",apiHandler("admin/setting:MailTest")},
                //link
                {"get","/admin/setting/link_list",apiHandler("admin/setting:LinkList")},
                {"post","/admin/setting/link_delete",apiHandler("admin/setting:LinkDelete")},
                {"post","/admin/setting/link_status_change",apiHandler("admin/setting:LinkStatusChange")},
                {"post","/admin/setting/link_edit_post",apiHandler("admin/setting:LinkEditPost")},
                //site
                {"post","/admin/setting/site_post",apiHandler("admin/setting:SitePost")},
                //upload
                {"post","/admin/setting/upload_post",apiHandler("admin/setting:UploadPost")},
                //tag
                {"get","/admin/portal/tag_list",apiHandler("admin/portal:TagList")},
                {"post","/admin/portal/tag_delete",apiHandler("admin/portal:TagDelete")},
                {"post","/admin/portal/tag_status_change",apiHandler("admin/portal:TagStatusChange")},
                {"post","/admin/portal/tag_edit_post",apiHandler("admin/portal:TagEditPost")},
                //cate
                {"post","/admin/portal/cate_delete",apiHandler("admin/portal:CateDelete")},
                {"post","/admin/portal/cate_status_change",apiHandler("admin/portal:CateStatusChange")},
                {"post","/admin/portal/cate_edit_post",apiHandler("admin/portal:CateEditPost")},
                //article
                {"get","/admin/portal/article_list",apiHandler("admin/portal:ArticleList")},
                {"post","/admin/portal/article_delete",apiHandler("admin/portal:ArticleDelete")},
                {"post","/admin/portal/article_status_change",apiHandler("admin/portal:ArticleStatusChange")},
                {"post","/admin/portal/article_edit_post",apiHandler("admin/portal:ArticleEditPost")},
                {"post","/admin/upload",base.MakeReqDataMiddleware(
                    utils.MakeUploadHandler(client.GetClientEndpoint(service.Name),"admin:Upload"))},
                //user
                {"get","/",htmlHandler("user:Index")},
                {"get","/index",htmlHandler("user:Index")},
            })
        default:
            grpclog.Error("Not found server by name")
            return
        }
    }
    cliConf.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir(global.StaticPath))))
    cliConf.router.PathPrefix("/upload/").Handler(http.StripPrefix("/upload/",http.FileServer(http.Dir(global.UploadPath))))
    c = make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)
    wg.Add(1)
    go http.ListenAndServe(":"+cliConf.port, cliConf.router)
    go Producer()
    wg.Wait()
}

