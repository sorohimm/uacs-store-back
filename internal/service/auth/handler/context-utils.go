package handler

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

var ErrNoToken = errors.New("no token in context")

// Type for context keys
type contextKey int

const (
	// our unique key used for storing the request in the context
	requestContextKey contextKey = 0
)

// SetAccessTokenInContext sets the user's session into the context.  This has the effect of logging the user
// in as that userId.  The grpc json gateway will set the UID in the user's session in this case
func SetAccessTokenInContext(ctx context.Context, token string) error {
	// create a header that the gateway will watch for
	header := metadata.Pairs("cookie-access-token", token)
	// send the header back to the gateway
	return grpc.SetHeader(ctx, header)
}

func SetRefreshTokenInContext(ctx context.Context, token string) error {
	// create a header that the gateway will watch for
	header := metadata.Pairs("cookie-refresh-token", token)
	// send the header back to the gateway
	return grpc.SetHeader(ctx, header)
}

// SetDeleteSessionFlagInContext sets a flag telling the gateway to delete the session, if any
func SetDeleteSessionFlagInContext(ctx context.Context) error {
	// create a header that the gateway will watch for
	header := metadata.Pairs("session-delete", "true")
	// send the header back to the gateway
	return grpc.SetHeader(ctx, header)
}

// get the first metadata value with the given name
func firstMetadataWithName(md metadata.MD, name string) string {
	values := md.Get(name)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetTokenFromContext returns the userId that has been stored in Context, if available.
// This will return 0 if the user has not logged in.  If there is an error attempting to return
// the userId it will be returned.  It's valid for this function to return 0 with no
// error, which indicates the user has not logged in.
func GetTokenFromContext(ctx context.Context) (string, error) {
	// retrieve incoming metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// get the first (and presumably only) user ID from the request metadata
		token := firstMetadataWithName(md, "session-token")
		if token != "" {
			return token, nil
		}
	}
	return "", ErrNoToken
}

// pull the request from context (set in middleware above)
func getRequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(requestContextKey).(*http.Request)
}
