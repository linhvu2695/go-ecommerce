package initialization

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()

	InitLogger()
	InitMySql()
	InitRedis()
	InitSmtp()
	InitKafka()
	InitPrometheus()

	r := NewRouter()

	return r
}
