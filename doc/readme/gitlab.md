# Gitlab接入配置

将Gitlab webhook event通过PrometheusAlert通知到其它软件。

<br/>

目前支持的gitlab event类型:

- Push
- Tag Push
- Issue
- Comment
- Merge Request
- Wiki Page
- Pipeline
- Job
- Deployment
- Feature Flag
- Release

<br/>

目前支持通知到:

- 企业微信机器人(`/gitlab/weixin?wxurl=xxx`)
- 钉钉机器人(`/gitlab/dingding?ddurl=xxx`)

地址不支持写多个，如需要多个，可创建多个gitlab webhoook。

<br/>
<br/>


## 配置步骤

进入gitlab仓库界面，在`设置->Webhooks`里，填写接口地址和机器人地址，选择触发来源事件，之后就可以测试了。

![gitlab配置](/doc/images/gitlab_setup.png)


<br/>
<br/>


## 展示效果

<br/>

### Push事件

微信机器人效果:

![push-weixin](/doc/images/gitlab_push_weixin.png)

钉钉机器人效果:

![push-dingding](/doc/images/gitlab_push_dingding.png)

