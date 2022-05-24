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
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *service) Login(ctx context.Context, username, password string) (pkg.ClientInfo, error) {
	staff, err := s.repo.StaffLogin(ctx, model.StaffLoginParams{
		Username: username,
		Password: helper.EncryptPassword(password, s.config.Salt),
	})
	if err != nil {
		return pkg.ClientInfo{}, err
	}

	now := time.Now().UTC()
	if err = s.repo.UpdateStaffLogin(ctx, model.UpdateStaffLoginParams{
		ServingStatus: types.StaffServingStatusClosed,
		LastLoginTime: sql.NullTime{Time: now, Valid: true},
		UpdatedAt:     now,
		ID:            staff.ID,
	}); err != nil {
		return pkg.ClientInfo{}, err
	}

	token := genToken()

	staffInfo := pkg.ClientInfo{
		ID:            staff.ID,
		Type:          pkg.ClientTypeStaff,
		Name:          staff.Name,
		Username:      staff.Username,
		ServingStatus: types.StaffServingStatusClosed,
		Token:         token,
	}

	result, err := json.Marshal(staffInfo)
	if err != nil {
		return pkg.ClientInfo{}, err
	}

	err = s.lua.RemoveToken(ctx, "staff", staff.Username)
	if err != nil {
		log.Error().Msgf("clear token error: %s\n", err)
		return pkg.ClientInfo{}, err
	}

	err = s.lua.SetToken(ctx, "staff", staff.Username, token, result, 24*time.Hour)
	if err != nil {
		log.Error().Msgf("set token error: %s\n", err)
		return pkg.ClientInfo{}, err
	}

	return staffInfo, nil
}

func (s *service) Logout(ctx context.Context, staffInfo pkg.ClientInfo) error {
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

func (s *service) SetClientInfo(clientType pkg.ClientType) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxKey := fmt.Sprintf("%s_info", clientType)
		token := c.GetHeader("X-Token")
		if token == "" {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		redisKey := getRedisKey(token, clientType)
		result, err := s.redis.Get(c.Request.Context(), redisKey)
		if err != nil {
			ginTool.ErrorAuth(c)
			c.Abort()
			return
		}

		var tmp pkg.ClientInfo
		err = json.Unmarshal([]byte(result), &tmp)
		if err != nil {
			ginTool.Error(c, err)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), ctxKey, tmp)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *service) GetClientInfo(ctx context.Context, clientType pkg.ClientType) (pkg.ClientInfo, error) {
	ctxKey := fmt.Sprintf("%s_info", clientType)
	clientInfo, ok := ctx.Value(ctxKey).(pkg.ClientInfo)
	if clientInfo.ID == 0 || !ok {
		return pkg.ClientInfo{}, errors.ErrorAuth
	}

	return clientInfo, nil
}

func getRedisKey(token string, clientType pkg.ClientType) string {
	return fmt.Sprintf("token:%s:%s", clientType, token)
}

func genToken() string {
	str := time.Now().String() + xid.New().String()
	str = fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return str
}
