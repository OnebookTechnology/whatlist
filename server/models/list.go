package models

type List struct {
	ListID           uint64 `json:"list_id"` 				// 书单ID
	ListName         string `json:"list_name"`				// 书单名称
	ListAuthor       string `json:"list_author"`			// 书单作者
	ListCategoryID   uint64 `json:"list_category_id"`		// 书单分类ID
	ListIntro        string `json:"list_intro"`				// 书单简介
	ListImg          string `json:"list_img"`				// 书单图片
	ListCreateTime   string `json:"list_create_time"`		// 书单创建时间
	ListLastEditTime string `json:"list_last_edit_time"`	// 书单最后编辑时间
	ListClickCount   uint64 `json:"list_click_count"`		// 书单点击量
}
