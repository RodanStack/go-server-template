package serializer

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Status   string `json:"status" binding:"required,oneof=active inactive suspended"`
}
