package tests_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/dump"
)

func TestGetLastTopic(t *testing.T) {
	mnClient := getClient(t)

	msg, err := mnClient.GetTopicLastMessage(t.Context(), "0.0.10078932")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("last secuence is %d", msg.SequenceNumber)
	dump.P(msg)
}
