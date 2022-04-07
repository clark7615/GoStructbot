Champ StructBot 
---
此專案為一個Go/Golang套件庫，讓使用者能方便的從其他資料型態轉換為Struct

系統需求為 [Go Programming Language](https://golang.org/dl/), version 1.15 and above.

安裝
---
```shell script
$ go get gitlab.com/champinfo/go/structbot
```
或者編輯您專案中的go.mod檔案
```
module your_project_name

go 1.15.1

require (
    gitlab.com/champinfo/go/champiris vx.x.x
)
```

```shell script
$ go build
```

如何使用
---
StructBot會依照您在struct中填入的tag自動地幫你分配解析的資料格式

例如<br>
 ``` ID int `json:"id"` ```bot將會用json解析器解析資料<br>
 ``` ID int `yaml:"id"` ```bot將會採用yaml解析器解析資料<br>
 ``` ID int `xml:"id"` ```bot將會採用xml解析器解析資料<br>
 當然你也可以在其中使用兩種以上的Tag例如``` ID   int    `json:"id" yaml:"id" xml:"id"` ``` <br>
在最新的版本中我們添加了mpa的解析

```go
import (
	"gitlab.com/champinfo/go/structbot"
)

const (
	dataJson = `
{"id":1,"data":"test"}
`
	dataYaml = `
id: 1
data: test
`
	dataXml = `
<root>
   <data>test</data>
   <id>1</id>
</root>
`
)

type ExpJson struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

type ExpYaml struct {
	Id   int    `yaml:"id"`
	Data string `yaml:"data"`
}

type ExpXml struct {
	Id   int    `xml:"id"`
	Data string `xml:"data"`
}

func main() {
	json := &ExpJson{}
	if err := structbot.MakeStruct(dataJson, json); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("inputData: %sstruct: %+v \n", dataJson, json)
	//inputData: 
	// {"id":1,"data":"test"}
	// struct: &{Id:1 Data:test} 
	yaml := &ExpYaml{}
	if err := structbot.MakeStruct(dataYaml, yaml); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("inputData: %sstruct: %+v \n", dataYaml, yaml)
	//inputData: 
	//id: 1 
	//data: test 
	//struct: &{Id:1 Data:test} 
	xml := &ExpXml{}
	if err := structbot.MakeStruct(dataXml, xml); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("inputData: %sstruct: %+v \n", dataXml, xml)
	//inputData:
	//<root> 
	//  <data>test</data> 
	//  <id>1</id> 
	//</root> 
	//struct: &{Id:1 Data:test}
	data := map[string]interface{}{
		"id":   1,
		"data": "測試",
	}
	yaml := &ExpYaml{}
	if err := structbot.MakeStruct(data, yaml); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("inputData: %+v struct: %+v \n", data, yaml) 
    //inputData: map[data:測試 id:1] struct: &{Id:1 Data:測試} 
}
```

有限制的功能
---
為了開發上的方便，我們新增了一個讀取config.yaml的功能<br>
須在專案目錄下新增config資料夾並在資料夾中建立一個config.yaml檔案<br>
此檔案必須為yaml格式
```go
type ExpYaml struct {
	Id   int    `yaml:"id"`
	Data string `yaml:"data"`
}
func main() {
    config := &ExpYaml{}
    structbot.FileMakeStruct(config)
    //inputData: 
    //id: 1 
    //data: test 
    //struct: &{Id:1 Data:test} 
}
```

###若有相關疑問請以 [Email](clark@championtek.com.tw) 詢問我們或是到我們的 [官方網站](https://www.championtek.com.tw/) 留言給我們
