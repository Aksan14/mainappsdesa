package model

type MstRole struct {
    IdRole   string `json:"id_role"`
    RoleName string `json:"role_name"`
    IsAdmin  bool   `json:"is_admin"`
}