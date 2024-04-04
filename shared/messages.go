package shared

const (
	LoginActionName            = "login"
	LogoutActionName           = "logout"
	GetUsersActionName         = "users"
	GetUsersCountActionName    = "users_count"
	SendMessageActionName      = "send_message"
	NewMessageNotificationName = "new_message_notification"
	NewClientNotificationName  = "new_client_notification"
)

type Request struct {
	Name    string `json:"name"`
	Payload string `json:"payload,omitempty"`
}

type Response struct {
	Name    string      `json:"name"`
	Err     bool        `json:"error"`
	Payload interface{} `json:"payload"`
}

type TextMessage struct {
	Username  string    `json:"username"`
	UserColor ColorCode `json:"user_color"`
	Message   string    `json:"message"`
}
