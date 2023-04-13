package ES

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (s *Student) run() {
	fmt.Printf("%s在跑...", s.Name)
}
func InitEs() {
	cli, err := elastic.NewClient(elastic.SetURL("https://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es successed!")
	p1 := Student{
		Name:    "rion",
		Age:     22,
		Married: false,
	}
	put1, err := cli.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)

	}
	fmt.Printf("Indexed user %s to index %s ,type %s\n", put1.Id, put1.Index, put1.Type)
}
