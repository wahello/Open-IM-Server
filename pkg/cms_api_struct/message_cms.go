package cms_api_struct

type BroadcastRequest struct {
	Message string `json:"message"`
}

type BroadcastResponse struct {
}

type MassSendMassageRequest struct {
}

type MassSendMassageResponse struct {
}

type GetChatLogsRequest struct {
	SessionType int    `form:"session_type"`
	ContentType int    `form:"content_type"`
	Content     string `form:"content"`
	UserId      string `form:"user_id"`
	GroupId     string `form:"group_id"`
	Date        string `form:"date"`

	RequestPagination
}

type ChatLog struct {
	SessionType    int    `json:"session_type"`
	ContentType    int    `json:"content_type"`
	SenderNickName string `json:"sender_nick_name"`
	SenderId       int    `json:"sender_id"`
	SearchContent  string `json:"search_content"`
	WholeContent   string `json:"whole_content"`

	ReceiverNickName string `json:"receiver_nick_name"`
	ReceiverID       int    `json:"receiver_id"`

	GroupName string `json:"group_name"`
	GroupId   string `json:"group_id"`

	Date string `json:"date"`
}

type GetChatLogsResponse struct {
	ChatLogs []ChatLog `json:"chat_logs"`

	ResponsePagination
}
