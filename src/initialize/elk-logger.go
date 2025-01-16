package initialize

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func InitElasticsearch() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	esClient = es
}


