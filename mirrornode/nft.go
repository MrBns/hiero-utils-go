package mirrornode

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	hiero "github.com/hiero-ledger/hiero-sdk-go/v2/sdk"
)

type NFTCollectionPage struct {
	NFTs  []NFT
	Next  *string
	Limit int
}

// Get Nft Info.
func (c *Client) GetNftInfo(ctx context.Context, nftID string) (*NFT, error) {
	if c == nil {
		return nil, errors.New("mirror node client is nil")
	}

	parsedNftID, err := hiero.NftIDFromString(nftID)
	if err != nil {
		return nil, fmt.Errorf("invalid nft id %q: %w", nftID, err)
	}

	path := fmt.Sprintf("/api/v1/tokens/%s/nfts/%d", parsedNftID.TokenID, parsedNftID.SerialNumber)
	resp, err := c.doGet(ctx, path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nft NFT
	if err := json.NewDecoder(resp.Body).Decode(&nft); err != nil {
		return nil, fmt.Errorf("decode nft response: %w", err)
	}

	return &nft, nil
}

// Get Nft metadata URL.
func (c *Client) GetNftMetadataURL(ctx context.Context, nftID string) (string, error) {
	nft, err := c.GetNftInfo(ctx, nftID)
	if err != nil {
		return "", err
	}

	if nft.Metadata == nil || *nft.Metadata == "" {
		return "", errors.New("nft metadata is empty")
	}

	rawMetadata := strings.TrimSpace(*nft.Metadata)
	decoded, err := base64.StdEncoding.DecodeString(rawMetadata)
	if err == nil {
		decodedMetadata := strings.TrimSpace(string(decoded))
		if decodedMetadata != "" {
			return decodedMetadata, nil
		}
	}

	return rawMetadata, nil
}

// Iterator Of Nft Collection. take parameter how much item should be loaded of 10-100 per iteration.
type NftCollectionIterator struct {
	path             string
	client           *Client
	limit            int
	ctx              context.Context
	iter_count       int
	isFirstIteration bool
}

func (iter *NftCollectionIterator) HasNext() bool {
	return iter.path != ""
}

func (iter *NftCollectionIterator) Next() ([]NFT, error) {

	if iter.client == nil {
		return nil, errors.New("mirror node client is nil")
	}

	if iter.path == "" {
		return []NFT{}, nil
	}

	resp, err := iter.client.doGet(iter.ctx, iter.path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload NFTsResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}

	// terminate on empty page to avoid an extra "phantom" iteration
	if len(payload.NFTs) == 0 {
		iter.path = ""
		return []NFT{}, nil
	}

	if payload.Links.Next == nil {
		iter.path = ""
		return payload.NFTs, nil
	} else {
		iter.path = *payload.Links.Next
	}

	iter.iter_count++

	return payload.NFTs, nil
}

func (iter *NftCollectionIterator) Index() int {
	return iter.iter_count
}

type GetNftCollectionOptions struct {
	Limit        int
	Order        string
	SerialNumber int
	AccountId    string
}

func (c *Client) GetNftCollection(ctx context.Context, tokenID string, opt *GetNftCollectionOptions) *NftCollectionIterator {

	if opt == nil {
		opt = &GetNftCollectionOptions{
			Limit: 25,
			Order: "asc",
		}
	} else {
		if opt.Limit > 100 {
			opt.Limit = 100
		}

		if opt.Order == "" {
			opt.Order = "asc"
		}
	}

	// Path string builder
	pB := strings.Builder{}
	pB.WriteString("/api/v1/tokens/")
	pB.WriteString(tokenID)
	pB.WriteString("/nfts?limit=")
	pB.WriteString(strconv.Itoa(opt.Limit))
	pB.WriteString("&order=")
	pB.WriteString(opt.Order)

	if opt.AccountId != "" {
		pB.WriteString("&account.id=")
		pB.WriteString(opt.AccountId)
	}

	if opt.SerialNumber > 0 {
		pB.WriteString("&serialnumber=")
		pB.WriteString(strconv.Itoa(opt.SerialNumber))
	}

	return &NftCollectionIterator{
		ctx:              ctx,
		path:             pB.String(),
		limit:            opt.Limit,
		client:           c,
		iter_count:       0,
		isFirstIteration: false,
	}

}
