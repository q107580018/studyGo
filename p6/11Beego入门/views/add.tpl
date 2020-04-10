<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>添加文章内容</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
</head>

<body>
    <div class="header">
        <a href="#" class="logo fl"><img src="/static/images/logo.png" alt="logo"></a>
        <a href="#" class="logout fr">退 出</a>
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
            当前位置：文章管理>添加文章
        </div>
        <div class="pannel">
            <form method="POST" enctype="multipart/form-data" action="/add">
                <h3 class="review_title">添加文章</h3>
                <div class="form_group">
                    <label>文章标题：</label>
                    <input type="text" class="input_txt2" name="articleName" maxlength="20">
                </div>
                <div class="form_group">
                    <label>文章类型：</label>
                    <select class="sel_opt" name="select">
                        {{range .articleType}}
                        <option>{{.TypeName}}</option>
                        {{end}}
                    </select>
                </div>
                <div class="form_group">
                    <label>文章内容：</label>
                    <textarea class="input_multxt" name="content"></textarea>
                </div>
                <div class="form_group">
                    <label>上传图片：</label>
                    <input type="file" class="input_file" name="uploadname">
                </div>
                <div class="form_group indent_group line_top">
                    <input type="submit" value="添 加" class="confirm">
                    <span>{{.errmsg}}</span>
                </div>
            </form>
        </div>
    </div>

</body>

</html>