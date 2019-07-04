
layui.use(['table', 'form'], function(){
    var table = layui.table
        ,form = layui.form
        ,$ = layui.jquery;
    $("#cate").on('click','tr',function () {
        if($(this).attr('status')=='true'){
            $(this).attr('status','false');
            cateId = $(this).attr('cate-id');
            $("tbody tr[fid="+cateId+"]").show();
        }else{
            cateIds = [];
            $(this).attr('status','true');
            cateId = $(this).attr('cate-id');
            getCateId(cateId);
            for (var i in cateIds) {
                $("tbody tr[cate-id="+cateIds[i]+"]").hide();
            }
        }
    });
    form.on('checkbox(checkAll)', function(data){
        if (data.elem.checked == true){
            $("input[name='cateCheckBox']:not(:checked)").each(function (index, item) {
                item.checked = true
            });
        }else{
            $("input[name='cateCheckBox']:checked").each(function (index, item) {
                item.checked = false
            });
        }
        form.render('checkbox');
    });

    form.on('checkbox(cateCheck)', function(data){
        if (data.elem.checked == false){
            $("input[name='checkAllBox']:checked").each(function (index, item) {
                item.checked = false
            });
        }
        form.render('checkbox');
    });

});

function deleteCateWithBatch(action,tips) {
    var cates = [];
    $('input[type="checkbox"][name="cateCheckBox"]:checked').each(function () {
        cates.unshift($(this).val())
    });
    if (cates.length >0 ){
        layer.confirm(tips, function(index){
            $.ajax({
                url: '/admin/portal/cate_delete',
                data: {"id": cates,"action":action},
                type: "post",
                dataType: "json",
                traditional: true,
                success: function (ret) {
                    var message = ret.msg + ret.code;
                    if (ret.code === 0) {
                        message = ret.msg
                        // $('input[type="checkbox"][name="cateCheckBox"]:checked').each(function () {
                        //     $(this).closest('tr').remove()
                        // })
                    }
                    layer.msg(message, {icon: 1, time: 1000}, function () {});
                    window.location.reload();
                }
            });
        });
    }else{
        layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
    }
}
function changeCateStatusWithBatch(status) {
    var cates = [];
    $('input[type="checkbox"][name="cateCheckBox"]:checked').each(function () {
        cates.unshift($(this).val())
    });
    if (cates.length >0 ){
        $.ajax({
            url: '/admin/portal/cate_status_change',
            data: {"id": cates,"status":status},
            type: "post",
            dataType: "json",
            traditional: true,
            success: function (ret) {
                var message = ret.msg + ret.code;
                if (ret.code ===0) {
                    message = ret.msg
                }
                layer.msg(message, {icon: 1, time: 1000}, function () {});
                window.location.reload();
            }
        });
    }else{
        layer.msg("未选中目标", {icon: 1, time: 1000}, function () {});
    }
}
function changeCateStatus(id,obj) {
    var status = 2;
    var title= "不发布";
    var cla="layui-icon layui-icon-play";
    if($(obj).attr('title')=='不发布'){
        title = "发布"
        cla = "layui-icon layui-icon-pause x-green"
        status = 1
    }
    $.ajax({
        url: '/admin/portal/cate_status_change',
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

function editCate(title,id,p_id,level) {
    xadmin.open(title,'/admin/portal/cate_edit?id=' + id + '&parent_id='+p_id + '&level=' + level,600,500);
}

function deleteCate(id, action, tips,obj) {
    layer.confirm(tips, function(index){
        $.ajax({
            url: '/admin/portal/cate_delete',
            data: {"id": id, "action":action},
            type: "post",
            dataType: "json",
            success: function (ret) {
                var message = ret.msg + ret.code;
                if (ret.code === 0) {
                    message = ret.msg;
                    $(obj).closest('tr').remove()
                }
                layer.msg(message, {icon: 1, time: 1000}, function () {
                });
            }
        });
    });
}

$(function(){
    //$("tbody.x-cate tr[fid!='0']").hide();
    $('.x-table-view').click(function () {
        if($(this).attr('status')=='true'){
            $("tbody.x-cate tr[fid!='0']").hide();
            $(this).attr('status','false');
            $(this).html('&#xe623;');
            $(this).attr('title','全部展开');
        }else{
            $(this).attr('status','true');
            $(this).html('&#xe625;');
            $(this).attr('title','全部隐藏');
            $("tbody.x-cate tr[fid!='0']").show();
        }
    });
})

var cateIds = [];
function getCateId(cateId) {
    $("tbody tr[fid="+cateId+"]").each(function(index, el) {
        id = $(el).attr('cate-id');
        cateIds.push(id);
        getCateId(id);
    });
}
