package models

import (
	"gopkg.in/guregu/null.v4"
)

type OrganizationUnit struct {
	Name            string             `json:"name" gorm:" type:varchar(255)"`
	DisplayName     string             `json:"text" gorm:" type:varchar(255)"`
	ParentId        null.Int           `json:"parentId" `
	SortOrder       int                `json:"sortorder"`
	Path            string             `json:"path" gorm:"type:varchar(2048)"`
	Children        []OrganizationUnit `json:"nodes" gorm:"foreignkey:ParentId;association_foreignkey:Id"`
	SnowflakeRecord `gorm:"embedded"`
}

// //TableName 数据表名称
//
//	func (m *OrganizationUnit) TableName() string {
//		return "OrganizationUnit"
//	}
func (m *OrganizationUnit) GetID() interface{} {
	return m.Id
}

//	func (m *OrganizationUnit) SetID(id interface{}) {
//		m.Id = fmt.Sprintf("%v", id)
//	}
func (m *OrganizationUnit) ParentID() interface{} {
	if !m.ParentId.Valid {
		return ""
	}
	return m.ParentId
}

func (m *OrganizationUnit) SetChildren(children []interface{}) {
	m.Children = make([]OrganizationUnit, 0)
	for _, c := range children {
		m.Children = append(m.Children, *c.(*OrganizationUnit))
	}
}
