package middleware

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

const (
	ContextKeyUserID ContextKey = "user_id"
)
