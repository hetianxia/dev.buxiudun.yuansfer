{{define "header"}}
<!DOCTYPE html>

<html>
  	<head>
		<link rel="shortcut icon" href="/static/img/favicon.jpg" />
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		 <!-- Stylesheets -->
		<link href="/static/css/bootstrap.min.css" rel="stylesheet" />
        <script src="/static/js/bootstrap.min.js" > </script>

    <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
    <script src="http://res.wx.qq.com/open/js/jweixin-1.4.0.js"></script>
    <script>
        var noncestr ;
        var timestamp 
        var url ;
        var appid ;
        var signature ;
        $(document).ready(function(){
console.log(window.location.href.split('#')[0])
    var addr = "http://yinghe.io/wx/sign?addr=" + window.location.href.split('#')[0];
    $.get(addr,function(data,status){
      var json=JSON.parse(data);
      console.log("JSON:" + json);
      console.log("noncestr:" + json.noncestr);
      console.log("noncestr-0:", noncestr)
      console.log("timestamp:" + json.timestamp);
      console.log("link:" + json.url);
      console.log("imgUrl:" + json.url);
      console.log("signature:" + json.signature);
      console.log("appid:" + json.appid);
      noncestr = json.noncestr
      timestamp = json.timestamp
      url = json.url
      appid = json.appid
      signature = json.signature

    wx.config({
    debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
    appId: appid, // 必填，公众号的唯一标识
    timestamp: timestamp, // 必填，生成签名的时间戳
    nonceStr: noncestr, // 必填，生成签名的随机串
    signature: signature,// 必填，签名
    jsApiList: ['updateAppMessageShareData', 'updateTimelineShareData'] // 必填，需要使用的JS接口列表
});
         wx.ready(function(){

                wx.updateAppMessageShareData({
                    title: "Yuansfer Sandbox", // 分享标题
                    desc: "Yuansfer Sandbox", // 分享描述
                    link: window.location.href.split('#')[0], // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
                    imgUrl: "https://oss.yuansfer.com/images/marketing/icon512.jpg?x-oss-process=image/resize,h_40",
                    // imgUrl: '', // 分享图标
                    type: '', // 分享类型,music、video或link，不填默认为link
                    dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                    success: function () {
                        // alert("成功")
                        //console.log("updateAppMessageShareData success")
                        // 用户点击了分享后执行的回调函数
                    }
                });

                wx.updateTimelineShareData({
			        title: "Yuansfer Sandbox", // 分享标题
			        link: window.location.href.split('#')[0], // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
			        imgUrl: "https://oss.yuansfer.com/images/marketing/icon512.jpg?x-oss-process=image/resize,h_40", // 分享图标
			        success: function () {
			          // 设置成功
			        }
                })

            });
});
    });

</script>
{{end}}