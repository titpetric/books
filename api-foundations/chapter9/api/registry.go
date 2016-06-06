package api

import "foundations/bootstrap"
import "github.com/garyburd/redigo/redis"

import "expvar"

var (
	countGet       *expvar.Int
	countSet       *expvar.Int
	countDel       *expvar.Int
	countGetAll    *expvar.Int
	countGetAllGet *expvar.Int
)

func init() {
	countGet = expvar.NewInt("registry.get")
	countSet = expvar.NewInt("registry.set")
	countDel = expvar.NewInt("registry.del")
	countGetAll = expvar.NewInt("registry.getAll")
	countGetAllGet = expvar.NewInt("registry.getAll.get")
}

type Registry struct {
	Name string
}

func (r Registry) GetKey(key string) string {
	return r.Name + ":" + key
}
func (r Registry) Get(key string) (string, error) {
	k := r.GetKey(key)
	countGet.Add(1)
	return redis.String(bootstrap.RedigoDo("GET", k))
}
func (r Registry) Del(key string) (interface{}, error) {
	k := r.GetKey(key)
	countDel.Add(1)
	return bootstrap.RedigoDo("DEL", k)
}
func (r Registry) Set(key string, value string) (interface{}, error) {
	k := r.GetKey(key)
	countSet.Add(1)
	return bootstrap.RedigoDo("SET", k, value)
}
func (r Registry) GetAll() (map[string]string, error) {
	k := r.GetKey("*")
	countGetAll.Add(1)
	keys, err := redis.Strings(bootstrap.RedigoDo("KEYS", k))
	allkeys := map[string]string{}
	if len(keys) == 0 || err != nil {
		return allkeys, nil
	}
	countGetAllGet.Add(int64(len(keys)))
	for _, value := range keys {
		value_redis, err := redis.String(bootstrap.RedigoDo("GET", value))
		if err == nil {
			allkeys[value[len(r.Name+":"):]] = value_redis
		}
	}
	return allkeys, nil
}
