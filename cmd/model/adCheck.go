package model

import "time"

type AcInviteCode struct {
	AcInviteCodeId int64     `gorm:"column:ac_invite_code_id;primary_key" json:"acInviteCodeId,string"`
	AppId          string    `gorm:"column:app_id" json:"appId"`
	InviteCode     string    `gorm:"column:invite_code" json:"inviteCode"`
	IsUsed         int       `gorm:"column:is_used" json:"isUsed"`
	BrowserId      string    `gorm:"column:browser_id" json:"browserId"`
	UsedTime       time.Time `gorm:"column:used_time" json:"usedTime"`
	CreateTime     time.Time `gorm:"column:create_time;<-:false" json:"-"`
	UpdateTime     time.Time `gorm:"column:update_time;<-:false" json:"-"`
}
