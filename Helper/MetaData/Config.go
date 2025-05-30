package Config

import "os"

var RedisConnString string
var DbConnString string

var MainURL string = "localhost:9030"

func InitiateConfig() {
	currDbStringName := "DBConnProd"
	//currDbStringName := "DBConnTest"

	currRedisStringName := "RedisConnProd"
	//currRedisStringName := "RedisConnTest"

	dbConnString := os.Getenv(currDbStringName)
	DbConnString = dbConnString

	redisConnString := os.Getenv(currRedisStringName)
	RedisConnString = redisConnString
}
