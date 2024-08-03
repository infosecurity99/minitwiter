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

// LikeTweet godoc
// @Router       /like [POST]
// @Summary      Like a tweet
// @Description  Allows a user to like a tweet
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        like body models.LikeRequest true "Like Request"
// @Success      201  {object}  models.Like
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) LikeTweet(c *gin.Context) {
	likeRequest := models.LikeRequest{}

	if err := c.ShouldBindJSON(&likeRequest); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Like().Like(ctx, likeRequest)
	if err != nil {
		handleResponse(c, h.log, "error while liking tweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// UnlikeTweet godoc
// @Router       /unlike [POST]
// @Summary      Unlike a tweet
// @Description  Allows a user to unlike a tweet
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        unlike body models.LikeRequest true "Unlike Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UnlikeTweet(c *gin.Context) {
	unlikeRequest := models.LikeRequest{}

	if err := c.ShouldBindJSON(&unlikeRequest); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Like().Unlike(ctx, unlikeRequest); err != nil {
		handleResponse(c, h.log, "error while unliking tweet", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "successfully unliked")
}

// GetLikes godoc
// @Router       /likes/{tweet_id} [GET]
// @Summary      Get likes
// @Description  Get a list of users who liked a tweet
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        tweet_id path string true "Tweet ID"
// @Success      200  {object}  models.LikesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetLikes(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	id, err := uuid.Parse(tweetID)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	likes, err := h.services.Like().GetLikes(ctx, id.String())
	if err != nil {
		handleResponse(c, h.log, "error while getting likes", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, likes)
}
