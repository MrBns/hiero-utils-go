package mirrornode

import (
	"context"
	"encoding/json"
	"fmt"
)

type GetTopicMessagesByIdOptions struct {
	Encoding       string
	Limit          int64
	SequenceNumber int64

	// The order in which items are listed either  "asc" or "desc"
	Order string

	// The consensus timestamp as a Unix timestamp in seconds.nanoseconds format with an
	Timestamp []string
}

func GetTopicMessagesById(topicID string, options GetTopicMessagesByIdOptions) (*GetTopicMessagesResponse, error) {
	return nil, nil
}

func (client *Client) GetTopicLastMessage(ctx context.Context, topicID string) (*TopicMessage, error) {
	resp, err := client.doGet(ctx, fmt.Sprintf("/api/v1/topics/%s/messages", topicID))

	if err != nil {
		return nil, err
	}

	var topicMessages GetTopicMessagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&topicMessages); err != nil {
		return nil, err
	}
	if len(topicMessages.Messages) == 0 {
		return nil, ErrNoTopicMessage
	}

	return &topicMessages.Messages[0], nil

}
