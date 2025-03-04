package ch12
import "fmt"
func countTo(max int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < max; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

func Master3() {
    for i := range countTo(10) {
        fmt.Println(i)
    }
}