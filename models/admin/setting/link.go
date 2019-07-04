package setting

import (
    "encoding/base64"
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "time"
    "github.com/LongMarch7/higo-admin/db/object/models"
)

func GetSysLinkListWhereSql(where map[string]string) (string, []interface{}) {
    var sql = "1 = 1"
    var value []interface{}
    if v, ok := where["link_name"]; ok && v != "" {
        keywords := define.TrimString(v)
        sql += " AND name like ?"
        value = append(value,"%"+keywords+"%")
    }

    if v, ok := where["create_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND create_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["create_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND create_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["update_start_time"]; ok && v != "" {
        startTime := define.GetTimestamp(v)
        sql += " AND update_time >= ?"
        value = append(value,strconv.Itoa(int(startTime)))
    }

    if v, ok := where["update_end_time"]; ok && v != "" {
        endTime := define.GetTimestamp(v)
        sql += " AND update_time <= ?"
        value = append(value,strconv.Itoa(int(endTime)))
    }

    if v, ok := where["status_flag"]; ok && v != "" {
        status, _ :=strconv.Atoi(v)
        if status >0 && status <3{
            sql += " AND status = ?"
            value = append(value,status)
        }
    }
    return sql, value
}

func GetLinkList(where map[string]string, page int, rows int)  (link_list []models.MicroLink, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    count = 0
    sql , value:= GetSysLinkListWhereSql(where)
    link_list = make([]models.MicroLink, 0)
    err := engine.Table(models.MicroLink{}.TableName()).Where(sql, value...).Limit(rows,(page-1)*rows).
        Find(&link_list)
    if err != nil {
        grpclog.Error("get link list err: ",err.Error())
        return
    }
    ret,err :=engine.Table(models.MicroLink{}.TableName()).Where(sql, value...).Count(&models.MicroLink{})
    if err != nil {
        grpclog.Error("get link count err: ",err.Error())
        return
    }
    count = int(ret)
    return
}

func LinkDelete(links_id []string) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg ="执行出错"
    success = false
    for _,value := range links_id{
        link :=models.MicroLink{}
        if _,err :=engine.Table(models.MicroLink{}.TableName()).ID(value).Delete(&link); err != nil{
            grpclog.Error("delete link failed !",err.Error())
            msg = value + "删除链接失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}

func LinkStatusChange(links_id []string, links_status int) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    link :=models.MicroLink{}
    success = false
    msg = "操作出错"
    for _,value := range links_id{
        link.Status = links_status
        link.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroLink{}.TableName()).ID(value).Cols("status").Update(&link); err != nil{
            grpclog.Error("update link failed ", err.Error())
            msg = value + "链接状态更新失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}
func GetLinkInfo(link_id string)(link_info models.MicroLink){
    engine := db.NewDb(db.DefaultNAME).Engine()
    link_info = models.MicroLink{}
    engine.Table(models.MicroLink{}.TableName()).ID(link_id).Get(&link_info)
    description, _ := base64.StdEncoding.DecodeString(link_info.Description)
    link_info.Description = string(description)
    return
}
func LinkEdit(link_data utils.LinkPostData) (success bool,msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    link :=models.MicroLink{
        Name: link_data.LinkName,
        UpdateTime: nowTime,
        Image: link_data.LinkImage,
        Url: link_data.LinkUrl,
        Description: base64.StdEncoding.EncodeToString([]byte(link_data.LinkDescription)),
    }
    var link_id = link_data.LinkId
    if link_id >0{
        if  _,err := engine.Table(models.MicroLink{}.TableName()).ID(link_id).Update(&link); err != nil{
            grpclog.Error("Link update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        link.CreateTime = nowTime
        link.Status = 1
        if  _,err := engine.Table(models.MicroLink{}.TableName()).Insert(&link); err != nil{
            grpclog.Error("Link insert:",err.Error())
            return
        }
    }
    success = true
    return
}
