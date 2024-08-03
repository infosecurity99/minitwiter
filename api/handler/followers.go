package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/models"
	"strconv"
	"time"
)

// CreateFollower godoc
// @Router       /follower [POST]
// @Summary      Creates a new follower relationship
// @Description  Create a new follower relationship between two users
// @Tags         follower
// @Accept       json
// @Produce      json
// @Param        follower body models.CreateFollower true "follower"
// @Success      201  {object}  models.Follower
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateFollower(c *gin.Context) {
	var createFollower models.CreateFollower

	if err := c.ShouldBindJSON(&createFollower); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Followers().Create(ctx, createFollower)
	if err != nil {
		handleResponse(c, h.log, "error while creating follower relationship", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetFollower godoc
// @Router       /follower/{id} [GET]
// @Summary      Get follower by ID
// @Description  Get a follower relationship by its ID
// @Tags         follower
// @Accept       json
// @Produce      json
// @Param        id path string true "follower_id"
// @Success      200  {object}  models.Follower
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetFollower(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Followers().Get(ctx, uid)
	if err != nil {
		handleResponse(c, h.log, "error while getting follower by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// GetFollowerList godoc
// @Router       /followers [GET]
// @Summary      Get list of followers
// @Description  Get a list of follower relationships
// @Tags         follower
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Success      200  {object}  models.FollowersResponse
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetFollowerList(c *gin.Context) {
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
	resp, err := h.services.Followers().GetList(ctx, models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting list of followers", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// DeleteFollower godoc
// @Router       /follower/{id} [DELETE]
// @Summary      Delete follower relationship
// @Description  Delete a follower relationship by ID
// @Tags         follower
// @Accept       json
// @Produce      json
// @Param        id path string true "follower_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteFollower(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Followers().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, h.log, "error while deleting follower relationship", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "follower relationship successfully deleted")
}
