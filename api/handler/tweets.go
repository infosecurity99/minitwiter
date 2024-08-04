package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"test/api/models"
	"time"
)

// CreateTweet godoc
// @Router       /tweet [POST]
// @Summary      Creates a new tweet
// @Description  Create a new tweet
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        tweet body models.CreateTweet true "tweet"
// @Success      201  {object}  models.Tweet
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTweet(c *gin.Context) {
	createTweet := models.CreateTweet{}

	if err := c.ShouldBindJSON(&createTweet); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Tweets().Create(ctx, createTweet)
	if err != nil {
		handleResponse(c, h.log, "error while creating tweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetTweet godoc
// @Router       /tweet/{id} [GET]
// @Summary      Get tweet
// @Description  Get tweet by ID
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        id path string true "tweet_id"
// @Success      200  {object}  models.Tweet
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTweet(c *gin.Context) {
	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tweet, err := h.services.Tweets().Get(ctx, id.String())
	if err != nil {
		handleResponse(c, h.log, "error while getting tweet by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, tweet)
}

// GetTweetList godoc
// @Router       /tweets [GET]
// @Summary      Get tweet list
// @Description  Get list of tweets
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.TweetsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTweetList(c *gin.Context) {
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
	resp, err := h.services.Tweets().GetList(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting tweets", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "success!", http.StatusOK, resp)
}

// UpdateTweet godoc
// @Router       /tweet/{id} [PUT]
// @Summary      Update tweet
// @Description  Update a tweet
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        id path string true "tweet_id"
// @Param        tweet body models.UpdateTweet true "tweet"
// @Success      200  {object}  models.Tweet
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTweet(c *gin.Context) {
	updateTweet := models.UpdateTweet{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, h.log, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateTweet.ID = uid

	if err := c.ShouldBindJSON(&updateTweet); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Tweets().Update(ctx, updateTweet)
	if err != nil {
		handleResponse(c, h.log, "error while updating tweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// DeleteTweet godoc
// @Router       /tweet/{id} [DELETE]
// @Summary      Delete tweet
// @Description  Delete a tweet
// @Tags         tweet
// @Accept       json
// @Produce      json
// @Param        id path string true "tweet_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteTweet(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.Tweets().Delete(ctx, models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, h.log, "error while deleting tweet by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "data successfully deleted")
}
