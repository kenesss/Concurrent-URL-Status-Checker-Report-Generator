package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"
)

type Result struct {
    URL          string
    Status       string
    ResponseTime string
    Error        string
}

func readURLs(file *os.File) ([]string, error) {
    var urls []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        url := scanner.Text()
        if url != "" {
            urls = append(urls, url)
        }
    }
    return urls, scanner.Err()
}

func checkURL(url string) Result {
    start := time.Now()
    resp, err := http.Get(url)
    duration := time.Since(start).Milliseconds()

    if err != nil {
        return Result{URL: url, Status: "N/A", ResponseTime: "N/A", Error: err.Error()}
    }
    defer resp.Body.Close()

    return Result{URL: url, Status: fmt.Sprintf("%d", resp.StatusCode), ResponseTime: fmt.Sprintf("%d", duration), Error: ""}
}

func checkURLsWithProgress(urls []string, concurrency int) []Result {
    var wg sync.WaitGroup
    ch := make(chan Result, len(urls))
    sem := make(chan struct{}, concurrency)

    completed := 0
    total := len(urls)
    mu := sync.Mutex{}

    for _, url := range urls {
        wg.Add(1)

        go func(url string) {
            defer wg.Done()
            sem <- struct{}{}

            result := checkURL(url)
            ch <- result

            mu.Lock()
            completed++
            fmt.Printf("\rProgress: %d/%d URLs checked", completed, total)
            mu.Unlock()

            <-sem
        }(url)
    }

    wg.Wait()
    close(ch)

    var results []Result
    for result := range ch {
        results = append(results, result)
    }
    return results
}

func generateReport(results []Result, outputPath string) error {
    file, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"URL", "Status", "ResponseTime(ms)", "Error"})

    for _, result := range results {
        writer.Write([]string{result.URL, result.Status, result.ResponseTime, result.Error})
    }

    return nil
}

func calculateStats(results []Result) (total int, failed int, avgTime float64) {
    total = len(results)
    var totalTime int64
    for _, result := range results {
        if result.Error != "" {
            failed++
        } else if result.ResponseTime != "N/A" {
            var rt int64
            fmt.Sscanf(result.ResponseTime, "%d", &rt)
            totalTime += rt
        }
    }

    if total-failed > 0 {
        avgTime = float64(totalTime) / float64(total-failed)
    }
    return
}
