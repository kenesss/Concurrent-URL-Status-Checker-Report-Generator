Concurrent URL Status Checker
A command-line tool written in Go to check the HTTP status of URLs concurrently and generate a CSV report with the results. This tool allows you to specify a concurrency limit and handles errors robustly.

Build or run the project:
- go run main.go checker.go -f urls.txt -c 10


Features
Concurrent URL Checking: The tool uses Go goroutines to check multiple URLs at the same time.
Configurable Concurrency: Set the concurrency limit using the -c flag.
Error Handling: Properly handles malformed URLs or unreachable servers.
Progress Indicator: Displays real-time progress while checking URLs.
CSV Report: Outputs the results (including URL, HTTP status code, response time, and errors) to a report.csv file.
Summary Stats: At the end of the process, the tool prints summary stats such as total URLs, failed URLs, and average response time.

Installation
Clone the repository:
git clone https://github.com/yourusername/url-status-checker.git

Change to the project directory:
cd url-status-checker

Command-line Arguments
-f: Specify the path to the text file containing URLs (default: urls.txt).
-c: Set the concurrency limit for checking URLs (default: 10).

Example Usage
To run the project with a concurrency limit of 5:
go run main.go checker.go -f urls.txt -c 5

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

The program does not crash due to failed requests.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Customization (Optional)
You can modify the concurrency limit (-c) to adjust the number of concurrent requests.

Add more sophisticated error handling as needed for specific use cases.

Extend the program to handle more detailed statistics, such as response headers or request retries.

Contributions
Feel free to fork the repository, create issues, or submit pull requests for any improvements or bug fixes.
