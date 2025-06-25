//  Question 5: (  Storage Engine)
// You want to write a program that can store data into:
// Local file
// PostgreSQL DB
// Redis cache
// Each storage type must implement:
// Save(data string)
// Load(key string) string

package interface_example3

import (
	"fmt"
)

type StorageEngine interface {
	Save(data string)
	Load(key string) string
}

type Localfile struct {
	path string
}

func (l Localfile) Save(data string) {
	// getting data from somewhere and saving to local file
	fmt.Println("saving data locally on : ", l.path, "\ndata: ", data)
}
func (l Localfile) Load(key string) string {
	// data gets load from this funciton from localfile to UI it will return files
	// here we take data from path l.path then print here
	return fmt.Sprintln("this data is drawn from local storage path: ", l.path, "\ndata is : ...<dynamic>...\nkey is: ", key)
}

// postgresDB
type postgresDB struct {
	ip      string
	db_name string
	port    int
}

func (p postgresDB) Save(data string) {
	// getting data from somewhere and saving to postgres db
	// get data
	fmt.Printf("saving data:%s  on ip: %s running on port: %d on db:%s\n", data, p.ip, p.port, p.db_name)
}
func (p postgresDB) Load(key string) string {
	// here we are getting data from database and putting on ui or cli
	return fmt.Sprintf("getting data from : %s and here is data: <dynamic> and key is %s", p.db_name, key)
}

type RedisCache struct {
	RedisIP string
}

func (r RedisCache) Save(data string) {
	fmt.Printf("\nhere we aresaving data to redis: %s cache  and data is %s ", r.RedisIP, data)
}
func (r RedisCache) Load(key string) string {
	return fmt.Sprintf("\nhere we are getting data from redis: %s and data is <dynamic> and key is :%s ", r.RedisIP, key)
}

func binder(s StorageEngine, data string, key string) {
	s.Save(data)
	l := s.Load(key)
	fmt.Println(l)
}

func main() {

	var s StorageEngine
	s = Localfile{
		path: "/var/log/local",
	}
	binder(s, "this is data for local storage", "AmritKeylocal")
	s = postgresDB{
		ip:      "10.21.32.100",
		db_name: "main-db",
		port:    3390,
	}
	binder(s, "this is data for postgres", "AmritKeypostgres")
	s = RedisCache{
		RedisIP: "10.43.100.24",
	}
	binder(s, "this is data for redis", "AmritKeyForRedis")

}
