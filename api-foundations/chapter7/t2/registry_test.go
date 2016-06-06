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

func TestRegistrySet(t *testing.T) {
	redisPool := bootstrap.RedigoPool()
	defer redisPool.Close()

	reg := Registry{Name: "test"}
	status, err := reg.Set("name", "Tit Petric")
	if status != "OK" || err != nil {
		t.Errorf("Error when using SET: %s", err)
	}
	val, err := reg.Get("name")
	if err != nil || val != "Tit Petric" {
		t.Errorf("Got error when getting name: %s/%s\n", val, err)
	}
}
