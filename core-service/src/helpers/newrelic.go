package helpers

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func GetCtxNewRelic(app *newrelic.Application, nameTransaction string) (context.Context, *newrelic.Transaction) {
	txn := app.StartTransaction(nameTransaction)

	ctx := newrelic.NewContext(context.Background(), txn)

	return ctx, txn
}
