# Drive into Golang

You will see the raw codes that I will write in Go. Create issue, gives me motivation.
[__Go Roadmap from roadmap.sh__](https://roadmap.sh/golang) 

### Repl of this Repo

1. [Hands dirty for understanding the basic](Basics/booking-app)
2. [Concept base deep drive](concept)
    - [Pointers](concept/pointers.go)
    - [Struct](concept/structs.go)
    - [Time](concept/time.go)
3. [Http server from scratch](httpserver/main.go)
4. [Be fluent in DSA & Algo by solving LeetCode problem](LeetCode)
    - [Binary Search](LeetCode/BinarySearch.go)
    - [Cyclic Sort](LeetCode/CyclicSort.go)
    - [Recursion](LeetCode/Recursion.go)


5. [Todo-API](/Todo-API)<br>

   #### Features

   - GetTodos
   - GetTodo
   - PostTodo
    - ToggleStatus

   #### How to run

    * Machine must have to installed [GO](https://go.dev/)
    * Go to Todo-API folder and run `go run main.go`
    * If everyting goes well, project is running on port `3000` as mentioned into `main.go` file.

   [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/8196637-d4c60aa9-b0b5-4640-90b3-e2e150a0ed6d?action=collection%2Ffork&collection-url=entityId%3D8196637-d4c60aa9-b0b5-4640-90b3-e2e150a0ed6d%26entityType%3Dcollection%26workspaceId%3D08445e69-71f6-45bd-bc3e-55df1a0de819)

   _Note: Build with `go version go1.19 linux/amd64`, Used [gin](https://pkg.go.dev/github.com/gin-gonic/gin@v1.8.1)_
   package for developing REST endpoints.