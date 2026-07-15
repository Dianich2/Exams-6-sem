// тут короче в браузере просто несколько вкладок нужно быстренько открыть
// и тогда увидим параллельную обработку
// то есть на каждый запрос создается свой handler
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := time.Now().UnixNano()
	fmt.Printf("%d:start\n", id)
	time.Sleep(5 * time.Second)
	fmt.Printf("%d:end\n", id)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Server start on :3000\n")
	http.ListenAndServe(":3000", nil)
}
