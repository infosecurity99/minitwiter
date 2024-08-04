package handler

import (
	"context"
	"net/http"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRetweet godoc
// @Router       /retweet [POST]
// @Summary      Creates a new retweet
// @Description  Create a new retweet for a tweet
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        retweet body models.CreateRetweet true "retweet"
// @Success      201  {object}  models.Retweet
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateRetweet(c *gin.Context) {
	var createRetweet models.CreateRetweet

	if err := c.ShouldBindJSON(&createRetweet); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	id, err := h.services.Retweets().Create(ctx, createRetweet)
	if err != nil {
		handleResponse(c, h.log, "error while creating retweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, id)
}

// DeleteRetweet godoc
// @Router       /retweet/{id} [DELETE]
// @Summary      Delete retweet
// @Description  Delete a retweet by ID
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        id path string true "retweet_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteRetweet(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Retweets().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, h.log, "error while deleting retweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "retweet successfully deleted")
}
