{{template "header"}}
<title>Forced Auth - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
        <form class="form-horizontal" method="POST" action="/forced-auth">
            <div class="form-group">
                <label class="col-lg-4 control-label">MerchantNo：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">StoreNo：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Amount：</label>
                <div class="col-lg-6">
                    <input id="amt" class="form-control" name="amt" placeholder="Amount">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">AuthCode：</label>
                <div class="col-lg-6">
                    <input id="authCode" class="form-control" name="authCode" placeholder="Auth Code">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">ExpirationDate：</label>
                <div class="col-lg-6">
                    <input id="expirationDate" class="form-control" name="expirationDate" placeholder="ExpirationDate">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">CardNumber：</label>
                <div class="col-lg-6">
                    <input id="cardNumber" class="form-control" name="cardNumber" placeholder="CardNumber">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">CardType</label>
                <div class="col-lg-6">
                    <input id="cardType" class="form-control" name="cardType" placeholder="cardType">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Address：</label>
                <div class="col-lg-6">
                    <input id="addr" class="form-control" name="addr" placeholder="Address">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Zip：</label>
                <div class="col-lg-6">
                    <input id="zip" class="form-control" name="zip" placeholder="Zip">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/forced-auth" target="_blank">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">提交</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/forced-auth";
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