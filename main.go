package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    filePath := flag.String("f", "urls.txt", "Path to the input file containing URLs")
    concurrency := flag.Int("c", 10, "Number of concurrent URL checks")
    flag.Parse()

    file, err := os.Open(*filePath)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer file.Close()

    urls, err := readURLs(file)
    if err != nil {
        log.Fatalf("Error reading URLs: %v", err)
    }

    fmt.Printf("Checking %d URLs with concurrency %d...\n", len(urls), *concurrency)
    start := time.Now()
    results := checkURLsWithProgress(urls, *concurrency)
    duration := time.Since(start)

    err = generateReport(results, "report.csv")
    if err != nil {
        log.Fatalf("Error generating report: %v", err)
    }

    total, failed, avgTime := calculateStats(results)
    fmt.Println("\nReport generated: report.csv")
    fmt.Printf("\nSummary:\n")
    fmt.Printf("Total URLs: %d\n", total)
    fmt.Printf("Failed URLs: %d\n", failed)
    fmt.Printf("Average Response Time: %.2f ms\n", avgTime)
    fmt.Printf("Total Time Taken: %s\n", duration)
}
