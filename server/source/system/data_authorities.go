package system

import (
	"reflect"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DataAuthorities = new(dataAuthorities)

type dataAuthorities struct{}

func (a *dataAuthorities) TableName() string {
	var entity AuthoritiesResources
	return entity.TableName()
}

func (a *dataAuthorities) Initialize() error {
	entities := []AuthoritiesResources{
		{AuthorityId: 1, ResourcesId: 1},
		{AuthorityId: 1, ResourcesId: 2},
		{AuthorityId: 1, ResourcesId: 3},
		{AuthorityId: 2, ResourcesId: 2},
		{AuthorityId: 2, ResourcesId: 3},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *dataAuthorities) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("sys_authority_authority_id = ? AND data_authority_id_authority_id = ?", 3, 3).First(&AuthoritiesResources{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// AuthoritiesResources 角色资源表
type AuthoritiesResources struct {
	AuthorityId uint `gorm:"column:sys_authority_authority_id"`
	ResourcesId uint `gorm:"column:data_authority_id_authority_id"`
}

func (a *AuthoritiesResources) TableName() string {
	var entity system.SysAuthority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("DataAuthorityId"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}
