package system

import (
	"reflect"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	var entity AuthorityMenus
	return entity.TableName()
}

func (a *authoritiesMenus) Initialize() error {
	menuBlocks := make(map[uint][]string)
	menuBlocks[1] = []string{
		MenuedDashboard,
		MenuedAbout,
		MenuedAdmin,
		MenuedAuthority,
		MenuedMenu,
		MenuedApi,
		MenuedUser,
		MenuedDictionary,
		MenuedDictionaryDetail,
		MenuedOperation,
		MenuedPerson,
		MenuedExample,
		MenuedExcel,
		MenuedUpload,
		MenuedBreakpoint,
		MenuedCustomer,
		MenuedSimpleUploader,
		MenuedSystemTools,
		MenuedAutoCodeAdmin,
		MenuedAutoCode,
		MenuedFormCreate,
		MenuedSystem,
		MenuedHome,
		MenuedState,
		MenuedAutoCodeEdit,
	}
	menuBlocks[2] = []string{
		MenuedDashboard,
		MenuedExample,
	}
	menus := getSysBaseMenuData()
	var entities []AuthorityMenus
	for authorityId, blocks := range menuBlocks {
		for _, block := range blocks {
			for _, menu := range menus {
				if menu.Name == block {
					entities = append(entities, AuthorityMenus{
						AuthorityId: authorityId,
						BaseMenuId:  menu.ID,
					})
					continue
				}
			}
		}

	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(
		global.
			GVA_DB.
			Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?", 17, 3).
			First(&AuthorityMenus{}).
			Error,
		gorm.ErrRecordNotFound,
	) { // 判断是否存在数据
		return false
	}
	return true
}

type AuthorityMenus struct {
	AuthorityId uint `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint `gorm:"column:sys_base_menu_id"`
}

func (a *AuthorityMenus) TableName() string {
	var entity system.SysAuthority
	types := reflect.TypeOf(entity)
	if s, o := types.FieldByName("SysBaseMenus"); o {
		m1 := schema.ParseTagSetting(s.Tag.Get("gorm"), ";")
		return m1["MANY2MANY"]
	}
	return ""
}
