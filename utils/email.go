package utils

import (
	"bookmark/config"
	"bookmark/logger"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sync"
	"time"
)

var EmailSender *emailSender

type emailSender struct {
	wg sync.WaitGroup
	ch chan *email.Email
	emailPool *email.Pool
}

func InitEmailSender(poolSize int) error {
	var sender emailSender
	sender.wg.Add(10)
	sender.ch = make(chan *email.Email, 5)
	addr := config.GlobalConfig.GetString("email.host") + ":25"
	username := config.GlobalConfig.GetString("email.username")
	password := config.GlobalConfig.GetString("email.password")
	host := config.GlobalConfig.GetString("email.host")

	var err error
	sender.emailPool, err = email.NewPool(addr,
										poolSize,
										smtp.PlainAuth("", username, password, host))
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer sender.wg.Done()
			for em := range sender.ch {
				e := sender.emailPool.Send(em, 10 * time.Second)
				if e != nil {
					logger.SugarLogger.Infof("Email %v send fail. Error: %s", em, e.Error())
				}
			}
		}()
	}
	EmailSender = &sender
	return nil
}

func (sender *emailSender) Send(from string, to string, subject string, text string)  {
	e := email.NewEmail()
	// 设置发送邮件的基本信息
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	sender.ch <- e
}

func (sender *emailSender) Close()  {
	close(sender.ch)
	sender.wg.Wait()
}