package models

import (
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type OrganizationUnit struct {
	Id          string             `json:"id" gorm:"primary_key;type:varchar(36);not null"`
	Name        string             `json:"name" gorm:" type:varchar(255)"`
	DisplayName string             `json:"text" gorm:" type:varchar(255)"`
	Parent      null.String        `json:"parent" gorm:"type:varchar(36)"`
	SortOrder   int                `json:"sortorder"`
	Path        string             `json:"path" gorm:"type:varchar(2048)"`
	Children    []OrganizationUnit `json:"nodes" gorm:"foreignkey:Parent;association_foreignkey:Id"`
	Record      `gorm:"embedded"`
}

// //TableName 数据表名称
//
//	func (m *OrganizationUnit) TableName() string {
//		return "OrganizationUnit"
//	}
func (m *OrganizationUnit) GetID() interface{} {
	return m.Id
}
func (m *OrganizationUnit) SetID(id interface{}) {
	m.Id = fmt.Sprintf("%v", id)
}
func (m *OrganizationUnit) ParentID() interface{} {
	return m.Parent
}

func (m *OrganizationUnit) SetChildren(children []interface{}) {
	m.Children = make([]OrganizationUnit, 0)
	for _, c := range children {
		m.Children = append(m.Children, *c.(*OrganizationUnit))
	}
}
