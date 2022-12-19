package main

import (
	"context"
	"fmt"
	"encoding/json"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

const CreateTransactionQuery = `
INSERT INTO transactions (price, status) values ($1, $2) RETURNING id
`

const CreateEventQuery = `
INSERT INTO events (event_type, payload) values ($1, $2)
`

func (p Postgres) CreateTransaction(
	ctx context.Context,
	input Transaction,
) (string, error) {
	ctx, span := otel.Tracer("outbox-pattern").Start(ctx, "postgres.CreateTransaction")
	defer span.End()

	var output string
	
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("begin txx: %w", err)
	}
	
	err = tx.QueryRowContext(
		ctx, CreateTransactionQuery,
		input.Price,
		"PENDING",
	).Scan(&output)
	if err != nil {
		fmt.Println("insert transaction: ", err)
		err = fmt.Errorf("insert transaction: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		if errRollback := tx.Rollback(); err != nil {
			fmt.Println(errRollback)
			err = fmt.Errorf("rollback transaction: %w: %v", errRollback, err)
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			return "", errRollback
		}

		return "", err
	}

	payload, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf("insert transaction: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		if errRollback := tx.Rollback(); err != nil {
			fmt.Println(errRollback)
			err = fmt.Errorf("rollback transaction: %w: %v", errRollback, err)
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			return "", errRollback
		}

		return "", err
	}
	
	_, err = tx.ExecContext(
		ctx, CreateEventQuery,
		"transaction_created",
		payload,
	)
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf("insert event: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		if errRollback := tx.Rollback(); err != nil {
			fmt.Println(errRollback)
			err = fmt.Errorf("rollback transaction: %w: %v", errRollback, err)
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			return "", errRollback
		}

		return "", err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println(err)
		err = fmt.Errorf("commit transaction: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return "", err
	}

	return output, nil
}
