{{template "header"}}
<title>pay - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/instore-detail">
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
                <label class="col-lg-4 control-label">StoreAdminNo</label>
                <div class="col-lg-6">
                    <input id="storeAdminNo" class="form-control" name="storeAdminNo" placeholder="Store Admin No.">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">transactionNo：</label>
                <div class="col-lg-6">
                    <input id="transNo" class="form-control" name="transNo" placeholder="transactionNo">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/instore-detail">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">提交</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/instore-detail/";
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