package script

import "article/pkg/elasticsearch"

func CreateIndex(es *elasticsearch.Elasticsearch) error {
	return es.CreateIndex()
}
