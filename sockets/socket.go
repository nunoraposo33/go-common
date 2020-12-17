package sockets

// SocketMessage Describes the messages sent by socket between agents
type SocketMessage struct {
	Command           string `json:"command"`
	OriginalMessageID string `json:"originalMessageId"`
	MessageID         string `json:"messageId"`
	Timestamp         string `json:"timestamp"`
	Payload           string `json:"payload"`
}

// SyncMessage Messaged used to sync the plataform with ERP's
type SyncMessage struct {
	ID        int    `json:"id"`
	Command   string `json:"command"`
	Response  string `json:"response"`
	MessageID string `json:"message_id"`
	ClientID  string `json:"client_id"`
	Processed bool   `json:"processed"`
	Payload   string `json:"payload"`
}
