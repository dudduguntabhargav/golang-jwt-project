package main
import(
	"github.com/gin-gonic/gin"
	routes "github.com/bhargavduddugunta/golang-jwt-project/routes"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main(){

			err := godotenv.Load(".env")
			if err!=nil{
				log.Fatal("error in loading .env file")
			}

			port := os.Getenv("PORT")
			if port==""{
				port = "8000"
			}
			
			router := gin.New()
			router.Use(gin.Logger())

			routes.UserRoutes(router)
      routes.AuthRoutes(router)

			router.GET("api-1/", func(c *gin.Context){
				c.JSON(200, gin.H{"success": "Access granted for api-1"})
			})

			router.GET("api-2/", func(c *gin.Context){
				c.JSON(200, gin.H{"success": "Access granted for api-1"})
			})

			router.Run(":" + port)
}