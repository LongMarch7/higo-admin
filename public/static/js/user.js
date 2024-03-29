
layui.use(['table', 'laydate', 'form'], function(){
    var table = layui.table
        ,form = layui.form
        ,laydate = layui.laydate;

    laydate.render({
        elem: '#start_time',
        type: 'datetime'
    });

    laydate.render({
        elem: '#end_time',
        type: 'datetime'
    });
    form.on('submit(sreach)', function (data) {
        var loading = layer.load(1, {shade: [0.1, '#FF0000']});

        table.reload('listRender', {
            page: {
                curr: 1 //重新从第 1 页开始
            }
            ,where: data.field
        });
        layer.close(loading);
        return false
    })

    table.on('toolbar(listAction)', function(obj){
        var checkStatus = table.checkStatus(obj.config.id);
        var data = checkStatus.data;
        switch(obj.event){
            case 'add':
                xadmin.open('添加用户','/admin/user/user_edit',600,500);
                break;
            case 'startBatch':
                changeStatusWithBatch(data,1);
                break;
            case 'stopBatch':
                changeStatusWithBatch(data,3);
                break;
            case 'deleteBatch':
                deleteWithBatch(data)
                break;
        };
    });

    table.on('tool(listAction)', function(obj){
        switch(obj.event){
            case 'delete':
                deleteObj(obj);
                break
            case 'edit':
                editObj(obj);
                break;
        }
    });


    function deleteObj(obj) {
        var data = obj.data;
        layer.confirm('确定删除吗', function(index){
            $.ajax({
                url: '/admin/user/user_delete',
                data: {"id": data.id,"name":data.user_login},
                type: "post",
                dataType: "json",
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code === 0) {
                        message = ret.msg
                        obj.del();
                        layer.close(index);
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {
                    });
                }
            });
        });
    }
    function deleteWithBatch(data) {
        var list = [];
        var list_name = [];
        data.forEach(function(value,i) {
            list.unshift(value.id)
            list_name.unshift(value.user_login)
        });
        if (list.length >0 ){
            layer.confirm('确定删除吗', function(index){
                $.ajax({
                    url: '/admin/user/user_delete',
                    data: {"id": list,"name":list_name},
                    type: "post",
                    dataType: "json",
                    traditional: true,
                    success: function (ret) {
                        var message = ret.msg + ret.code;
                        if (ret.code === 0) {
                            message = ret.msg
                            table.reload('listRender')
                        }
                        layer.msg(message, {icon: 1, time: 1000}, function () {
                        });
                    }
                });
            });
        }else{
            layer.msg("未选中用户", {icon: 1, time: 1000}, function () {});
        }
    }

    function editObj(obj){
        var data = obj.data;
        xadmin.open('编辑用户','/admin/user/user_edit?user_name='+data.user_login+'&user_id='+data.id,600,500);
    }

    function changeStatusWithBatch(data, status) {
        var list = [];
        data.forEach(function(value,i) {
            list.push(value.id)
        });
        if (list.length >0 ){
            $.ajax({
                url: '/admin/user/user_status_change',
                data: {"id": list,"status":status},
                type: "post",
                dataType: "json",
                traditional: true,
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code ===0) {
                        message = ret.msg
                        table.reload('listRender')
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {});
                }
            });
        }else{
            layer.msg("未选中用户", {icon: 1, time: 1000}, function () {});
        }
    }
});

function changeStatus(id,obj) {
    var status = 3;
    var title= "停用";
    var cla="layui-icon layui-icon-play";
    if($(obj).attr('title')=='停用'){
        title = "启用"
        cla = "layui-icon layui-icon-pause"
        status = 1
    }
    $.ajax({
        url: '/admin/user/user_status_change',
        data: {"id": id,"status":status},
        type: "post",
        dataType: "json",
        success: function (ret) {
            var message = ret.msg;
            if (ret.code < 0) {
                message += ret.code;
            }else{
                $(obj).attr('title',title)
                $(obj).find('i').attr("class", cla)
            }
            layer.msg(message, {icon: 1, time: 1000}, function () {});
        }
    });
}