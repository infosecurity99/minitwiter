package handler

import (
	"context"
	"net/http"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// FollowUser godoc
// @Router       /follow [POST]
// @Summary      Follow a user
// @Description  Allows a user to follow another user
// @Tags         followers
// @Accept       json
// @Produce      json
// @Param        follow body models.GetListRequest true "Follow Request"
// @Success      201  {object}  models.Follower
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) FollowUser(c *gin.Context) {
	followRequest := models.GetListRequest{}

	if err := c.ShouldBindJSON(&followRequest); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Follower().Follow(ctx, followRequest)
	if err != nil {
		handleResponse(c, h.log, "error while following user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// UnfollowUser godoc
// @Router       /unfollow [POST]
// @Summary      Unfollow a user
// @Description  Allows a user to unfollow another user
// @Tags         followers
// @Accept       json
// @Produce      json
// @Param        unfollow body models.FollowRequest true "Unfollow Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UnfollowUser(c *gin.Context) {
	unfollowRequest := models.FollowRequest{}

	if err := c.ShouldBindJSON(&unfollowRequest); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Follower().Unfollow(ctx, unfollowRequest); err != nil {
		handleResponse(c, h.log, "error while unfollowing user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "successfully unfollowed")
}

// GetFollowers godoc
// @Router       /followers/{user_id} [GET]
// @Summary      Get followers
// @Description  Get a list of followers for a user
// @Tags         followers
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200  {object}  models.FollowersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetFollowers(c *gin.Context) {
	uid := c.Param("user_id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	followers, err := h.services.Follower().GetFollowers(ctx, id.String())
	if err != nil {
		handleResponse(c, h.log, "error while getting followers", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, followers)
}

// GetFollowing godoc
// @Router       /following/{user_id} [GET]
// @Summary      Get following
// @Description  Get a list of users that a user is following
// @Tags         followers
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200  {object}  models.FollowingResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetFollowing(c *gin.Context) {
	uid := c.Param("user_id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	following, err := h.services.Follower().GetFollowing(ctx, id.String())
	if err != nil {
		handleResponse(c, h.log, "error while getting following list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, following)
}
