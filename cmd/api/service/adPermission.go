package service

import (
	"adCheck/cmd/common"
	"adCheck/cmd/model"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"reflect"
	"time"
)

func GenerateCode(context *gin.Context, bo model.GeneralCodeBO) *model.Result {
	// 定义可用字符集：小写字母和数字
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 8)

	for i := range result {
		// 生成一个小于字符集长度的随机数
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		// 从字符集中选择一个字符
		result[i] = charset[n.Int64()]
	}
	newCode := string(result)
	db := common.GetDb()
	var count int64
	db.Model(model.AcInviteCode{}).Where("invite_code = ?", newCode).Count(&count)
	if count == 0 {
		db.Model(model.AcInviteCode{}).Create(&model.AcInviteCode{
			AppId:      bo.AppId,
			InviteCode: newCode,
		})
		return model.Success(newCode)
	} else {
		return model.Failed(http.StatusInternalServerError, "邀请码生成重复，请再试一次", nil)
	}

}

/*
*
40001 邀请码不存在
40002 邀请码不属于该应用
40003 邀请码已绑定其他用户
*/
func CheckCode(context *gin.Context, bo model.CheckCodeBO) *model.Result {
	db := common.GetDb()
	var inviteCode model.AcInviteCode
	db.Model(model.AcInviteCode{}).Where("invite_code = ?", bo.InviteCode).First(&inviteCode)
	if reflect.ValueOf(inviteCode).IsZero() {
		return model.Failed(http.StatusBadRequest, "40001", false)
	}
	if inviteCode.AppId != bo.AppId {
		return model.Failed(http.StatusBadRequest, "40002", false)
	}

	if inviteCode.IsUsed == 1 {
		if inviteCode.BrowserId != bo.BrowserId {
			return model.Failed(http.StatusBadRequest, "40003", false)
		} else {
			return model.Success(true)
		}
	} else {
		db.Debug().Model(model.AcInviteCode{}).Where("invite_code = ?", bo.InviteCode).UpdateColumns(model.AcInviteCode{
			IsUsed:    1,
			BrowserId: bo.BrowserId,
			UsedTime:  time.Now(),
		})
		return model.Success(true)
	}

}
