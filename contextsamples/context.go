package contextsamples

import (
	"context"
	"fmt"
)

type userKey struct{}

func contextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey{}).(string)
	return user, ok
}

func ContextWithData() {
	userContext := contextWithUser(context.Background(), "this is user")
	if userData, ok := userFromContext(userContext); ok {
		fmt.Println("User Data:", userData)
	}
}
