package api

import "foundations/bootstrap"
import "github.com/garyburd/redigo/redis"

type Registry struct {
	Name string
}

func (r Registry) GetKey(key string) string {
	return r.Name + ":" + key
}
func (r Registry) Get(key string) (string, error) {
	k := r.GetKey(key)
	return redis.String(bootstrap.RedigoDo("GET", k))
}
func (r Registry) Del(key string) (interface{}, error) {
	k := r.GetKey(key)
	return bootstrap.RedigoDo("DEL", k)
}
func (r Registry) Set(key string, value string) (interface{}, error) {
	k := r.GetKey(key)
	return bootstrap.RedigoDo("SET", k, value)
}
