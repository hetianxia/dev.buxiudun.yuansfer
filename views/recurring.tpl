{{template "header"}}
<title>recurring - Yuansfer</title>
</head>

<body onload="startTime()">
    <div class="container" style="width: 500px;">
        {{template "navbar" .}}
<!--  pay -->
        <form class="form-horizontal" method="POST" action="/recurring/pay">
            <div class="form-group">
                <label class="col-lg-4 control-label">merchantNo：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">storeNo：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">scheduleNo：</label>
                <div class="col-lg-6">
                    <input id="scheduleNo" class="form-control" name="scheduleNo" placeholder="scheduleNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">paymentDate：</label>
                <div class="col-lg-6">
                    <input id="paymentDate" class="form-control" name="paymentDate" placeholder="paymentDate">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/recurring/pay" target="_blank">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">提交</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/recurring";
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

                        //var vendor = document.getElementById("vendor");
                        //if (vendor.value.length == 0) {
                        //    alert("请输入支付方式");
                        //    return false;
                        //}
                    }
                    </script>
                </div>
            </div>
        </form>
<!--  search -->
        <form class="form-horizontal" method="POST" action="/recurring/search">
            <div class="form-group">
                <label class="col-lg-4 control-label">merchantNo：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">storeNo：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">scheduleNo：</label>
                <div class="col-lg-6">
                    <input id="scheduleNo" class="form-control" name="scheduleNo" placeholder="scheduleNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">paymentDate：</label>
                <div class="col-lg-6">
                    <input id="paymentDate" class="form-control" name="paymentDate" placeholder="paymentDate">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/recurring/search" target="_blank">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">查询</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/recurring";
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

                        //var vendor = document.getElementById("vendor");
                        //if (vendor.value.length == 0) {
                        //    alert("请输入支付方式");
                        //    return false;
                        //}
                    }
                    </script>
                </div>
            </div>
        </form>
<!--  update -->
        <form class="form-horizontal" method="POST" action="/recurring/update">
            <div class="form-group">
                <label class="col-lg-4 control-label">merchantNo：</label>
                <div class="col-lg-6">
                    <input id="merchantNo" class="form-control" name="merchantNo" placeholder="merchantNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">storeNo：</label>
                <div class="col-lg-6">
                    <input id="storeNo" class="form-control" name="storeNo" placeholder="storeNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">scheduleNo：</label>
                <div class="col-lg-6">
                    <input id="scheduleNo" class="form-control" name="scheduleNo" placeholder="scheduleNo">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">金额：</label>
                <div class="col-lg-6">
                    <input id="amt" class="form-control" name="amt" placeholder="amount">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Count</label>
                <div class="col-lg-6">
                    <input id="count" class="form-control" name="count" placeholder="Count">
                </div>
            </div>
            <div class="form-group">
                <label class="col-lg-4 control-label">Status</label>
                <div class="col-lg-6">
                    <input id="stat" class="form-control" name="stat" placeholder="Status">
                </div>
            </div>
            <div class="form-group">
                <div class="col-lg-offset-2 col-lg-10">
                    <a href="/recurring/update" target="_blank">
                        <button type="submit" class="btn btn-default" onclick="return checkInput();">修改</button>
                    </a>
                    <button class="btn btn-default" onclick="return backToHome();">返回</button>
                    <script type="text/javascript">
                    function backToHome() {
                        window.location.href = "/recurring/";
                        return false;
                    }

                    function checkInput() {
                    }
                    </script>
                </div>
            </div>
        </form>

    </div>
</body>

</html>