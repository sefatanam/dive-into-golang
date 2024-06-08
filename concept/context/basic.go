package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunBasicContext() {
	// If request exced 400ms+ then request get automatically get cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)

	defer cancel()

	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/users/1",
		"https://jsonplaceholder.typicode.com/posts/1",
	}

	results := make(chan string)

	for _, url := range urls {
		go FetchApi(ctx, url, results)
	}

	for range urls {
		log.Println(<-results)
	}

}

func FetchApi(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		results <- fmt.Sprintf("error creating request, %s , %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("error creating request, %s , %s", url, err.Error())
		return
	}

	defer response.Body.Close()
	results <- fmt.Sprintf("Response from %s: %d", url, response.StatusCode)

}

/*
Article that I followed
https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea
*/
