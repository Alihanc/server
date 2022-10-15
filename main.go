package main

import (
	"context"
	"os"
)

func main() {
	var errVal string
	if len(os.Args) > 1 {
		errVal = os.Args[1]
	}

	ss := slowServer()
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()

	ctx := context.Background()
	callBoth(ctx, errVal, ss.URL, fs.URL)
}
