// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkUserExistsStmt, err = db.PrepareContext(ctx, checkUserExists); err != nil {
		return nil, fmt.Errorf("error preparing query CheckUserExists: %w", err)
	}
	if q.createPasswordResetTokenStmt, err = db.PrepareContext(ctx, createPasswordResetToken); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePasswordResetToken: %w", err)
	}
	if q.createSessionStmt, err = db.PrepareContext(ctx, createSession); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSession: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteExpiredSessionsStmt, err = db.PrepareContext(ctx, deleteExpiredSessions); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExpiredSessions: %w", err)
	}
	if q.deleteSessionStmt, err = db.PrepareContext(ctx, deleteSession); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSession: %w", err)
	}
	if q.deleteUserSessionsStmt, err = db.PrepareContext(ctx, deleteUserSessions); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserSessions: %w", err)
	}
	if q.getSessionStmt, err = db.PrepareContext(ctx, getSession); err != nil {
		return nil, fmt.Errorf("error preparing query GetSession: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.resetPasswordStmt, err = db.PrepareContext(ctx, resetPassword); err != nil {
		return nil, fmt.Errorf("error preparing query ResetPassword: %w", err)
	}
	if q.updateSessionStmt, err = db.PrepareContext(ctx, updateSession); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSession: %w", err)
	}
	if q.verifyUserStmt, err = db.PrepareContext(ctx, verifyUser); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkUserExistsStmt != nil {
		if cerr := q.checkUserExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkUserExistsStmt: %w", cerr)
		}
	}
	if q.createPasswordResetTokenStmt != nil {
		if cerr := q.createPasswordResetTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPasswordResetTokenStmt: %w", cerr)
		}
	}
	if q.createSessionStmt != nil {
		if cerr := q.createSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSessionStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteExpiredSessionsStmt != nil {
		if cerr := q.deleteExpiredSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExpiredSessionsStmt: %w", cerr)
		}
	}
	if q.deleteSessionStmt != nil {
		if cerr := q.deleteSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionStmt: %w", cerr)
		}
	}
	if q.deleteUserSessionsStmt != nil {
		if cerr := q.deleteUserSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserSessionsStmt: %w", cerr)
		}
	}
	if q.getSessionStmt != nil {
		if cerr := q.getSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.resetPasswordStmt != nil {
		if cerr := q.resetPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing resetPasswordStmt: %w", cerr)
		}
	}
	if q.updateSessionStmt != nil {
		if cerr := q.updateSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSessionStmt: %w", cerr)
		}
	}
	if q.verifyUserStmt != nil {
		if cerr := q.verifyUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	checkUserExistsStmt          *sql.Stmt
	createPasswordResetTokenStmt *sql.Stmt
	createSessionStmt            *sql.Stmt
	createUserStmt               *sql.Stmt
	deleteExpiredSessionsStmt    *sql.Stmt
	deleteSessionStmt            *sql.Stmt
	deleteUserSessionsStmt       *sql.Stmt
	getSessionStmt               *sql.Stmt
	getUserByEmailStmt           *sql.Stmt
	resetPasswordStmt            *sql.Stmt
	updateSessionStmt            *sql.Stmt
	verifyUserStmt               *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		checkUserExistsStmt:          q.checkUserExistsStmt,
		createPasswordResetTokenStmt: q.createPasswordResetTokenStmt,
		createSessionStmt:            q.createSessionStmt,
		createUserStmt:               q.createUserStmt,
		deleteExpiredSessionsStmt:    q.deleteExpiredSessionsStmt,
		deleteSessionStmt:            q.deleteSessionStmt,
		deleteUserSessionsStmt:       q.deleteUserSessionsStmt,
		getSessionStmt:               q.getSessionStmt,
		getUserByEmailStmt:           q.getUserByEmailStmt,
		resetPasswordStmt:            q.resetPasswordStmt,
		updateSessionStmt:            q.updateSessionStmt,
		verifyUserStmt:               q.verifyUserStmt,
	}
}
