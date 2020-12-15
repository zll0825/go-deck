package service

import (
	"go-deck/app/global"
	"go-deck/app/model/entity"
)

func GetMenuTree() (total int, data interface{}, err error) {
	allMenus , err := global.DB.AllMenus()
	if err != nil {
		return 0, nil, err
	}

	// 先用parentId构建map
	treeMap := map[int][]entity.Menu{}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}

	// 递归构建所有紫菜单
	baseMenu := treeMap[0]
	total = len(baseMenu)
	for i := 0; i < total; i++ {
		err = getChildrenList(&baseMenu[i], treeMap)
	}

	return total, baseMenu, err
}

func getChildrenList(menu *entity.Menu, treeMap map[int][]entity.Menu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
