package common

// message server需要用到的结构体

type MessageSendEvent struct {
	UserId     uint   `json:"user_id,omitempty"`
	ToUserId   uint   `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId uint   `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
