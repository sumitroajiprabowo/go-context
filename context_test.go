package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	ctx := context.Background()                 // create a new context with default values
	fmt.Println("context.Background() = ", ctx) // context.Background() =  context.Context{}

	todo := context.TODO()                 // create a new context with default values and a TODO key
	fmt.Println("context.TODO() = ", todo) // context.TODO() =  context.Context{context.TODO}

}

func TestContextWithValue(t *testing.T) {
	type name string
	k := name("language")
	contextA := context.WithValue(context.Background(), k, "Programming")
	contextB := context.WithValue(contextA, k, "Python")
	contextC := context.WithValue(contextA, k, "JavaScript")

	contextD := context.WithValue(contextB, k, "Django")
	contextE := context.WithValue(contextB, k, "Flask")
	contextF := context.WithValue(contextC, k, "Vue")

	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextA.Value(k))
	fmt.Println(contextB.Value(k))
	fmt.Println(contextC.Value(k))
	fmt.Println(contextD.Value(k))
	fmt.Println(contextE.Value(k))
	fmt.Println(contextF.Value(k))
}

func TestExampleContextWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
