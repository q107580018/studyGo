<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>文章详情</title>
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
            当前位置：文章管理>文章列表>文章详情
        </div>
        <div class="pannel">
            <h3 class="review_title">文章详情</h3>
            <div class="form_group">
                <label>文章标题：</label>
                <p class="detail"><b>{{.article.Aname}}</b></p>
            </div>
            <div class="form_group">
                <label>文章类型：</label>
                <p class="detail">体育新闻</p>
            </div>
            <div class="form_group">
                <label>文章内容：</label>
                <p class="detail"><img src="">{{.article.Acontent}}</p>
            </div>
            <div class="form_group">
                <label>阅读次数：</label>
                <p class="detail">{{.article.Acount}}</p>
            </div>
            <div class="form_group">
                <label>创建时间：</label>
                <p class="detail">{{.article.Atime.Format "2006-01-02 15:04:05"}}</p>
                <span>{{.errmsg}}</span>
            </div>          
        </div>
    </div>

</body>
</html>