
package app

import 	"github.com/sandeepmendiratta/gin-go-boilerplate/controllers"


func initializeRoutes() {
	router.GET("/health", controllers.GetHealth)
	router.GET("/api", controllers.HandleGet)
	router.POST("/api", controllers.HandleVerification)
	router.OPTIONS("/api", controllers.HandleVerification)

}