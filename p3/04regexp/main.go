package main

import (
	"fmt"
	"regexp"
)

func main() {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta name="renderer" content="webkit" />
	<meta httpEquiv="X-UA-Compatible" content="IE=edge" />
	<title>福建烟草网 | 新商盟卷烟订货登录</title>
	<meta name="keywords" content="在线订货,烟草订货,网上订货" />
	<meta name="description" content="福建烟草官网:打造九地市卷烟在线订货集成窗口,直通新商盟网上订货平台,轻松进行烟草订货,省去卷烟订货繁复过程.发布福建烟草政策、烟草资讯、零售业界动态." />
	<link rel="stylesheet" href="css/xsm.css" />
	<link rel="stylesheet" href="xsmcs.css" />
	<link rel="stylesheet" href="topb.css" />
	
	<!--[if IE]><link type="text/css" rel="stylesheet" href="topb_ie.css" /><![endif]-->
	<base target="_blank" />
	</head>
	<body>
	<div class="xsm_box clearfix">
		<a class="xsm_box1" href="/" title="点击返回福建烟草网">
			<p class="xsm_text">
				返回福建烟草网首页
			</p>
		</a>
	</div>
	<div class="content">
		<div class="xsm_box4 clearfix">
			<iframe src="ind_AD4.html" width="100%" height="290px" scrolling="no" frameborder="0" marginwidth="0" marginheight="0" style="overflow:hidden;">
			 </iframe>
		</div>
		<form action="" id="loginform">
			<div id="form-desc"></div>
			<div id="label-username">登录名</div>
			<input id="username" autocomplete="off" type="text" maxlength="16" value="" tabindex="1" placeholder="用户名" title="请输入在线订货用户名" autofocus />
			<div class="content_right_jz clearfix">
				<div id="label-password">登录密码</div>
				<a id="findpwdlink">忘记密码？</a>
			</div>
			<input id="password" maxlength="30" type="password" autocomplete="off" value="" tabindex="2" placeholder="密码" title="请输入在线订货密码" />
			<div class="valcode-area clearfix valcodeHide">
				<input id="valcode" name="valcode" type="text" maxlength="4" tabindex="3" />
				<img id="valcodeimg">
				<input id="submit" type="submit" value="" />
				<a id="changevalcode">换一张</a>
			</div>
			<div class="content_right_jizhu">
				<label id="lable-remember" for="rememberun"><input type="checkbox" class="checkbox" id="rememberun" tabindex="4" title="是否记住在线订货的帐号和密码" />记住账号</label>
			</div>
	<div style="text-align:center">
			<input id="loginbtn"  type="button" title="登录新商盟订货平台" value="登 录" tabindex="5" />
	</div>
		<p class="xsm_text6">
				<strong>
					提示：
				</strong>
				新商盟用户由此登录订烟
			</p>
		</form>
	</div>
	<div class="contentbotton"></div>
	<div class="xsm_box13 clearfix">
					<p class="xsm_text7">
						订货平台登录
					</p>
					<div class="xsm_box14 clearfix">
						<a href="http://www.xinshangmeng.com" target="_blank" title="登录福州在线订货平台">
							<p>
								福州
							</p>
						</a>
						<a href="http://www.xinshangmeng.com" target="_blank" title="登录厦门在线订货平台">
							<p>
								厦门
							</p>
						</a>
						<a href="http://www.xinshangmeng.com" target="_blank" title="登录宁德在线订货平台">
							<p>
								宁德
							</p>
						</a>
						<a href="http://www.xinshangmeng.com" target="_blank" title="登录莆田在线订货平台">
							<p>
								莆田
							</p>
						</a>
						<a href="http://www.fjycgs.com" target="_blank" title="登录泉州在线订货平台">
							<p>
								泉州
							</p>
						</a>
						<a href="http://www.fjycgs.com" target="_blank" title="登录漳州在线订货平台">
							<p>
								漳州
							</p>
						</a>
						<a href="http://www.fjycgs.com/" target="_blank" title="登录龙岩在线订货平台">
							<p>
								龙岩
							</p>
						</a>
						<a href="http://www.fjycgs.com/" target="_blank" title="登录三明在线订货平台">
							<p>
								三明
							</p>
						</a>
						<a href="http://www.xinshangmeng.com" target="_blank" title="登录南平在线订货平台">
							<p>
								南平
							</p>
						</a>
					</div>
				</div>
				<form name="form1" method="post" action="dh.aspx" id="form1">
	<div>
	<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKLTI4MTc0NzE2NQ9kFgICAQ9kFgICAw8WAh4JaW5uZXJodG1sBaUFPGxpPjxhIGhyZWY9Ii9OZXdzLzIwMjAwMS8yMDIwMDEyMDU2MzcwMS5zaHRtbCIgdGFyZ2V0PSJfYmxhbmsiPjxwIGNsYXNzPSJ4c21fdGV4dDExIj7Ct+eOsOS7o+mbtuWUrue7iOerr+WuouaIt+KAnOWFreimgeS8muKAnTwvcD48cCBjbGFzcz0ieHNtX3RleHQxMiI+MjAyMC0wMS0yMDwvcD48L2E+PC9saT48bGk+PGEgaHJlZj0iL05ld3MvMjAyMDAxLzIwMjAwMTIwNTYzNjk4LnNodG1sIiB0YXJnZXQ9Il9ibGFuayI+PHAgY2xhc3M9InhzbV90ZXh0MTEiPsK35aaC5L2V5oiQ5Li65a6i5oi355qE4oCc5L6d6LWW5Lq64oCdPC9wPjxwIGNsYXNzPSJ4c21fdGV4dDEyIj4yMDIwLTAxLTIwPC9wPjwvYT48L2xpPjxsaT48YSBocmVmPSIvTmV3cy8yMDIwMDEvMjAyMDAxMjA1NjM2OTcuc2h0bWwiIHRhcmdldD0iX2JsYW5rIj48cCBjbGFzcz0ieHNtX3RleHQxMSI+wrflpb3mlLvnlaXliqnnm4jliKnoioLoioLljYc8L3A+PHAgY2xhc3M9InhzbV90ZXh0MTIiPjIwMjAtMDEtMjA8L3A+PC9hPjwvbGk+PGxpPjxhIGhyZWY9Ii9OZXdzLzIwMjAwMS8yMDIwMDEyMDU2MzY5Ni5zaHRtbCIgdGFyZ2V0PSJfYmxhbmsiPjxwIGNsYXNzPSJ4c21fdGV4dDExIj7Ct+aXuuWto+e7j+iQpeKAnOmHkeeCueWtkOKAnTwvcD48cCBjbGFzcz0ieHNtX3RleHQxMiI+MjAyMC0wMS0yMDwvcD48L2E+PC9saT5kZE7JvaRGGh8zj3HO+hD3y5DSls1Y" />
	</div>
	
	<div class="xsm_box15 clearfix">
		<div class="xsm_box16 clearfix">
			<div class="xsm_box17 clearfix">
				<p class="xsm_text9">
					零售要闻
				</p>
				<a class="xsm_text10" href="/NewsList.aspx?GUID=48-48-50-48-48-49-48-48-49" target="_blank" title="更多烟草资讯">
					更多
				</a>
			</div>
			<div class="xsm_box18 clearfix">
				<ul id="div002001003"><li><a href="/News/202001/20200120563677.shtml" target="_blank"><p class="xsm_text11">·去年我省网络零售额全国第六</p><p class="xsm_text12">2020-01-20</p></a></li><li><a href="/News/202001/20200119563617.shtml" target="_blank"><p class="xsm_text11">·茅台19家区域KA卖场投放飞天茅台迎战春节服务期</p><p class="xsm_text12">2020-01-19</p></a></li><li><a href="/News/202001/20200115563374.shtml" target="_blank"><p class="xsm_text11">·苏宁易购联手斯柯达，新型汽车门店正式上线</p><p class="xsm_text12">2020-01-15</p></a></li><li><a href="/News/202001/20200113563160.shtml" target="_blank"><p class="xsm_text11">·再拓商超渠道 茅台新增10万瓶53度飞天直供沃尔玛</p><p class="xsm_text12">2020-01-13</p></a></li></ul>
			</div>
		</div>
		<div class="xsm_box16b clearfix">
			<div class="xsm_box17 clearfix">
				<p class="xsm_text9">
					经营技能
				</p>
				<a class="xsm_text10" href="/NewsList.aspx?GUID=48-48-50-48-48-49-48-48-50" title="更多零售新闻">
					更多
				</a>
			</div>
			<div class="xsm_box18 clearfix">
				<ul id="div12"><li><a href="/News/202001/20200120563701.shtml" target="_blank"><p class="xsm_text11">·现代零售终端客户“六要会”</p><p class="xsm_text12">2020-01-20</p></a></li><li><a href="/News/202001/20200120563698.shtml" target="_blank"><p class="xsm_text11">·如何成为客户的“依赖人”</p><p class="xsm_text12">2020-01-20</p></a></li><li><a href="/News/202001/20200120563697.shtml" target="_blank"><p class="xsm_text11">·好攻略助盈利节节升</p><p class="xsm_text12">2020-01-20</p></a></li><li><a href="/News/202001/20200120563696.shtml" target="_blank"><p class="xsm_text11">·旺季经营“金点子”</p><p class="xsm_text12">2020-01-20</p></a></li></ul>
			</div>
		</div>
		
	</div>
	<div style="width:100%;">
	<div class="topb_box12 clearfix">
		<p>
		<a href="/about_company.aspx">闽烟概况</a>&nbsp;&nbsp;|&nbsp;&nbsp;
		<a href="/other/pubinfo_webmap.aspx" target="_blank">网站地图</a>&nbsp;&nbsp;|&nbsp;&nbsp;
		<a href="http://vip.fjycw.com/site/three" target="_blank">会员登录</a>&nbsp;&nbsp;|&nbsp;&nbsp;
		<a href="/other/pubinfo_lxwm.aspx" target="_blank">联系我们</a>
	<!--    <a href="http://www.gov.cn/c125533/jbxxk.htm" target="_blank">我为政府网站找错</a>
	-->    </p>
	</div>
	<div class="topb_box13 clearfix">
		<p>
		本站面向烟草业和成年烟民开放，含有烟草内容，如果您是18岁以下人士，敬请回避<br /><strong>主管：福建省烟草专卖局（公司）办公室　 承办：福建省海晟文化传媒有限公司</strong><br />地址：福建省福州市北环中路133号 邮编：350003<br/><strong>专卖举报：12313  服务投诉：0591-87069114  网站纠错：0591-87069450</strong><br>增值电信业务经营许可证：闽B2-20090018　闽ICP备06036078号-2　网站标识码：bm64130001<br />Copyright©2017-2021 福建烟草网 版权所有
		</p>
		
		<!--网站验证图标-->
		<div class="topb_box14 clearfix" style="margin-bottom:20px;">
			<a href="http://www.tobacco.gov.cn" target="_blank">
			<div class="topb_box15 clearfix">
				<img src="/img/inu11.jpg"/>
				<p>
				国家烟草<br>专卖局网站
				</p>
			</div>
			</a>
			<!--<a href="http://www.fjaic.gov.cn/" target="_blank">
			<div class="topb_box15 clearfix">
				<img src="/img/inu12.gif"/>
				<p>
				经营性网站<br>备案信息
				</p>
			</div>
			</a>-->
			<a href="http://www.12377.cn/" target="_blank">
			<div class="topb_box15 clearfix">
				<img src="/img/inu13.gif"/>
				<p>
				不良信息<br>举报中心
				</p>
			</div>
			</a>
			<a href="http://www.gov.cn/c125533/jbxxk.htm" target="_blank">
			<div class="topb_box15b clearfix">
				<img class="topb_image6 image" src="/img/inu15.png"/>
			</div>
			</a>
			<a href="http://biaozhi.conac.cn/" target="_blank">
			<div class="topb_box15 clearfix">
				<img src="/img/dzjg.png"/>
				<p>
				党政机关<br>网站标识
				</p>
			</div>
			</a>
		</div>
		<!--网站验证图标-end-->
		<iframe id="ifcount" src="" style="display:none; height:0px; width:0px;"></iframe>
		<div style="display:none;">
			<script type="text/javascript">var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://"); document.write(unescape("%3Cspan id='cnzz_stat_icon_1258291724'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s11.cnzz.com/stat.php%3Fid%3D1258291724%26show%3Dpic1' type='text/javascript'%3E%3C/script%3E"));</script>
			<script type="text/javascript">
	
				var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://"); document.write(unescape("%3Cspan id='cnzz_stat_icon_1770494'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s5.cnzz.com/stat.php%3Fid%3D1770494%26show%3Dpic' type='text/javascript'%3E%3C/script%3E"));
	
	
				//baidu
				var _hmt = _hmt || [];
				(function () {
					var hm = document.createElement("script");
					hm.src = "//hm.baidu.com/hm.js?fddf33d052f5f53d3b4b95fe92b20b65";
					var s = document.getElementsByTagName("script")[0];
					s.parentNode.insertBefore(hm, s);
				})();
	
				var hrt = location.href;
				var hostName = document.domain;
				var prot = document.location.protocol;
				var urlstr = prot + "//" + hostName + "/";
				hrt = hrt.replace(urlstr, "");
				//document.getElementById("ifcount").src="/ct.aspx?GUID="+hrt;
			</script>
		</div>
	</div>
	</div>
	</form>
	<script type="text/javascript" src="js/jquery-1.7.1.min.js"></script>
	<script type="text/javascript" src="js/swfobject.js"></script>
	<script type="text/javascript" src="js/index.js" charset="UTF-8"></script>
	<script type="text/javascript" src="js/config.js"></script>
	</body>
	</html>
	
	`
	// 题目一：提取出title
	pattern := regexp.MustCompile(`<title>(.*?)</title>`)
	str := pattern.FindStringSubmatch(html)
	fmt.Println(str[1])

	// 题目二：提取出所有链接
	pattern = regexp.MustCompile(`href=".*?\.[a-z]+"`)
	s2 := pattern.FindAllString(html, -1)
	// 通过字典去重后输出
	x := make(map[string]int)
	for _, v := range s2 {
		x[v] = 0
	}
	for k := range x {
		fmt.Println(k)
	}
}
