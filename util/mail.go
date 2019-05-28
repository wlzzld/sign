package util

import (
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/gomail.v2"
	"wlzzld.cn/sign/baidu"
)

func SendMail(mailTo []string, subject string, body string) error {
	user := GetValStr("mail_from_user")
	pass := GetValStr("mail_from_pass")
	host := GetValStr("mail_from_host")
	port, _ := GetValInt("mail_from_port")

	m := gomail.NewMessage()
	m.SetHeader("From", user)       //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(host, port, user, pass)

	err := d.DialAndSend(m)
	return err

}

type data struct {
	Tiebas     []baidu.Tieba
	Count      int
	TotalScore int64
}

func BuildBody(tiebas []baidu.Tieba) string {
	count := len(tiebas)
	var totalScore int64 = 0
	for _, t := range tiebas {
		totalScore += t.AddScore
	}

	tmpl, _ := template.ParseFiles("template/sign_mail.tpl")
	file, _ := os.Create("template/sign_mail.html")
	_data := data{
		Tiebas:     tiebas,
		Count:      count,
		TotalScore: totalScore,
	}
	tmpl.Execute(file, _data)
	byt, _ := ioutil.ReadFile("template/sign_mail.html")
	defer file.Close()

	body := string(byt)
	return body
}
