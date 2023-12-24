package user_po

type UserLogin struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserReq struct {
	Name            string `json:"name" binding:"required,min=3,max=30" `
	Phone           string `json:"phone" binding:"required"`
	Password        string `json:"password" binding:"required,min=6,max=16"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,max=16"`
}

type UserLoginResponse struct {
	UserId string `json:"user_id"`
}

type UserProfile struct {
	AvatarUrl string `json:"avatar_image_url"`
	Name      string `json:"user_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserId    string `json:"user_id" `
	IsAdmin   bool   `json:"is_admin"`
}

//
//type ModifyPasswordReq struct {
//	UserId          string `json:"user_id" binding:"required"`
//	OldPassword     string `json:"old_password" binding:"required,min=6,max=16"`
//	NewPassword     string `json:"new_password" binding:"required,min=6,max=16"`
//	ConfirmPassword string `json:"confirm_password" binding:"required,min=6,max=16"`
//}
//
//type DeactivateReq struct {
//	UserId string `json:"user_id" binding:"required"`
//}
//
//type GenerateApiSecret struct {
//	UserId string `json:"user_id"`
//	ApiKey string `json:"api_key"`
//}
//
//type GetApiSecret struct {
//	ApiKey    string `json:"api_key"`
//	ApiSecret string `json:"api_secret"`
//}
//
//type UserListReq struct {
//	Pager *common_po.Pager `json:"pager"`
//	Name  string           `json:"name" form:"name"`
//}
//
//type User struct {
//	Id          string    `json:"id"`
//	Name        string    `json:"name"`
//	Email       string    `json:"email"`
//	Password    string    `json:"password"`
//	ApiKey      string    `json:"api_key"`
//	ApiSecret   string    `json:"api_secret"`
//	AvatarUrl   string    `json:"avatar_url"`
//	PhoneNumber string    `json:"phone_number"`
//	Balance     float64   `json:"balance"`
//	Currency    string    `json:"currency"`
//	CreateTime  time.Time `json:"create_time"`
//	ModifyTime  time.Time `json:"modify_time"`
//}
//
//type UserListResp struct {
//	Pager *common_po.Pager `json:"pager"`
//	List  []*User          `json:"list"`
//}
//
//type SaveUserReq struct {
//	Id          string  `json:"id"`
//	Name        string  `json:"name"`
//	Email       string  `json:"email"`
//	Password    string  `json:"password"`
//	PhoneNumber string  `json:"phone_number"`
//	Balance     float64 `json:"balance"`
//	Currency    string  `json:"currency"`
//}
//
//type UserBalance struct {
//	Id         string    `json:"id"`
//	UserId     string    `json:"userId"`
//	Balance    float64   `json:"balance"`
//	Currency   string    `json:"currency"`
//	CreateTime time.Time `json:"create_time"`
//	ModifyTime time.Time `json:"modify_time"`
//}
//
//type GetStatResp struct {
//	TotalUsage            float64 `json:"total_usage"`
//	TotalUsageCurrency    string  `json:"total_usage_currency"`
//	LastWeekUsage         float64 `json:"last_week_usage"`
//	LastWeekUsageCurrency string  `json:"last_week_usage_currency"`
//}
//
//type GetTrendResp struct {
//	Data []*TrendItem `json:"data"`
//}
//
//type TrendItem struct {
//	Day        string  `json:"day"`
//	TotalCount float64 `json:"total_count"`
//}
