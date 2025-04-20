package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"master-finanacial-planner/internal/constant"
	"master-finanacial-planner/internal/entity"
	"master-finanacial-planner/internal/handler"
	"master-finanacial-planner/internal/helper"
	"master-finanacial-planner/internal/repo"
	"master-finanacial-planner/internal/usecase/finance"
	"master-finanacial-planner/internal/usecase/user"
	"net/http"
)

func main() {

	// Initialize the database connection
	db, err := repo.InitializeDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close() // Ensure the DB connection closes when

	// setting up the internals
	dataSourceRepo := repo.NewResource(db)
	financeUsecase := finance.NewFinanceUsecase(dataSourceRepo)
	userUsecase := user.NewUserUsecase(dataSourceRepo)
	handler := handler.NewFinanceHandler(userUsecase, financeUsecase)

	// setting up the route
	router := chi.NewRouter()

	// health check
	router.Get("/service-health", func(w http.ResponseWriter, r *http.Request) {
		var response entity.ApiResponse
		response.Data = map[string]interface{}{"message": "Working fine"}
		response.Error = nil
		response.Success = true
		helper.WriteCustomResp(w, http.StatusOK, response)
	})

	// user-route
	router.Post("/sign-up", handler.SignUpHandler)
	router.Post("/sign-in", handler.SignInHandler)

	// finance-route

	// get routes
	router.Get("/get/asset-classes", handler.GetAssetClassHandler)
	// get effective returns on allocation type
	router.Get("/get/allocation/effective-assets", handler.GetEffectiveReturnAllocationTypeHandler)
	// investing surplus
	router.Get("/investing-surplus", handler.GetInvestingSurplusHandler)
	// investing
	router.Get("/net-worth", handler.GetNetWorthHandler)

	fmt.Printf("Master-financial Server Started at port %s\n", constant.ConfigPort)
	err = http.ListenAndServe(constant.ConfigPort, router)
	if err != nil {
		fmt.Println("Error while starting the master-financial server", err.Error())
	}

}
