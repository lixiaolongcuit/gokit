package v7

import (
	"context"
	"testing"

	"github.com/lixiaolongcuit/gokit/pkg/elasticx"
)

func TestNewElastic(t *testing.T) {
	cfg := &elasticx.Config{
		Username:     "elastic",
		Password:     "123456",
		DisableSniff: true,
	}
	es, err := NewElastic(cfg, nil)
	if err != nil {
		t.Fatal("connect error: ", err)
	}
	exist, err := es.IndexExists("kibana_sample_data_flights").Do(context.Background())
	if err != nil {
		t.Fatal("index exists error:", err)
	}
	t.Logf("index exists: %v", exist)
}
