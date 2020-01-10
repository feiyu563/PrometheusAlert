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
      width: 100%;
      margin-left: auto;
      margin-right: auto;
      text-align:center;
    }

    .logo {
      background-image: url('/static/img/prometheus-ico.png');
      background-repeat: no-repeat;
      -webkit-background-size: 70px 70px;
      background-size: 70px 70px;
      background-position: center top;
      text-align: center;
      font-size: 32px;
      padding: 100px 0 10px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }
    .box {
      text-align: center;
      height: 200px;
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
      height: 500px;
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
  </header>
  
   <div class="box">
     <table border="0" style="width: 60%;margin:auto;text-shadow: 0px 1px 2px #ddd;">
       <tr>
         <td><p>webhook</p></td>
         <td><p>短信</p></td>
         <td><p>电话</p></td>
       </tr>
       <tr>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="dd">钉钉告警测试</button></p></td>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="hwdx">华为云短信告警测试</button></p></td>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="txdh">腾讯云电话告警测试</button></p></td>
       </tr>
       <tr>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="wx">微信告警测试</button></p></td>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="txdx">腾讯云短信告警测试</button></p></td>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="alydh">阿里云电话告警测试</button></p></td>
       </tr>
       <tr>
         <td></td>
         <td><p><button class="btn btn-primary" data-toggle='modal' id="alydx">阿里云短信告警测试</button></p></td>
         <td></td>
       </tr>
       <tr>
         <td colspan="3"><p>PS:短信测试和电话测试请提前配置好app.conf中的[defaultphone]配置项</p></td>
       </tr>
     </table>
   </div>
   <div class="description">
    <a href="https://feiyu563.github.io">Go to My GitHub and find how to use it or get new version !</a>
    <p>  Contact me: </p>
    <p> <img height="10%" width="10%" src="/static/img/wx.png" /> </p>
   </div>
   
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