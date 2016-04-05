package main

import "fmt"
import "os"
import "io"
import "bufio"
import r "github.com/solos/requests"


func fetch(url string, proxy string) string {
    req := &r.Request{}
    fmt.Println(url)
    resp, _ := req.MakeRequest("GET", url, r.Timeout(10), r.Proxies(proxy))
    return resp.Content
}


func main() {
    channel := make(chan string)

    inputFile, inputError := os.Open("urls.txt")
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" + "Does the file exist?\n" + "Have you got acces to it?\n")
    }

    proxy := "http://183.207.228.122"

    defer inputFile.Close()
    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        if readerError != io.EOF {
            fmt.Printf("The input was: %s", inputString)
            go func() {
                s := fetch("http://www.baidu.com", proxy)
                channel <- s
                fmt.Println(s)
            }()
        }
    }

    //proxy := "http://183.207.228.122"
    //for i:=0; i < 1000; i++ {
    //}


    //for i := range channel {
    for i:=0; i<2;i++ {
        <- channel
        fmt.Println(i)
        //fmt.Println(j)
    }
}
