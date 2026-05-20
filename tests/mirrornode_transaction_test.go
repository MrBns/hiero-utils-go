package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/mrbns/hiero-utils-go/mirrornode"
)

func TestTransactionIterator(t *testing.T) {
	client, err := mirrornode.NewClient(mirrornode.MainnetNetwork)

	if err != nil {
		t.Error(err)
	}

	iter := client.GetTransactionsIter(&mirrornode.GetTransactionOptions{
		AccountId: "0.0.10075301",
		Limit:     100,
	})

	for iter.Next() {
		val, err := iter.Values(t.Context())
		if err != nil {
			t.Error(err)
			break
		}
		dump.P(val)
	}

}

func TestGetAllTransaction(t *testing.T) {

	client, err := mirrornode.NewClient(mirrornode.MainnetNetwork)

	if err != nil {
		t.Error(err)
		return
	}

	timestamp := time.Now().Add(-(time.Hour * 24 * 70))

	txs, err := client.GetAllTransaction(t.Context(), &mirrornode.GetTransactionOptions{
		AccountId: "0.0.10075301",
		Limit:     100,
		TimeStamps: []string{
			fmt.Sprintf("gte:%d.00000000", timestamp.Unix()),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}
	dump.P(len(txs))
	dump.P(txs[10:])

}
