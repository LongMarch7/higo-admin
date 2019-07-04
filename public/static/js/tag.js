
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
        table.reload('tagListRender', {
            page: {
                curr: 1 //重新从第 1 页开始
            }
            ,where: data.field
        });
        layer.close(loading);
        return false
    })

    table.on('toolbar(tagListAction)', function(obj){
        var checkStatus = table.checkStatus(obj.config.id);
        var data = checkStatus.data;
        switch(obj.event){
            case 'addTag':
                editTag(0, '');
                break;
            case 'startTags':
                changeTagStatusWithBatch(data,1);
                break;
            case 'stopTags':
                changeTagStatusWithBatch(data,2);
                break;
            case 'deleteTags':
                deleteTagWithBatch(data)
                break;
        };
    });

    //监听状态操作
    form.on('switch(statusSwitch)', function(obj){
        changeTagStatus(this.value, obj)
    });

    table.on('tool(tagListAction)', function(obj){
        switch(obj.event){
            case 'delete':
                deleteTag(obj);
                break;
            case 'edit':
                editTag(obj.data.id, obj.data.name);
                break;
        }
    });


    function deleteTag(obj) {
        var data = obj.data;
        layer.confirm('确定删除吗', function(index){
            $.ajax({
                url: '/admin/portal/tag_delete',
                data: {"id": data.id},
                type: "post",
                dataType: "json",
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code === 0) {
                        message = ret.msg;
                        obj.del();
                        layer.close(index);
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {
                    });
                }
            });
        });
    }
    function deleteTagWithBatch(data) {
        var tags = [];
        data.forEach(function(value,i) {
            tags.unshift(value.id)
        });
        if (tags.length >0 ){
            layer.confirm('确定删除吗', function(index){
                $.ajax({
                    url: '/admin/portal/tag_delete',
                    data: {"id": tags},
                    type: "post",
                    dataType: "json",
                    traditional: true,
                    success: function (ret) {
                        var message = ret.msg + ret.code;
                        if (ret.code === 0) {
                            message = ret.msg
                            table.reload('tagListRender')
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

    function editTag(id,name){
        layer.open({
            type: 1
            ,title: false //不显示标题栏
            ,closeBtn: false
            ,area: '300px;'
            ,shade: 0.8
            ,id: 'edit_tag' //设定一个id，防止重复弹出
            ,btn: ['提交', '返回']
            ,btnAlign: 'c'
            ,moveType: 1 //拖拽模式，0或者1
            ,content: '<div style="padding: 50px; line-height: 22px; background-color: #393D49; color: #fff; font-weight: 300;">' +
                '<label for="tag_name">' +
                '                    <span class="x-red">*</span>标签名称' +
                '                </label>' +
                '<div class="layui-input-inline">\n' +
                '    <input type="text" id="tag_name" name="tag_name" autocomplete="off" class="layui-input" value=' + name + '>' +
                '</div>' +
                '</div>'
            ,yes: function(index, layero){
                tag_name = $("#tag_name").val()
                $.ajax({
                    url: '/admin/portal/tag_edit_post',
                    data: {"tag_id": id,"tag_name":tag_name},
                    type: "post",
                    dataType: "json",
                    traditional: true,
                    success: function (ret) {
                        var message = ret.msg + ret.code;
                        if (ret.code ===0) {
                            message = ret.msg;
                            layer.close(index);
                            table.reload('tagListRender');
                        }
                        layer.msg(message, {icon: 1, time: 1000}, function () {});
                    }
                });
            }
        });
    }

    function changeTagStatusWithBatch(data, status) {
        var tags = [];
        data.forEach(function(value,i) {
            tags.push(value.id)
        });
        if (tags.length >0 ){
            $.ajax({
                url: '/admin/portal/tag_status_change',
                data: {"id": tags,"status":status},
                type: "post",
                dataType: "json",
                traditional: true,
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code ===0) {
                        message = ret.msg
                        table.reload('tagListRender')
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {});
                }
            });
        }else{
            layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
        }
    }
});

function changeTagStatus(id,obj) {
    var status = 2;
    var title= "不发布";
    var cla="layui-icon layui-icon-play";
    if($(obj).attr('title')=='不发布'){
        title = "发布"
        cla = "layui-icon layui-icon-pause x-green"
        status = 1
    }
    $.ajax({
        url: '/admin/portal/tag_status_change',
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
