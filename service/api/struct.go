package api



type LoginRequest struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type LoginResponse struct {
	Identifier string `json:"identifier"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type UpdateGroupRequest struct {
	Name string `json:"groupName"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Photo []byte `json:"photo"`
}

type Message struct {
    Id              string  `json:"id"`
    ConversationId  string  `json:"conversationId"`
    SenderId        string  `json:"senderId"`
    SenderName      string  `json:"senderName"`
    Content         string  `json:"content"`
    Timestamp       string  `json:"timestamp"`
    Attachment      []byte  `json:"attachment"`

    ReplyTo         string  `json:"replyTo"` // 🔥 ВАЖНО
    ReplyContent    string  `json:"replyContent"`
    ReplySenderName string  `json:"replySenderName"`
    ReplyAttachment []byte  `json:"replyAttachment"`
}




