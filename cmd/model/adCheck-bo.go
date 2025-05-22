package model

type GeneralCodeBO struct {
	AppId string `json:"appId" binding:"required,min=6"`
}

type CheckCodeBO struct {
	BrowserId  string `json:"browserId" binding:"required,min=6"`
	AppId      string `json:"appId" binding:"required,min=6"`
	InviteCode string `json:"inviteCode" binding:"required,min=6"`
}
