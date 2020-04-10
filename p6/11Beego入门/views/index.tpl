<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>后台管理页面</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
    <script type="text/javascript" src="/static/js/jquery-1.12.4.min.js"></script>
    <script type="text/javascript">
        window.onload = function () {
            $(".dels").click(function () {
                if (confirm("是否确认删除?")) {
                    return true
                }
            })
        }
    </script>

</head>

<body>

    <div class="header">
        <a href="/index" class="logo fl"><img src="/static/images/logo.png" alt="logo"></a>
        <a href="/login" class="logout fr">退 出</a>
    </div>

    <div class="side_bar">
        <div class="user_info">
            <img src="/static/images/person.png" alt="张大山">
            <p>欢迎你 <em>李雷</em></p>
        </div>

        <div class="menu_con">
            <div class="first_menu active"><a href="javascript:;" class="icon02">文章管理</a></div>
            <ul class="sub_menu show">
                <li><a href="/index" class="icon031">文章列表</a></li>
                <li><a href="/add" class="icon032">添加文章</a></li>
                <li><a href="/addType" class="icon034">添加分类</a></li>
            </ul>
        </div>
    </div>

    <div class="main_body" id="main_body">
        <div class="breadcrub">
            当前位置：文章管理>文章列表
        </div>
        <div class="pannel">
            <span class="sel_label">请选择文章分类：</span>
            <select name="select" id="select" class="sel_opt">
                {{range .articleType}}
                <option>{{.TypeName}}</option>
                {{end}}
            </select>

            <table class="common_table">
                <tr>
                    <th width="43%">文章标题</th>
                    <th width="10%">文章内容</th>
                    <th width="16%">添加时间</th>
                    <th width="7%">阅读量</th>
                    <th width="7%">删除</th>
                    <th width="7%">编辑</th>
                    <th width="10%">文章类型</th>
                </tr>


                {{range .articles}}
                <tr>
                    <td>{{.Aname}}</td>
                    <td><a href="/content?id={{.Id}}">查看详情</a></td>
                    <td>{{.Atime.Format "2006-01-02 15:04:05"}}</td>
                    <td>{{.Acount}}</td>
                    <td><a href="/delete?id={{.Id}}" class="dels">删除</a></td>
                    <td><a href="/update?id={{.Id}}">编辑</a></td>
                    <td>{{.ArtiType.TypeName}}</td>
                </tr>
                {{end}}
            </table>

            <ul class="pagenation">
                <li><a href="/index">首页</a></li>
                {{if compare .firstPage true}}
                <li>上一页</li>
                {{else}}
                <li><a href="/index?pageIndex={{.pageIndex| PrePage}}">上一页 </a></li>
                {{end}}
                {{if compare .lastPage true}}
                <li>下一页</li>
                {{else}}
                <li><a href="/index?pageIndex={{NextPage .pageIndex .pageCount}}">下一页</a></li>
                {{end}}
                <li><a href="/index?pageIndex={{.pageCount}}">末页</a></li>
                <li>共{{.count}}条记录/共{{.pageCount}}页/当前{{.pageIndex}}页</li>
            </ul>
        </div>
    </div>
</body>

</html>