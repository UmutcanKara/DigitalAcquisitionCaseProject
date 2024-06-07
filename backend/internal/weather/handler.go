package weather

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	Service
}

func NewHandler(service Service) Handler { return &handler{service} }

func (h handler) FindWeather(c *gin.Context) {
	ctx := c.Request.Context()
	hometown := c.Query("hometown")
	startDate := c.Query("start_date")

	res, err := h.Service.findWeather(ctx, hometown, startDate)
	if err != nil {
		// If hometown weather does not exist, add it.
		err = h.Service.addWeather(ctx, hometown)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// If it doesn't appear on the 2nd try, don't try it further
		res, err = h.findWeather(ctx, hometown, startDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h handler) UpdateWeather(c *gin.Context) {
	ctx := c.Request.Context()
	hometown := c.Query("hometown")

	err := h.Service.updateWeather(ctx, hometown)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
