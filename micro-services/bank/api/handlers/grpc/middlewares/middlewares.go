package middlewares

import (
	"context"
	appCtx "gholi-fly-bank/pkg/context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// TransactionInterceptor adds a DB transaction to the context for specific handlers.
func TransactionInterceptor(db *gorm.DB, methodsRequiringTx map[string]bool) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Check if the current method requires a transaction
		if !methodsRequiringTx[info.FullMethod] {
			return handler(ctx, req) // Call handler without a transaction
		}

		// Begin a new transaction
		tx := db.Begin()
		if tx.Error != nil {
			log.Printf("failed to start transaction: %v", tx.Error)
			return nil, tx.Error
		}

		// Wrap the context with the DB transaction and set commit flag
		ctx = appCtx.NewAppContext(ctx, appCtx.WithDB(tx, true))

		// Call the handler
		resp, err := handler(ctx, req)

		// Check for errors and rollback if needed
		if err != nil {
			if rollbackErr := appCtx.Rollback(ctx); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
				return nil, status.Errorf(codes.Internal, "rollback failed: %v", rollbackErr)
			}
			return nil, err // Return the original handler error
		}

		// Commit the transaction
		if commitErr := appCtx.Commit(ctx); commitErr != nil {
			log.Printf("failed to commit transaction: %v", commitErr)
			return nil, status.Errorf(codes.Internal, "transaction commit failed: %v", commitErr)
		}

		return resp, nil
	}
}

// LoggerInterceptor logs the details of each gRPC request and response.
func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		startTime := time.Now()

		// Log the incoming request
		log.Printf("gRPC Request: %s | Payload: %+v", info.FullMethod, req)

		// Call the handler
		resp, err := handler(ctx, req)

		// Log the response and execution time
		duration := time.Since(startTime)
		if err != nil {
			log.Printf("gRPC Error: %s | Duration: %v | Error: %v", info.FullMethod, duration, err)
		} else {
			log.Printf("gRPC Response: %s | Duration: %v | Response: %+v", info.FullMethod, duration, resp)
		}

		return resp, err
	}
}
