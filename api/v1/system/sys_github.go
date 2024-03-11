package system

import (
	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"

	// "encoding/json"
	"io"
	"net/http"
	"server-fiber/global"
	"server-fiber/model/common/request"
	"server-fiber/model/common/response"
	"server-fiber/model/system"
	"server-fiber/service"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type SystemGithubApi struct{}

var githubService = service.ServiceGroupApp.SystemServiceGroup.GithubService

type GithubCommit struct {
	URL         string `json:"url"`
	Sha         string `json:"sha"`
	NodeID      string `json:"node_id"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Commit      struct {
		URL    string `json:"url"`
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			URL string `json:"url"`
			Sha string `json:"sha"`
		} `json:"tree"`
		CommentCount int `json:"comment_count"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		URL string `json:"url"`
		Sha string `json:"sha"`
	} `json:"parents"`
}

func (g *SystemGithubApi) GetGithubList(c *fiber.Ctx) error {
	var searchInfo request.PageInfo
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")
	searchInfo.Page, _ = strconv.Atoi(page)
	searchInfo.PageSize, _ = strconv.Atoi(pageSize)

	if list, err := githubService.GetGithubList(searchInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		return response.FailWithMessage("获取失败", c)
	} else {
		return response.OkWithDetailed(response.PageResult{
			List:     list,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, "获取成功", c)
	}
}

func (g *SystemGithubApi) CreateGithub(c *fiber.Ctx) error {
	data := make([]system.SysGithub, 1)

	page := "1"
	per_page := "5"
	resp, err := http.Get("https://api.github.com/repos/JiangHaoCode/server-web/commits?page=" + page + "&per_page=" + per_page)
	defer func() {
		_ = resp.Body.Close()

	}()
	if err != nil {
		global.LOG.Error("请求Commit错误", zap.Error(err))
		return response.FailWithMessage("请求Commit错误", c)
		// log.Println(resp.Body)

	}
	body, _ := io.ReadAll(resp.Body)
	// respData := new([]GithubCommit)
	var respData []GithubCommit
	json.Unmarshal(body, &respData)
	time.LoadLocation("Asia/Shanghai")
	for _, val := range respData {
		var temp system.SysGithub
		temp.Author = val.Commit.Author.Name
		temp.CommitTime = val.Commit.Author.Date.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
		temp.Message = val.Commit.Message
		data = append(data, temp)
	}
	if total, err := githubService.CreateApi(data); err != nil {
		global.LOG.Error("创建commit有错误!", zap.Error(err))
		return response.FailWithMessage("创建commit有错误!", c)
	} else {
		return response.OkWithData(fiber.Map{"total": total}, c)
	}
}
