package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yogie/go-clean-api/cache"
	"github.com/yogie/go-clean-api/controller"
	"github.com/yogie/go-clean-api/repository"
	"github.com/yogie/go-clean-api/router"
	"github.com/yogie/go-clean-api/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)

	postRepository repository.PostRepository = repository.NewPostgreRepository() // Change YOUR DB Stack here
	postService    service.PostService       = service.NewPostService(postRepository)
	postChache     cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postChache)
	httpRouter     router.Router             = router.NewChiRouter() // Change Your Routing Framework Here
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome to golang")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	log.Println("Server listening on port", port)

	httpRouter.SERVE(port)

}
