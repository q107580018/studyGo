<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
</head>
<body>
<div class="login_logo">
        <img src="/static/images/logo.png" alt="">
    </div>  
    <form  class="login_form" name = "login" method = "post"  action="/register">     
        <h1 class="login_title">用户注册</h1>
        <input type="text" placeholder="用户名" class="input_txt" name="userName">
        <input type="password" placeholder="密码" class="input_txt" name = "pwd">
        <input type="submit" value="注 册" class="input_sub">
        <a href="/login" style="
        color: #6c6c6c;
        cursor: pointer;
        float:right;
        margin-right:10%;
        margin-top: 4%; ">已有账号</a>
    </form>
    <div class="login_bg"></div>
</body>
</html>