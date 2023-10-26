package core

import (
	"ElasticSearch/global"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func EsConnect(ESServerURL string) {
	//初始化es连接
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(ESServerURL),
	)
	if err != nil {
		log.Errorf("Failed to build elasticsearch connection: %s %s", ESServerURL, err.Error())
		return
	}
	global.ESclient = client
}
