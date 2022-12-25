package main

import (
	"fmt"
	"functions/visitor"
)

func main() {
	vi := visitor.New()
	vi.Join()
	vi.Join()
	vi.Join()
	vi.Join()
	vi.BlukActivity(string(visitor.LeftActivity))(3)
	vi.Join()
	vi.BlukActivity(string(visitor.JoinActivity))(10)
	fmt.Println(vi.ActiveVisitors())
}
