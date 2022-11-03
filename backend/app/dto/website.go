package dto

import (
	"github.com/1Panel-dev/1Panel/backend/app/model"
	"github.com/1Panel-dev/1Panel/backend/utils/nginx/components"
)

type WebSiteReq struct {
	PageInfo
}

type AppType string

const (
	NewApp       AppType = "new"
	InstalledApp AppType = "installed"
)

type WebSiteCreate struct {
	PrimaryDomain  string        `json:"primaryDomain" validate:"required"`
	Type           string        `json:"type" validate:"required"`
	Alias          string        `json:"alias" validate:"required"`
	Remark         string        `json:"remark"`
	OtherDomains   string        `json:"otherDomains"`
	AppType        AppType       `json:"appType" validate:"required"`
	AppInstall     NewAppInstall `json:"appInstall"`
	AppID          uint          `json:"appID"`
	AppInstallID   uint          `json:"appInstallID"`
	WebSiteGroupID uint          `json:"webSiteGroupID" validate:"required"`
}

type NewAppInstall struct {
	Name        string                 `json:"name"`
	AppDetailId uint                   `json:"appDetailID"`
	Params      map[string]interface{} `json:"params"`
}

type WebSiteDel struct {
	ID           uint `json:"id"`
	DeleteApp    bool `json:"deleteApp"`
	DeleteBackup bool `json:"deleteBackup"`
}

type WebSiteDTO struct {
	model.WebSite
}

type WebSiteGroupCreate struct {
	Name string `json:"name"`
}

type WebSiteGroupUpdate struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type WebSiteDomainCreate struct {
	WebSiteID uint   `json:"webSiteId"`
	Port      int    `json:"port"`
	Domain    string `json:"domain"`
}

type NginxConfig struct {
	FilePath      string             `json:"filePath"`
	ContainerName string             `json:"containerName"`
	Config        *components.Config `json:"config"`
	OldContent    string             `json:"oldContent"`
}