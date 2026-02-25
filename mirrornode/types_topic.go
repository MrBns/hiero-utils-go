package mirrornode

// TopicMessage
type TopicMessage struct {
	ChunkInfo          *ChunkInfo `json:"chunk_info"`
	ConsensusTimestamp string     `json:"consensus_timestamp"`
	Message            string     `json:"message"`
	PayerAccountID     *string    `json:"payer_account_id"`
	RunningHash        string     `json:"running_hash"`
	RunningHashVersion int64      `json:"running_hash_version"`
	SequenceNumber     int64      `json:"sequence_number"`
	TopicID            *string    `json:"topic_id"`
}

type ChunkInfo struct {
	InitialTransactionID *TransactionID `json:"initial_transaction_id,omitempty"`
	Number               *int64         `json:"number,omitempty"`
	Total                *int64         `json:"total,omitempty"`
	Scheduled            bool           `json:"scheduled"`
	Nonce                *int64         `json:"nonce"`
}

// TransactionId
type TransactionID struct {
	AccountID             *string `json:"account_id"`
	Nonce                 *int64  `json:"nonce"`
	Scheduled             *bool   `json:"scheduled"`
	TransactionValidStart *string `json:"transaction_valid_start,omitempty"`
}

// TopicMessagesResponse
type GetTopicMessagesResponse struct {
	Links    *Links         `json:"links,omitempty"`
	Messages []TopicMessage `json:"messages,omitempty"`
}
