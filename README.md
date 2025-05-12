Concurrent URL Status Checker
A command-line tool written in Go to check the HTTP status of URLs concurrently and generate a CSV report with the results. This tool allows you to specify a concurrency limit and handles errors robustly.

Build or run the project:
- go run main.go checker.go -f urls.txt -c 10

Example Usage
To run the project with a concurrency limit of 5:
go run main.go checker.go -f urls.txt -c 5

Features
Concurrent URL Checking: The tool uses Go goroutines to check multiple URLs at the same time.
Configurable Concurrency: Set the concurrency limit using the -c flag.
Error Handling: Properly handles malformed URLs or unreachable servers.
Progress Indicator: Displays real-time progress while checking URLs.
CSV Report: Outputs the results (including URL, HTTP status code, response time, and errors) to a report.csv file.
Summary Stats: At the end of the process, the tool prints summary stats such as total URLs, failed URLs, and average response time.

This will check the URLs listed in urls.txt with a maximum of 5 concurrent checks.
Input File Format
The input file (urls.txt) should contain one URL per line, for example:
https://google.com
https://example.com
https://nonexistent-url.xyz

Output Report Format
The tool will generate a report.csv file in the following format:
URL,Status,ResponseTime(ms),Error
https://google.com,200,143,
https://example.com,200,89,
https://nonexistent-url.xyz,N/A,N/A,dial tcp: no such host
Summary Stats (Printed to Console)

After completion, the following summary stats are displayed:
Summary:
Total URLs: 20
Failed URLs: 3
Average Response Time: 142.67 ms
Total Time Taken: 1.24s
Error Handling
If a URL is malformed or unreachable, the error is logged and reported in the report.csv file.
