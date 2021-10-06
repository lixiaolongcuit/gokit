package v6

import (
	"context"
	"testing"

	"github.com/lixiaolongcuit/gokit/pkg/elasticx"
)

func TestNewElastic(t *testing.T) {
	cfg := &elasticx.Config{
		Urls: []string{
			"http://172.30.3.116:9200",
		},
		Username: "admin",
		Password: "skg2102stg",
	}
	es, err := NewElastic(cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	exist, err := es.IndexExists("dbl_white_list_v1").Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("index exists: %v", exist)
}
