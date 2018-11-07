package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/beanscc/rango/database/mysql/gormutil/examples"
	"github.com/jinzhu/gorm"
)

func GetAPPByID(ID int) (*examples.App, error) {
	var resp examples.App
	err := examples.Instance().Slave().Where("id=?", ID).First(&resp).Error
	return &resp, err
}

func main() {
	resp, err := GetAPPByID(1)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logrus.Errorf("query first failed. err=%v", err)
			return
		}
	}

	logrus.Infof("query resp=%v", resp)
}
