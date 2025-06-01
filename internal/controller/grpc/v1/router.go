package v1

import (
	"go-clean-template/internal/controller/http/v1"
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"

	"github.com/go-playground/validator/v10"
	pbgrpc "google.golang.org/grpc"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		v1.RegisterTranslationServer(app, r)
	}
}

func NewRouter(r *gin.Engine, pegawaiUC usecase.PegawaiUsecase) {
	h := handler.NewPegawaiHandler(pegawaiUC)

	v1 := r.Group("/v1")
	{
		v1.GET("/pegawai", h.GetAll)
	}
}
