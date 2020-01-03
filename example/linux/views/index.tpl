<!DOCTYPE html>
<html>
 <head> 
  <title>Prometheus Alert System</title> 
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <link href="/static/css/bootstrap.min.css" rel="stylesheet">
   <script src="/static/js/jquery.js"></script>
   <script src="/static/js/bootstrap.min.js"></script>
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    .footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
      text-align:center;
    }

    .logo {
      background-image: url('https://avatars1.githubusercontent.com/u/3380462?s=200&v=4');
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 200px 0 20px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }
    .box {
      text-align: center;
      font-size: 16px;
      height: 100px;
    }
    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style> 
 </head> 
 <body class="backdrop"> 
  <header>
   <h1 class="logo">Welcome to Prometheus Alert System</h1>
   <div class="box">
     <p>PS:短信测试和电话测试请提前配置好app.conf中的[defaultphone]配置项</p>
     <button class="btn btn-primary" data-toggle='modal' id="dd">钉钉告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="wx">微信告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="txdx">腾讯云短信告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="txdh">腾讯云电话告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="hwdx">华为云短信告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="alydx">阿里云短信告警测试</button>
     <button class="btn btn-primary" data-toggle='modal' id="alydh">阿里云电话告警测试</button>
   </div>
   <div class="description">
    <a href="https://github.com/feiyu563/PrometheusAlert">Go to My GitHub and find how to use it or get new version !</a>
    <p>  Contact me: </p>
    <p> <img height="10%" width="10%" src="/static/img/wx.png" /> </p>
   </div>
  </header>
  <script src="/static/js/reload.min.js"></script>
   <script>
     $(function(){
        $(".btn").click(function(){
          $.post("/alerttest",{mtype:$(this).attr("id")},function(result){
            alert(result);
            });
         });
      })
   </script>
 </body>
</html>