package auth

import (
	"context"
	"crypto/md5"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *service) Login(ctx context.Context, username, password string) (pkg.StaffInfo, error) {
	staff, err := s.repo.StaffLogin(ctx, model.StaffLoginParams{
		Username: username,
		Password: helper.EncryptPassword(password, s.config.Salt),
	})
	if err != nil {
		return pkg.StaffInfo{}, err
	}

	now := time.Now().UTC()
	if err = s.repo.UpdateStaffLogin(ctx, model.UpdateStaffLoginParams{
		ServingStatus: types.StaffServingStatusClosed,
		LastLoginTime: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            staff.ID,
	}); err != nil {
		return pkg.StaffInfo{}, err
	}

	token := genToken()

	staffInfo := pkg.StaffInfo{
		ID:            staff.ID,
		Type:          pkg.WsClientTypeStaff,
		Name:          staff.Name,
		Username:      staff.Username,
		ServingStatus: types.StaffServingStatusClosed,
		Token:         token,
	}

	result, err := json.Marshal(staffInfo)
	if err != nil {
		return pkg.StaffInfo{}, err
	}

	err = s.lua.RemoveToken(ctx, "staff", staff.Username)
	if err != nil {
		log.Error().Msgf("clear token error: %s\n", err)
		return pkg.StaffInfo{}, err
	}

	err = s.lua.SetToken(ctx, "staff", staff.Username, token, result, 24*time.Hour)
	if err != nil {
		log.Error().Msgf("set token error: %s\n", err)
		return pkg.StaffInfo{}, err
	}

	return staffInfo, nil
}

func (s *service) Logout(ctx context.Context, staffInfo pkg.StaffInfo) error {
	params := model.UpdateStaffServingStatusParams{
		ServingStatus: types.StaffServingStatusClosed,
		UpdatedBy:     staffInfo.ID,
		UpdatedAt:     time.Now().UTC(),
		ID:            staffInfo.ID,
	}
	if err := s.repo.UpdateStaffServingStatus(ctx, params); err != nil {
		return err
	}

	if err := s.lua.RemoveToken(ctx, "staff", staffInfo.Username); err != nil {
		log.Error().Msgf("clear token error: %s\n", err)
		return err
	}

	return nil
}

func (s *service) SetStaffInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if token == "" {
			ctx := context.WithValue(c.Request.Context(), "staff_info", pkg.StaffInfo{})
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}

		redisKey := getRedisKey(token)
		result, err := s.redis.Get(c.Request.Context(), redisKey)
		if err != nil {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		var tmp pkg.StaffInfo
		err = json.Unmarshal([]byte(result), &tmp)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "staff_info", tmp)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *service) GetStaffInfo(ctx context.Context) (pkg.StaffInfo, error) {
	staffInfo, ok := ctx.Value("staff_info").(pkg.StaffInfo)
	if staffInfo.ID == 0 || !ok {
		return pkg.StaffInfo{}, errors.ErrorAuth
	}

	return staffInfo, nil
}

func getRedisKey(token string) string {
	return fmt.Sprintf("token:staff:%s", token)
}

func genToken() string {
	str := time.Now().String() + xid.New().String()
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return str
}
