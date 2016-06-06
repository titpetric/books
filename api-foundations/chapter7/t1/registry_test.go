package api

import "testing"
import "foundations/bootstrap"

func TestRegistryGet(t *testing.T) {
	redisPool := bootstrap.RedigoPool()
	defer redisPool.Close()

	reg := Registry{Name: "test"}
	reg.Del("name")
	val, err := reg.Get("name")
	if err == nil || val != "" {
		t.Errorf("Unexpected result when getting name: %s/%s\n", val, err)
	}
}
