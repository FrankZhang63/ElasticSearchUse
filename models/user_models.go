package models

type UserModel struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	CreateAt string `json:"create_at"`
	Title    string `json:"title"`
	Age      int    `json:"age"`
}

func (UserModel) Index() string {
	return "user_index"
}
func (UserModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "nick_name": {"type": "text"},    //模糊匹配
	  "user_name": {"type": "keyword"}, //精准匹配     
	  "id": {"type": "integer"},
	  "age": {"type": "integer"},
	  "title": {
	    "type": "text",
		"fields": {
		  "keyword": {
			"type": "keyword",
			"ignore_above": 256
		  }
		}
	  },
	  "create_at": {
        "type": "date",
		"null_value": "null",
		"format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}`
}
