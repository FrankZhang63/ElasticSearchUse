package indexs

import (
	"ElasticSearch/global"
	"context"
	"fmt"
)

func CreateIndex(esindex string, esMapping string) {
	if ExistIndex(esindex) {
		fmt.Println(esindex, "索引存在")
		//索引存在，先删除，再创建
		DeleteIndex(esindex)
	}
	createIndex, err := global.ESclient.
		CreateIndex(esindex).
		BodyString(esMapping).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createIndex, "索引创建成功")
}

// ExistIndex 判断索引是否存在
func ExistIndex(esindex string) bool {
	exist, err := global.ESclient.IndexExists(esindex).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return exist
}

// DeleteIndex 删除索引
func DeleteIndex(esindex string) {
	_, err := global.ESclient.DeleteIndex(esindex).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(esindex, "索引删除成功")
}

//  global.ESclient.CreateIndex().BodyString().Do(context.Background())
//  global.ESclient.IndexExists().Do(context.Background())
//  global.ESclient.DeleteIndex().Do(context.Background())
