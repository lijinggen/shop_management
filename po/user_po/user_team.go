package user_po

import (
	"github.com/shop_management/po/common_po"
)

type SubUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type SubUserListResp struct {
	Pager *common_po.Pager `json:"pager"`
	List  []*SubUser       `json:"list"`
}
