package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
	"test/api/models"
)

// CreateRetweet godoc
// @Router       /retweet [POST]
// @Summary      Retweet a tweet
// @Description  Allows a user to retweet an existing tweet
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        retweet body models.RetweetRequest true "Retweet Request"
// @Success      201  {object}  models.Retweet
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateRetweet(c *gin.Context) {
	retweetRequest := models.RetweetRequest{}

	if err := c.ShouldBindJSON(&retweetRequest); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Retweet().Create(ctx, retweetRequest)
	if err != nil {
		handleResponse(c, h.log, "error while creating retweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetRetweet godoc
// @Router       /retweet/{id} [GET]
// @Summary      Get retweet
// @Description  Get retweet by ID
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        id path string true "retweet_id"
// @Success      200  {object}  models.Retweet
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetRetweet(c *gin.Context) {
	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	retweet, err := h.services.Retweet().Get(ctx, id.String())
	if err != nil {
		handleResponse(c, h.log, "error while getting retweet by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, retweet)
}

// GetRetweetList godoc
// @Router       /retweets [GET]
// @Summary      Get retweet list
// @Description  Get list of retweets
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.RetweetResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetRetweetList(c *gin.Context) {
	var (
		page, limit int
		search      string
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

	search = c.Query("search")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Retweet().GetList(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting retweets", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "success!", http.StatusOK, resp)
}


// UpdateReTweet godoc
// @Router       /retweet/{id} [PUT]
// @Summary      Update retweet
// @Description  Update a retweet
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        id path string true "retweet_id"
// @Param        tweet body models.UpdateReTweet true "retweet"
// @Success      200  {object}  models.ReTweet
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTweet(c *gin.Context) {
	updateTweet := models.UpdateReTweet{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, h.log, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateTweet.ID = uid

	if err := c.ShouldBindJSON(&updateReTweet); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.ReTweet().Update(ctx, updateTweet)
	if err != nil {
		handleResponse(c, h.log, "error while updating tweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// DeleteRetweet godoc
// @Router       /retweet/{id} [DELETE]
// @Summary      Delete retweet
// @Description  Delete a retweet
// @Tags         retweet
// @Accept       json
// @Produce      json
// @Param        id path string true "retweet_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteRetweet(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.Retweet().Delete(ctx, models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, h.log, "error while deleting retweet by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "data successfully deleted")
}
