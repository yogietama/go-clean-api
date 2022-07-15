package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yogie/go-clean-api/controller"
	"github.com/yogie/go-clean-api/repository"
	"github.com/yogie/go-clean-api/router"
	"github.com/yogie/go-clean-api/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	// carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)

	postRepository repository.PostRepository = repository.NewFirestoreRepository() // Change YOUR DB Stack here
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter() // Change Your Routing Framework Here
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome to golang")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	log.Println("Server listening on port", port)

	httpRouter.SERVE(port)

}
