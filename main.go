package main

import (
	"ElasticSearch/core"
	"ElasticSearch/doc"
	"ElasticSearch/indexs"
	"ElasticSearch/models"
)

func main() {
	var ESServerURL string = "http://ip:port"
	//初始化es连接
	core.EsConnect(ESServerURL)
	var esIndex = "user_index"
	indexs.CreateIndex(esIndex, models.UserModel{}.Mapping())
	doc.CreateDoc()
	doc.DeleteDoc()
	doc.DeleteDocBatch()
	doc.CreateDocBatch()
	doc.SelectAllDoc()
	doc.SelectTermDoc()
	doc.UpdateDoc()

	////boolquery 可用于组合查询
	////Must想当于且，Should相当于或，MustNot相当于非......
	//boolquery := elastic.NewBoolQuery()
	//boolquery.Must(elastic.NewMatchQuery("name", "java")) //查询name为java的
	//searchByMatch, err := ESClient.Search(esIndex).Type(esType).Query(boolquery).Do(context.Background())
	//for _, item := range searchByMatch.Each(reflect.TypeOf(resultType)) {
	//	language := item.(Language)
	//	fmt.Printf("search by match: %#v \n", language)
	//}

	////匹配查询
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("name", "py") //查询name包含py的
	//searchByPhrase, err := ESClient.Search(esIndex).Type(esType).Query(matchPhraseQuery).Do(context.Background())
	//for _, item := range searchByPhrase.Each(reflect.TypeOf(resultType)) {
	//	language := item.(Language)
	//	fmt.Printf("search by phrase: %#v \n", language)
	//}

	////条件查询
	//boolquery2 := elastic.NewBoolQuery()
	//boolquery2.Filter(elastic.NewRangeQuery("build_time").Gt(2000)) //查询build_time大于2000
	//searchByfilter, err := ESClient.Search(esIndex).Type(esType).Query(boolquery2).Do(context.Background())
	//for _, item := range searchByfilter.Each(reflect.TypeOf(resultType)) {
	//	language := item.(Language)
	//	fmt.Printf("search by filter: %#v \n", language)
	//}
}
