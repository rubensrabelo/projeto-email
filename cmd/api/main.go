package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := database.NewDb()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Patch("/campaigns/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaigns/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))

	log.Println("Servidor rodando em http://localhost:3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}