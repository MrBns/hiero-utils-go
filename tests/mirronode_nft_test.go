package tests_test

import (
	"testing"

	"github.com/mrbns/hiero-utils-go/mirrornode"
)

func getClient(t *testing.T) *mirrornode.Client {
	mnClient, err := mirrornode.NewClient(
		mirrornode.MainnetNetwork,
	)
	if err != nil {
		t.Error(err)
	}

	return mnClient
}

func TestDownloadAllNft(t *testing.T) {

	mnClient := getClient(t)
	iterator := mnClient.GetNftCollection(t.Context(), "0.0.10075276", &mirrornode.GetNftCollectionOptions{
		Limit: 10,
	})

	total := 0
	index := 0
	for iterator.HasNext() {
		nfts, err := iterator.Next()
		if err != nil {
			t.Error(err)
			break
		}

		total += len(nfts)

		t.Logf("iteration index %v", iterator.Index())

		for _, n := range nfts {
			t.Logf("%v@%v", n.TokenID, n.SerialNumber)
		}

		index++

	}

	t.Logf("total nft %v \n", total)

}

func TestGetNftDetails(t *testing.T) {

	mnClient := getClient(t)

	nft, err := mnClient.GetNftInfo(t.Context(), "0.0.10075276@10")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Nft Info %#v\n", nft)

}

func TestGetNftMetadata(t *testing.T) {

	mnClient := getClient(t)

	metadata, err := mnClient.GetNftMetadataURL(t.Context(), "0.0.10075276@10")
	if err != nil {
		t.Error(err)
	}

	t.Logf("metadata %v\n", metadata)
}
