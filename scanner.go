package main

import (
    "fmt"
    "net"
    "time"
)

func scan(host string, c chan<- string) {
    conn, err := net.DialTimeout("tcp", host, 3*time.Second)
    if err != nil {
        c <- "" 
        return
    }
    conn.Close()
    c <- fmt.Sprintf("%s\topen\n", host)
}

func main() {
    threads :=  100 
    max_port := 3500 

    host := "localhost"

    chunks := max_port / threads

    for i := 0; i < chunks; i++ {
        ports := make([]int, threads)

        /* fill up ports slice */
        for j:= 0; j < threads; j++ {
            ports[j] = j+1+i*threads
        }

        c := make(chan string)
        n := 0

        /* start scanning with the specified number of threads */
        for j := range ports {
            n++
            h := fmt.Sprintf("%s:%d", host, ports[j])
            go scan(h, c)
        }

        /* collect results */
        for j := 0; j < n; j++ {
            result := <-c
            fmt.Print(result)
        }
    }
}
