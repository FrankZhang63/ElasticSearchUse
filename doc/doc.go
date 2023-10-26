package doc

import (
	"ElasticSearch/global"
	"ElasticSearch/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func CreateDoc() {
	user := models.UserModel{
		ID:       10,
		UserName: "张三",
		Age:      24,
		NickName: "夜空中最亮的张三",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 第一个index()表明要添加数据，第二个index()表明索引值
	indexResponse, err := global.ESclient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", indexResponse)
}

// CreateDocBatch 根据id批量创建
func CreateDocBatch() {
	list := []models.UserModel{
		{
			ID:       11,
			UserName: "李四",
			Age:      25,
			NickName: "夜空中最亮的李四",
			Title:    "夜空中最亮的李四",
			CreateAt: time.Now().Format("2006-01-02 15:04:05")},
		{
			ID:       12,
			UserName: "王五",
			Age:      26,
			NickName: "夜空中最亮的王五",
			Title:    "夜空中最亮的王五",
			CreateAt: time.Now().Format("2006-01-02 15:04:05")},
		{
			ID:       13,
			UserName: "赵六",
			Age:      27,
			NickName: "夜空中最亮的赵六",
			Title:    "夜空中最亮的赵六",
			CreateAt: time.Now().Format("2006-01-02 15:04:05")},
	}
	// 批量操作的时候要先建立一个桶 bulk,在桶内进行相关的操作，Refresh true表示每一次操作都刷新索引。false可能会出现删除掉了，但是查询还在的情况
	bulk := global.ESclient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, model := range list {
		req := elastic.NewBulkCreateRequest().Doc(model)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(res.Succeeded()))
}

// DeleteDoc 根据id删除
func DeleteDoc() {
	deleteResponse, err := global.ESclient.Delete().
		Index(models.UserModel{}.Index()).Id("yxQyaosBhGoF22ZQ8ZU9").Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}

// DeleteDocBatch 根据id批量删除
func DeleteDocBatch() {
	idList := []string{
		"2xw2aosBHbPjJicxcfIC", "3BygaosBHbPjJicxIvKL",
	}
	bulk := global.ESclient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, s := range idList {
		// 删除传入id，创建传入doc
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(res.Succeeded())) //查看实际删除成功的个数
}
func UpdateDoc() {
	res, err := global.ESclient.Update().Index(models.UserModel{}.Index()).Id("bVXuaosB30HJR3AcWxgc").Doc(map[string]any{
		"user_name": "FrankZhang",
	}).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", res)
}
func SelectAllDoc() {
	// 全查   From是从第几条记录开始，Size是输出几条记录
	res, err := global.ESclient.Search(models.UserModel{}.Index()).From(0).Size(3).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// 统计搜索结果 1.res.Hits.TotalHits.Value  2.res.TotalHits()
	count := res.TotalHits()
	fmt.Println(count)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
func SelectTermDoc() {
	// Match主要用来匹配type = text类型的字段。模糊匹配，match匹配的时候会将参数的分词和库里的分词逐一进行比对，如果分词一样就匹配成功。
	// Term主要用来精准匹配 type = keyword 类型的字段。精准匹配
	//query := elastic.NewMatchQuery("nick_name", "夜空中最亮的李四")

	//"title": {
	//	"type": "text",
	//		"fields": {
	//		"keyword": {
	//			"type": "keyword",
	//				"ignore_above": 256
	//		}
	//	}
	//},
	// title是text类型，只能用模糊查询，但是如需精确匹配，可以title.keyword 的形式来精准搜索
	query := elastic.NewTermQuery("title.keyword", "夜空中最亮的李四")
	res, err := global.ESclient.Search(models.UserModel{}.Index()).Query(query).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// 统计搜索结果 1.res.Hits.TotalHits.Value  2.res.TotalHits()
	count := res.TotalHits()
	fmt.Println(count)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
