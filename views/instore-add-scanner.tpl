{{template "header"}}
<title>pay - Yuansfer</title>
</head>

<body onload="setCallback()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/instore-add-scanner">
            <div class="form-group">
                <label class="col-lg-4 control-label">商户号：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">店铺号：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">店员账号：</label>
                <div class="col-lg-6">
                    <input id="storeAdminNo" class="form-control" name="storeAdminNo" placeholder="StoreAdminNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">金额：</label>
                <div class="col-lg-6">
                    <input id="amt" class="form-control" name="amt" placeholder="amount">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">人民币金额：</label>
                <div class="col-lg-6">
                    <input id="rmbAmt" class="form-control" name="rmbAmt" placeholder="rmbAmt">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Reference：</label>
                <div class="col-lg-6">
                    <input id="reference" class="form-control" name="reference" value="{{.reference}}">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Description：</label>
                <div class="col-lg-6">
                    <input id="description" class="form-control" name="description" placeholder="description" value="{{.reference}}">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Note：</label>
                <div class="col-lg-6">
                    <input id="note" class="form-control" name="note" placeholder="note" value="{{.reference}}">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">IpnUrl：</label>
                <div class="col-lg-6">
                    <input id="ipnUrl" class="form-control" name="ipnUrl" placeholder="IpnUrl">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/instore-add-scanner">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">提交</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/instore-add-scanner/";
                        return false;
                    }

                    function checkInput() {
                        // var amt = document.getElementById("amt");
                        // if (amt.value.length == 0) {
                        //  alert("请输入金额");
                        //  return false;
                        // }

                        // var rmbAmt = document.getElementById("rmbAmt");
                        // if (rmbAmt.value.length == 0) {
                        //  alert("请输入金额");
                        //  return false;
                        // }

                        var vendor = document.getElementById("vendor");
                        if (vendor.value.length == 0) {
                            alert("请输入支付方式");
                            return false;
                        }
                    }
                    </script>
                </div>
            </div>
        </form>
    </div>
</body>

</html>