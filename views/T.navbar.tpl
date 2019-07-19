{{define "navbar"}}
<!-- <a class="navbar-brand" href="/">Yuansfer</a> -->
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="/">
            <img alt="Brand" src="http://oss.yuansfer.com/log_20190410.png?x-oss-process=image/resize,l_110" height="20">
        </a>
        <div>
            <ul class="nav navbar-nav">
                <li {{if .IsPay}}class="active" {{end}}><a href="/">Online Pay</a></li>
                <li {{if .IsInquire}}class="active" {{end}}><a href="/inquire">Online Query</a></li>
                <li {{if .IsRefund}}class="active" {{end}}><a href="/refund">Refund</a></li>
                <li {{if .IsInstoreAdd}}class="active" {{end}}><a href="/instore-add">In-store Add</a></li>
                <li {{if .IsInstoreAddScanner}}class="active" {{end}}><a href="/instore-add-scanner">In-store Add Scanner</a></li>
                <li {{if .IsInstorePay}}class="active" {{end}}><a href="/instore-pay" target="_blank">In-store Pay</a></li>
                <li {{if .IsInstoreDetail}}class="active" {{end}}><a href="/instore-detail" target="_blank">In-store Detail</a></li>
                <li {{if .IsCreateQrcode}}class="active" {{end}}><a href="/instore-qrcode">In-store Qrcode</a></li>
                <li {{if .IsReverse}}class="active" {{end}}><a href="/instore-reverse">In-store Reverse</a></li>
                <li {{if .IsMicroPay}}class="active" {{end}}><a href="/micropay">Micropay</a></li>
                <li {{if .IsRecurring}}class="active" {{end}}><a href="/recurring">Recurring</a></li>
                <li {{if .IsOnlineReverse}}class="active" {{end}}><a href="/online-reverse">Online Reverse</a></li>
                <li {{if .IsForced}}class="active" {{end}}><a href="/forced-auth">Forced Auth</a></li>
            </ul>
        </div>
        <!-- <div class="pull-right">
            <ul class="nav navbar-nav">
                <li>
                    <a class="navbar-brand" href="/">
                        <img alt="Brand" src="/static/img/7372_logo-invoice.png" height="20">
                    </a>
                </li>
            </ul>
            <a id="txt"></a>
        </div> -->
    </div>
</div>
<br/>
<br/>
<br/>
<br/>
<br/>
{{end}}