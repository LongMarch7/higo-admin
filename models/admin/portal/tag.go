package portal

import (
    "github.com/LongMarch7/higo-admin/models/utils"
    "github.com/LongMarch7/higo/db"
    "github.com/LongMarch7/higo/util/define"
    "google.golang.org/grpc/grpclog"
    "strconv"
    "time"
    "github.com/LongMarch7/higo-admin/db/object/models"
)

func GetSysTagListWhereSql(where map[string]string) (string, []interface{}) {
    var sql = "1 = 1"
    var value []interface{}
    if v, ok := where["tag_name"]; ok && v != "" {
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

func GetTagList(where map[string]string, page int, rows int)  (tag_list []models.MicroTag, count int){
    engine := db.NewDb(db.DefaultNAME).Engine()
    count = 0
    sql , value:= GetSysTagListWhereSql(where)
    tag_list = make([]models.MicroTag, 0)
    err := engine.Table(models.MicroTag{}.TableName()).Where(sql, value...).Limit(rows,(page-1)*rows).
        Find(&tag_list)
    if err != nil {
        grpclog.Error("get tag list err: ",err.Error())
        return
    }
    ret,err :=engine.Table(models.MicroTag{}.TableName()).Where(sql, value...).Count(&models.MicroLink{})
    if err != nil {
        grpclog.Error("get tag count err: ",err.Error())
        return
    }
    count = int(ret)
    return
}

func TagDelete(tags_id []string) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg ="执行出错"
    success = false
    for _,value := range tags_id{
        link :=models.MicroTag{}
        if _,err :=engine.Table(models.MicroTag{}.TableName()).ID(value).Delete(&link); err != nil{
            grpclog.Error("delete tag failed !",err.Error())
            msg = value + "删除标签失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}

func TagStatusChange(tags_id []string, tags_status int) (success bool, msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    link :=models.MicroTag{}
    success = false
    msg = "操作出错"
    for _,value := range tags_id{
        link.Status = tags_status
        link.UpdateTime = int(time.Now().Unix())
        if _,err :=engine.Table(models.MicroTag{}.TableName()).ID(value).Cols("status").Update(&link); err != nil{
            grpclog.Error("update tag failed ", err.Error())
            msg = value + "标签状态更新失败"
            return
        }
    }
    msg = "操作成功"
    success = true
    return
}
func TagEdit(tag_data utils.TagPostData) (success bool,msg string){
    engine := db.NewDb(db.DefaultNAME).Engine()
    msg = "执行错误"
    success = false

    nowTime := int(time.Now().Unix())
    tag :=models.MicroTag{
        Name: tag_data.TagName,
        UpdateTime: nowTime,
    }
    var tag_id = tag_data.TagId
    if tag_id >0{
        if  _,err := engine.Table(models.MicroTag{}.TableName()).ID(tag_id).Update(&tag); err != nil{
            grpclog.Error("Tag update:",err.Error())
            msg = "更新出错"
            return
        }
    }else{
        has,err :=engine.Table(models.MicroTag{}.TableName()).Where("name = ?", tag_data.TagName).Exist(&models.MicroTag{})
        if err == nil && has{
            grpclog.Error("tag exist")
            msg = tag_data.TagName + "标签名已存在"
            return
        }
        tag.CreateTime = nowTime
        tag.Status = 1
        if  _,err := engine.Table(models.MicroTag{}.TableName()).Insert(&tag); err != nil{
            grpclog.Error("Tag insert:",err.Error())
            return
        }
    }
    success = true
    return
}
