package main

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/robfig/cron"

	"wlzzld.cn/sign/baidu"
	"wlzzld.cn/sign/util"
)

//主机名
var hostname, _ = os.Hostname()
var configfile = "app.properties"

type Tiebas []baidu.Tieba

//Len()
func (s Tiebas) Len() int {
	return len(s)
}

//Less():贴吧当前经验将由高到低排序
func (s Tiebas) Less(i, j int) bool {
	return s[i].CurScore > s[j].CurScore
}

//Swap()
func (s Tiebas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	util.InitConfig(configfile)
	log.Println("定时器初始化开始！")
	c := cron.New()
	taskCron1 := "3 0 0 * * ? "
	c.AddFunc(taskCron1, func() {
		baiduSign()
	})
	log.Println("第一次签到定时初始化完成！")
	taskCron2 := "0 0 9,11,13,15,17,19 * * ? "
	c.AddFunc(taskCron2, func() {
		baiduSign()
	})

	log.Println("补签定时初始化完成！")
	c.Start()
	log.Println("定时器初始化完成！")

	for {
		runtime.Gosched()
	}

}
func baiduSign() {
	tiebas := Tiebas{}
	bduss := util.GetValStr("bduss")

	baidu.GetDoc("http://c.tieba.baidu.com/c/f/forum/like")
	userName := baidu.GetUserName(bduss)
	userId := baidu.GetUserID(userName)

	tiebaList := baidu.GetTiebaList(userId, bduss)

	a := time.Now().Unix()

	pool := new(util.GoroutinePool)
	poolSize, _ := util.GetValInt("pool_size")
	pool.Init(poolSize, len(tiebaList))

	for i := 0; i < len(tiebaList); i++ {
		tieba := tiebaList[i]
		pool.AddTask(func() error {
			_tieba := baidu.Sign(tieba, bduss)
			__tieba := baidu.AfterSign(_tieba, bduss)
			tiebas = append(tiebas, __tieba)
			return nil
		})
	}

	isFinish := false

	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
		}(&isFinish)
	})

	pool.Start()

	for !isFinish {
		time.Sleep(time.Millisecond * 100)
	}

	pool.Stop()
	b := time.Now().Unix()
	log.Printf("签到操作已结束，共个%d贴吧，耗时:%d秒\n", len(tiebas), (b - a))

	sort.Sort(tiebas)

	// 定义收件人
	mailTo := strings.Split(util.GetValStr("mail_to"), ",")
	// 邮件主题为"Hello"
	subject := "GO自动签到结果:" + time.Now().Format("2006-01-02 15:04:05") + hostname
	// 邮件正文
	body := util.BuildBody(tiebas)
	util.SendMail(mailTo, subject, body)
	c := time.Now().Unix()
	log.Printf("邮件发送已结束，耗时:%d秒\n", (c - b))
}
