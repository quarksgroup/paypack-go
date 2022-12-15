package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/quarksgroup/paypack-go/paypack"
)

func (c *Client) Cashin(ctx context.Context, tx *paypack.TransactionRequest) (*paypack.TransactionResponse, error) {
	const endpoint = "transactions/cashin"

	in := &transactionRequest{
		Amount: tx.Amount,
		Number: tx.Number,
	}

	header := http.Header{
		"X-Webhook-Mode": []string{tx.Mode},
	}

	out := new(paypack.TransactionResponse)

	_, err := c.do(ctx, "POST", endpoint, in, out, header)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) Cashout(ctx context.Context, tx *paypack.TransactionRequest) (*paypack.TransactionResponse, error) {
	const endpoint = "transactions/cashout"

	in := &transactionRequest{
		Amount: tx.Amount,
		Number: tx.Number,
	}

	header := http.Header{
		"X-Webhook-Mode": []string{tx.Mode},
	}

	out := new(paypack.TransactionResponse)

	_, err := c.do(ctx, "POST", endpoint, in, out, header)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) FindTx(ctx context.Context, ref string) (*paypack.Transaction, error) {

	endpoint := fmt.Sprintf("transactions/find/%s", ref)

	out := new(Transaction)

	_, err := c.do(ctx, "GET", endpoint, nil, out, nil)
	if err != nil {
		return nil, err
	}

	res := &paypack.Transaction{
		Ref:       out.Ref,
		Amount:    out.Amount,
		Kind:      out.Kind,
		Fee:       out.Fee,
		Provider:  out.Provider,
		Client:    out.Client,
		CreatedAt: out.Timestamp,
	}

	return res, nil
}

// List handles List http api request for https://payments.paypack.rw/api/transactions/list with paramas
func (c *Client) ListTx(ctx context.Context, options ...paypack.Option) (*paypack.Transactions, error) {

	var params string

	if len(options) > 0 {
		for _, option := range options {
			params += fmt.Sprintf("%s&", option)
		}
	}

	endpoint := fmt.Sprintf("transactions/list?%s", params)

	out := new(listTransactions)

	_, err := c.do(ctx, "GET", endpoint, nil, out, nil)
	if err != nil {
		return nil, err
	}

	res := &paypack.Transactions{
		Offset:       out.Offset,
		Limit:        out.Limit,
		From:         out.From,
		To:           out.To,
		Kind:         out.Kind,
		Fee:          out.Fee,
		Cashin:       out.Cashin,
		Cashout:      out.Cashout,
		Total:        out.Total,
		Transactions: make([]paypack.Transaction, 0),
	}

	for _, tx := range out.Transactions {
		res.Transactions = append(res.Transactions, paypack.Transaction{
			Ref:       tx.Ref,
			Amount:    tx.Amount,
			Kind:      tx.Kind,
			Provider:  tx.Provider,
			Fee:       tx.Fee,
			Client:    tx.Client,
			CreatedAt: tx.Timestamp,
		})
	}

	return res, nil
}
