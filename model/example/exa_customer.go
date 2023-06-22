package example

import "github.com/wubainian/godaemon/global"

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string `json:"customerName" form:"customerName" gorm:"comment:客户名"`                // 客户名
	CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`    // 客户手机号
	SysUserID          uint   `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                     // 管理ID
	SysUserAuthorityID uint   `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"` // 管理角色ID                    // 管理详情
}
