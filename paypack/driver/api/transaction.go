package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/quarksgroup/paypack-go/paypack"
)

type transactionService struct {
	http *wrapper
}

func (s *transactionService) Cashin(ctx context.Context, tx *paypack.TxPayload) (*paypack.TxResponse, error) {
	endpoint := "transactions/cashin"

	in := &transactionRequest{
		Amount: tx.Amount,
		Number: tx.Number,
	}

	if detectProvider(in.Number) == airtel {
		return nil, &paypack.Error{Code: http.StatusBadRequest, Message: "airtel not currently supported"}
	}

	if detectProvider(in.Number) == uknown {
		return nil, &paypack.Error{Code: http.StatusBadRequest, Message: "unsupported provider"}
	}
	out := new(paypack.TxResponse)
	_, err := s.http.do(ctx, "POST", endpoint, in, out)
	return out, err
}

func (s *transactionService) Cashout(ctx context.Context, tx *paypack.TxPayload) (*paypack.TxResponse, error) {
	endpoint := "transactions/cashout"

	in := &transactionRequest{
		Amount: tx.Amount,
		Number: tx.Number,
	}

	if detectProvider(in.Number) == airtel {
		return nil, &paypack.Error{Code: http.StatusBadRequest, Message: "airtel not currently supported"}
	}

	if detectProvider(in.Number) == uknown {
		return nil, &paypack.Error{Code: http.StatusBadRequest, Message: "unsupported provider"}
	}

	out := new(paypack.TxResponse)
	_, err := s.http.do(ctx, "POST", endpoint, in, out)
	return out, err
}

func (s *transactionService) Find(ctx context.Context, ref string) (*paypack.Transaction, error) {
	endpoint := fmt.Sprintf("transactions/find/%s", ref)

	out := new(findTransactionResponse)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out)

	res := &paypack.Transaction{
		Ref:     out.Ref,
		Amount:  out.Amount,
		Kind:    out.Kind,
		Fee:     out.Fee,
		Client:  out.Client,
		Created: out.Timestamp,
	}
	return res, err
}

// List handles List http api request for https://payments.paypack.rw/api/transactions/list with paramas
func (s *transactionService) List(ctx context.Context, options ...string) (*paypack.Transactions, error) {

	var params string

	if len(options) > 0 {
		for _, option := range options {
			params += fmt.Sprintf("%s&", option)
		}
	}

	endpoint := fmt.Sprintf("transactions/list?%s", params)

	out := new(listTransactionResponse)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out)

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
			Ref:     tx.Ref,
			Amount:  tx.Amount,
			Kind:    tx.Kind,
			Fee:     tx.Fee,
			Client:  tx.Client,
			Created: tx.Timestamp,
		})
	}

	return res, err
}

var _ (paypack.TransactionService) = (*transactionService)(nil)
