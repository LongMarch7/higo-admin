
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

        table.reload('roleListRender', {
            page: {
                curr: 1 //重新从第 1 页开始
            }
            ,where: data.field
        });
        layer.close(loading);
        return false
    })

    table.on('toolbar(roleListAction)', function(obj){
        var checkStatus = table.checkStatus(obj.config.id);
        var data = checkStatus.data;
        switch(obj.event){
            case 'addRole':
                xadmin.open('添加角色','/admin/user/role_edit');
                break;
            case 'startRoles':
                changeRoleStatusWithBatch(data,1);
                break;
            case 'stopRoles':
                changeRoleStatusWithBatch(data,2);
                break;
            case 'deleteRoles':
                deleteRoleWithBatch(data)
                break;
        };
    });

    table.on('tool(roleListAction)', function(obj){
        switch(obj.event){
            case 'delete':
                deleteRole(obj);
                break
            case 'edit':
                editRole(obj);
                break;
        }
    });


    function deleteRole(obj) {
        var data = obj.data;
        layer.confirm('确定删除吗', function(index){
            $.ajax({
                url: '/admin/user/role_delete',
                data: {"id": data.id,"name":data.role_name},
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
    function deleteRoleWithBatch(data) {
        var roles = [];
        var roles_name = [];
        data.forEach(function(value,i) {
            roles.unshift(value.id)
            roles_name.unshift(value.role_name)
        });
        if (roles.length >0 ){
            layer.confirm('确定删除吗', function(index){
                $.ajax({
                    url: '/admin/user/role_delete',
                    data: {"id": roles,"name":roles_name},
                    type: "post",
                    dataType: "json",
                    traditional: true,
                    success: function (ret) {
                        var message = ret.msg + ret.code;
                        if (ret.code === 0) {
                            message = ret.msg
                            table.reload('roleListRender')
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

    function editRole(obj){
        var data = obj.data;
        xadmin.open('添加角色','/admin/user/role_edit?role_parent='+data.parent_id+'&role_id='+data.id);
    }

    function changeRoleStatusWithBatch(data, status) {
        var roles = [];
        data.forEach(function(value,i) {
            roles.push(value.id)
        });
        if (roles.length >0 ){
            $.ajax({
                url: '/admin/user/role_status_change',
                data: {"id": roles,"status":status},
                type: "post",
                dataType: "json",
                traditional: true,
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code ===0) {
                        message = ret.msg;
                        table.reload('roleListRender')
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {});
                }
            });
        }else{
            layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
        }
    }
});


function changeRoleStatus(id,obj) {
    var status = 2;
    var title= "停用";
    var cla="layui-icon layui-icon-play";
    if($(obj).attr('title')=='停用'){
        title = "启用"
        cla = "layui-icon layui-icon-pause"
        status = 1
    }
    $.ajax({
        url: '/admin/user/role_status_change',
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
