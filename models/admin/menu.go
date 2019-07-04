package admin

import (
    "github.com/LongMarch7/higo-admin/db/object/models"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "strconv"
)


func GetMenuList(user_info []utils.LoginInfo) []models.MicroAdminMenu{
    engine := db.NewDb(db.DefaultNAME).Engine()
    var menuList = make([]models.MicroAdminMenu, 0)
    engine.Table(models.MicroAdminMenu{}.TableName()).Find(&menuList)

    j := 0
    casbin := auth.NewCasbin()
    for _,menu := range menuList{
        for _, user := range user_info{
            if user.RoleId == define.SuperRoleId {
                j = len(menuList)
                break
            }
            if casbin.Enforcer().Enforce(user.RoleName, menu.Url + ":" + menu.Func, menu.Method){
                menuList[j] = menu
                j++
                break
            }
        }
        if j ==  len(menuList){
            break
        }
    }
    return menuList[:j]
}

func GetMenuViewList(user_info []utils.LoginInfo) []models.MicroAdminMenu{
    list :=GetMenuList(user_info)
    j :=0
    for _, value:= range list{
        if value.Type == 0 || value.Type == 1  {
            list[j] = value
            j++
        }
    }
    return list[:j]
}



func GetMenuPower(check []string, no_check []string) (list []utils.MenuPower, success bool){
    engine := db.NewDb(db.DefaultNAME).Engine()
    success = false

    for _,value := range check{
        var menu = models.MicroAdminMenu{}
        id,err :=strconv.Atoi(value)
        if err != nil{
            grpclog.Error("menu Atoi err:" , err.Error())
            return
        }
        _,err =engine.Table(models.MicroAdminMenu{}.TableName()).ID(id).Get(&menu)
        if err != nil{
            grpclog.Error("Get menu power err:" , err.Error())
            return
        }else if menu.Id != id {
            grpclog.Error("Get menu id invalid :",menu)
            return
        }
        list = append(list, utils.MenuPower{Pattern:menu.Url + ":" + menu.Func,Method:menu.Method,Status:true})
    }
    list = append(list, utils.MenuPower{Pattern:"admin:LoginPost", Method: "POST",Status:true})
    list = append(list, utils.MenuPower{Pattern:"admin:Index", Method: "GET",Status:true})
    list = append(list, utils.MenuPower{Pattern:"admin:Welcome", Method: "GET",Status:true})
    list = append(list, utils.MenuPower{Pattern:"admin/user:Update", Method: "GET",Status:true})
    list = append(list, utils.MenuPower{Pattern:"admin:Upload", Method: "POST",Status:true})
    for _,value := range no_check{
        var menu = models.MicroAdminMenu{}
        id,err :=strconv.Atoi(value)
        if err != nil{
            grpclog.Error("no_check  menu Atoi err:" , err.Error())
            return
        }
        _,err =engine.Table(models.MicroAdminMenu{}.TableName()).ID(id).Get(&menu)
        if err != nil {
            grpclog.Error("Get no_check menu power err:",err.Error())
            return
        }else if menu.Id != id {
            grpclog.Error("Get no_check menu id invalid :",menu)
            return
        }
        list = append(list, utils.MenuPower{Pattern:menu.Url + ":" + menu.Func,Method:menu.Method,Status:false})
    }

    success = true
    return
}
