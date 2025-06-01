package v1

import (
	"go-clean-template/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PegawaiHandler struct {
	uc usecase.PegawaiUsecase
}

func NewPegawaiHandler(uc usecase.PegawaiUsecase) *PegawaiHandler {
	return &PegawaiHandler{uc: uc}
}

func (h *PegawaiHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.uc.GetAllPegawai(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
