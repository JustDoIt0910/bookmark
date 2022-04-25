package cron

import "github.com/robfig/cron/v3"

var Cron *cron.Cron

func Init() error {
	Cron = cron.New()
	_, err := Cron.AddFunc("0 0 * * *", ClearRecycleBin)
	//_, err := Cron.AddFunc("@every 1m", ClearRecycleBin)
	if err != nil {
		return err
	}
	Cron.Start()
	return nil
}
