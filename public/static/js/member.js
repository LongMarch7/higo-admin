
layui.use(['table', 'laydate', 'form'], function(){
    var table = layui.table
        ,form = layui.form
        ,$ = layui.jquery
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

        table.reload('userListRender', {
            page: {
                curr: 1 //重新从第 1 页开始
            }
            ,where: data.field
        });
        layer.close(loading);
        return false
    })

    table.on('toolbar(userListAction)', function(obj){
        var checkStatus = table.checkStatus(obj.config.id);
        var data = checkStatus.data;
        switch(obj.event){
            case 'addUser':
                xadmin.open('添加会员','/admin/member/edit',500,700);
                break;
            case 'startUsers':
                changeUserStatusWithBatch(data,1);
                break;
            case 'stopUsers':
                changeUserStatusWithBatch(data,3);
                break;
            case 'deleteUsers':
                deleteUserWithBatch(data)
                break;
        };
    });

    table.on('tool(userListAction)', function(obj){
        switch(obj.event){
            case 'delete':
                deleteUser(obj);
                break
            case 'edit':
                editUser(obj);
                break;
        }
    });


    function deleteUser(obj) {
        var data = obj.data;
        layer.confirm('确定删除吗', function(index){
            $.ajax({
                url: '/admin/member/delete',
                data: {"id": data.id},
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
    function deleteUserWithBatch(data) {
        var users = [];
        data.forEach(function(value,i) {
            users.unshift(value.id)
        });
        if (users.length >0 ){
            layer.confirm('确定删除吗', function(index){
                $.ajax({
                    url: '/admin/member/delete',
                    data: {"id": users},
                    type: "post",
                    dataType: "json",
                    traditional: true,
                    success: function (ret) {
                        var message = ret.msg + ret.code;
                        if (ret.code === 0) {
                            message = ret.msg
                            table.reload('userListRender')
                        }
                        layer.msg(message, {icon: 1, time: 1000}, function () {
                        });
                    }
                });
            });
        }else{
            layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
        }
    }

    function editUser(obj){
        var data = obj.data;
        xadmin.open('编辑会员','/admin/member/edit?id='+data.id,500,700);
    }

    function changeUserStatusWithBatch(data, status) {
        var users = [];
        data.forEach(function(value,i) {
            users.push(value.id)
        });
        if (users.length >0 ){
            $.ajax({
                url: '/admin/member/status_change',
                data: {"id": users,"status":status},
                type: "post",
                dataType: "json",
                traditional: true,
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code ===0) {
                        message = ret.msg
                        table.reload('userListRender')
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {});
                }
            });
        }else{
            layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
        }
    }
});

function changeUserStatus(id,obj) {
    var status = 3;
    var title= "停用";
    var cla="layui-icon layui-icon-play";
    if($(obj).attr('title')=='停用'){title = "启用"
        cla = "layui-icon layui-icon-pause"
        status = 1
    }
    $.ajax({
        url: '/admin/member/status_change',
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
