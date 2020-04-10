<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>编辑文章类型</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
</head>

<body>

    <div class="header">
        <a href="#" class="logo fl"><img src="/static/images/logo.png" alt="logo"></a>
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
            当前位置：文章管理>添加分类
        </div>
        <div class="pannel">
            <table class="common_table">
                <tr>
                    <th width="10%">id</th>
                    <th width="70%">类别名称</th>
                    <th width="20%">管理操作</th>
                </tr>
                {{range .types}}
                <tr>
                    <td>{{.Id}}</td>
                    <td>{{.TypeName}}</td>
                    <td><a href="javascript:;" class="edit">删除</a><a href="javascript:;" class="edit">编辑</a></td>
                </tr>
                {{end}}

                <tr>
                    <td colspan="3">
                        <form action="/addType" method="POST">
                            <input type="text" class="type_txt" placeholder="添加分类" name="typeName">
                            <input type="submit" class="addtype" value="增加分类">
                        </form>
                    </td>
                </tr>
            </table>
        </div>
    </div>



</body>

</html>