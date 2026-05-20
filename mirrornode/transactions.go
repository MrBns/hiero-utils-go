package mirrornode

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

type GetTransactionOptions struct {
	Limit      int
	AccountId  string
	TimeStamps []string
	// asc or desc
	Order string
	// debit or credit
	BalanceChangeType string
	// either success of fail
	Result string
	// Transaction types
	TransactionTypes []string
}

type TransactionIterator struct {
	nextUrl string
	client  *Client
	count   int
}

func (iter *TransactionIterator) Next() bool {
	return iter.nextUrl != ""
}

func (iter *TransactionIterator) Values(ctx context.Context) ([]Transaction, error) {
	client := iter.client
	url, err := client.buildURL(iter.nextUrl)
	if err != nil {
		return nil, err
	}

	res, err := client.doGet(ctx, url)

	if err != nil {
		return nil, err
	}

	var data GetTransactionsResponse

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Links != nil && data.Links.Next != nil && len(data.Transactions) > 0 {
		iter.nextUrl = *data.Links.Next
		iter.count++
	} else {
		iter.nextUrl = ""
	}

	return data.Transactions, nil
}

func (c *Client) GetTransactionsIter(opt *GetTransactionOptions) *TransactionIterator {
	if opt == nil {
		opt = &GetTransactionOptions{Order: "desc", Limit: 25}
	}
	if opt.Limit <= 0 {
		opt.Limit = 25
	}

	var query = url.Values{
		"limit": []string{strconv.Itoa(opt.Limit)},
		"order": []string{"desc"},
	}

	if opt.AccountId != "" {
		query.Set("account.id", opt.AccountId)
	}
	if opt.BalanceChangeType != "" {
		query.Set("type", opt.BalanceChangeType)
	}
	if opt.Result != "" {
		query.Set("result", opt.Result)
	}
	if len(opt.TimeStamps) > 0 {
		for _, v := range opt.TimeStamps {
			query.Add("timestamp", v)
		}
	}
	if len(opt.TransactionTypes) > 0 {
		for _, v := range opt.TransactionTypes {
			query.Add("transactiontype", v)
		}
	}

	return &TransactionIterator{
		client:  c,
		nextUrl: "/api/v1/transactions?" + query.Encode(),
		count:   0,
	}
}

func (c *Client) GetAllTransaction(ctx context.Context, opt *GetTransactionOptions) ([]Transaction, error) {

	var iter = c.GetTransactionsIter(opt)
	var trx = []Transaction{}
	for iter.Next() {
		pulled, err := iter.Values(ctx)
		if err != nil {
			return nil, err
		}
		if len(pulled) <= 0 {
			break
		}
		trx = append(trx, pulled...)

	}

	return trx, nil
}
