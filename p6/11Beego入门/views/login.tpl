<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>登录</title>
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
</head>

<body>
    <div class="login_logo">
        <img src="/static/img/logo.png" alt="">
    </div>
    <form class="login_form" name="login" method="post" action="/login">
        <h1 class="login_title">用户登录</h1>
        <input type="text" class="input_txt" name="userName" placeholder="用户名">
        <input type="password" name="pwd" class="input_txt"  placeholder="密码">
        <div class="remember"><input type="checkbox" name="remember"><label>记住用户名</label></div>
        <input type="submit" value="登 录" class="input_sub">
        <a href="/register" style="
        color: #6c6c6c;
        cursor: pointer;
        float:right;
        margin-right:10%;
        margin-top: 4%; ">新用户注册</a>
    </form>
    <div class="login_bg"></div>
</body>

</html>