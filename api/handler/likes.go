package handler

import (
	"context"
	"net/http"
	"strconv"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateLike godoc
// @Router       /like [POST]
// @Summary      Creates a new like
// @Description  Create a new like for a tweet
// @Tags         like
// @Accept       json
// @Produce      json
// @Param        like body models.CreateLike true "like"
// @Success      201  {object}  models.Like
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateLike(c *gin.Context) {
	var createLike models.CreateLike

	if err := c.ShouldBindJSON(&createLike); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Likes().Create(ctx, models.CreateTweet{})
	if err != nil {
		handleResponse(c, h.log, "error while creating like", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetLike godoc
// @Router       /like/{id} [GET]
// @Summary      Get like by ID
// @Description  Get a like by its ID
// @Tags         like
// @Accept       json
// @Produce      json
// @Param        id path string true "like_id"
// @Success      200  {object}  models.Like
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetLike(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Likes().Get(ctx, uid)
	if err != nil {
		handleResponse(c, h.log, "error while getting like by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// GetLikeList godoc
// @Router       /likes [GET]
// @Summary      Get list of likes
// @Description  Get a list of likes
// @Tags         like
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Success      200  {object}  models.LikesResponse
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetLikeList(c *gin.Context) {
	var (
		page, limit int
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Likes().GetList(ctx, models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting list of likes", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// DeleteLike godoc
// @Router       /like/{id} [DELETE]
// @Summary      Delete like
// @Description  Delete a like by ID
// @Tags         like
// @Accept       json
// @Produce      json
// @Param        id path string true "like_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteLike(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Likes().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, h.log, "error while deleting like", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "like successfully deleted")
}
