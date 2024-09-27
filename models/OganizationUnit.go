package models

import "github.com/bwmarrin/snowflake"

type OrganizationUnit struct {
	Name            string             `json:"name" gorm:" type:varchar(255)"`
	DisplayName     string             `json:"text" gorm:" type:varchar(255)"`
	ParentId        snowflake.ID       `json:"parentId,omitempty" `
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
	return m.Id.Int64()
}

//	func (m *OrganizationUnit) SetID(id interface{}) {
//		m.Id = fmt.Sprintf("%v", id)
//	}
func (m *OrganizationUnit) ParentID() interface{} {
	result := m.ParentId.Int64()
	if result == 0 {
		return nil
	} else {
		return result
	}
}

func (m *OrganizationUnit) SetChildren(children []interface{}) {
	m.Children = make([]OrganizationUnit, 0)
	for _, c := range children {
		m.Children = append(m.Children, *c.(*OrganizationUnit))
	}
}
