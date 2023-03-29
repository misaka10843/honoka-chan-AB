package main

import (
	"honoka-chan/config"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	"honoka-chan/sifcap"
	_ "honoka-chan/tools"
	"honoka-chan/xclog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	xclog.Init(config.Conf.Log.LogDir, "", config.Conf.Log.LogLevel, config.Conf.Log.LogSave)
}

func main() {
	if config.Conf.SifCap.Enabled {
		sifcap.Start()
	} else {
		// router
		r := gin.Default()
		r.Use(middleware.KlabHeader)

		r.Any("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello, world!")
		})

		r.GET("/webview.php/announce/index", handler.AnnounceIndexHandler)

		r.POST("/main.php/login/authkey", handler.AuthKeyHandler)
		r.POST("/main.php/login/login", handler.LoginHandler)
		r.POST("/main.php/user/userInfo", handler.UserInfoHandler)
		r.POST("/main.php/gdpr/get", handler.GdprHandler)
		r.POST("/main.php/personalnotice/get", handler.PersonalNoticeHandler)
		r.POST("/main.php/tos/tosCheck", handler.TosCheckHandler)
		r.POST("/main.php/download/batch", handler.DownloadBatchHandler)
		r.POST("/main.php/download/event", handler.DownloadEventHandler)
		r.POST("/main.php/lbonus/execute", handler.LBonusExecuteHandler)
		r.POST("/main.php/api", handler.ApiHandler)
		r.POST("/main.php/announce/checkState", handler.AnnounceCheckStateHandler)
		r.POST("/main.php/scenario/startup", handler.ScenarioStartupHandler)
		r.POST("/main.php/user/setNotificationToken", handler.SetNotificationTokenHandler)
		r.POST("/main.php/user/changeNavi", handler.SetNotificationTokenHandler)
		r.POST("/main.php/event/eventList", handler.EventListHandler)
		r.POST("/main.php/payment/productList", handler.ProductListHandler)
		r.POST("/main.php/live/partyList", handler.PartyListHandler)
		r.POST("/main.php/live/play", handler.PlayLiveHandler)
		r.POST("/main.php/live/gameover", handler.GameOverHandler)
		r.POST("/main.php/unit/setDisplayRank", handler.SetDisplayRankHandler)
		r.POST("/main.php/unit/favorite", handler.SetDisplayRankHandler)

		r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
