package controllers

import (
	"encoding/json"
	"testing"

	"github.com/astaxie/beego"
)

func TestSendGitlabEvent(t *testing.T) {
	// 读取配置文件
	beego.LoadAppConfig("ini", "../conf/app.conf")
	// 构造一个测试json
	pushJSON := `
        {
          "object_kind": "push",
          "event_name": "push",
          "before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
          "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
          "ref": "refs/heads/master",
          "checkout_sha": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
          "message":null,
          "user_id": 4,
          "user_name": "John Smith",
          "user_username": "jsmith",
          "user_email": "john@example.com",
          "user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
          "project_id": 15,
          "project":{
            "id": 15,
            "name":"Diaspora",
            "description":"",
            "web_url":"http://example.com/mike/diaspora",
            "avatar_url":null,
            "git_ssh_url":"git@example.com:mike/diaspora.git",
            "git_http_url":"http://example.com/mike/diaspora.git",
            "namespace":"Mike",
            "visibility_level":0,
            "path_with_namespace":"mike/diaspora",
            "default_branch":"master",
            "homepage":"http://example.com/mike/diaspora",
            "url":"git@example.com:mike/diaspora.git",
            "ssh_url":"git@example.com:mike/diaspora.git",
            "http_url":"http://example.com/mike/diaspora.git"
          },
          "repository":{
            "name": "Diaspora",
            "url": "git@example.com:mike/diaspora.git",
            "description": "",
            "homepage": "http://example.com/mike/diaspora",
            "git_http_url":"http://example.com/mike/diaspora.git",
            "git_ssh_url":"git@example.com:mike/diaspora.git",
            "visibility_level":0
          },
          "commits": [
            {
              "id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
              "message": "Update Catalan translation to e38cb41.\n\nSee https://gitlab.com/gitlab-org/gitlab for more information",
              "title": "Update Catalan translation to e38cb41.",
              "timestamp": "2011-12-12T14:27:31+02:00",
              "url": "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
              "author": {
                "name": "Jordi Mallach",
                "email": "jordi@softcatala.org"
              },
              "added": ["CHANGELOG"],
              "modified": ["app/controller/application.rb"],
              "removed": []
            },
            {
              "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
              "message": "fixed readme",
              "title": "fixed readme",
              "timestamp": "2012-01-03T23:36:29+02:00",
              "url": "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
              "author": {
                "name": "GitLab dev user",
                "email": "gitlabdev@dv6700.(none)"
              },
              "added": ["CHANGELOG"],
              "modified": ["app/controller/application.rb"],
              "removed": []
            }
          ],
          "total_commits_count": 4
        }`

	// 测试反序列化json为struct
	event := GitlabEvent{}
	err := json.Unmarshal([]byte(pushJSON), &event)
	if err != nil {
		t.Errorf("Gitlab event json解析错误: %s", err)
	}

	// 测试发送消息到weixin, dingding
	wxurl := beego.AppConfig.DefaultString("wxurl", "")
	if wxurl == "" {
		t.Log("wxurl为空, 消息未能发送到企微机器人。")
	} else {
		resp := sendGitlabEvent(3, event, "Push Hook", "", wxurl)
		if resp != "消息发送完成!" {
			t.Errorf("发送gitlab push event消息到企业微信失败: %s", resp)
		}
	}

	ddurl := beego.AppConfig.DefaultString("ddurl", "")
	if ddurl == "" {
		t.Log("ddurl为空, 消息未能发送到钉钉机器人。")
	} else {
		resp := sendGitlabEvent(2, event, "Push Hook", "", ddurl)
		if resp != "消息发送完成!" {
			t.Errorf("发送gitlab push event消息到钉钉失败: %s", resp)
		}
	}
	fsUrl := beego.AppConfig.DefaultString("fsurl","")
	if fsUrl == "" {
		t.Log("fsUrl为空, 消息未能发送到飞书机器人。")
	} else {
		resp := sendGitlabEvent(4, event, "Push Hook", "", fsUrl)
		if resp != "消息发送完成!" {
			t.Errorf("发送gitlab push event消息到飞书失败: %s", resp)
		}
	}
}
