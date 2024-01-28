# Log Grepper Service

## Local Setup Steps
1. Clone the repo
2. Run go mod tidy
3. Run go run main.go

## Locally running S3
1. Install localstack from https://docs.localstack.cloud/getting-started/installation/
2. Install aws cli
3. cd to log-grepper
4. run docker-compose up

## Creating S3 bucket
1. Install aws cli
2. Run aws --endpoint-url=http://localhost:4566 s3 mb s3://2022-01-01

## Adding files to S3
1. Install aws cli
2. Create a local file (touch 01.txt, vim 01.txt, etc) 
3. Run aws --endpoint-url=http://localhost:4566 s3 cp 01.txt s3://2022-01-01

## Running the service
1. Run go run main.go
2. Run curl -X GET curl -X GET http://localhost:8080/search\?from\='2022-01-01'\&to\='2022-01-01'\&search_keyword\='error'

![Screenshot 2024-01-28 at 9.47.48 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2FScreenshot%202024-01-28%20at%209.47.48%20PM.png)

## Running the tests
1. Run go test ./...

## Design of project
The service is currently designed more towards low latency rather than low memory. Considering memory is cheap and also a major
assumption that logfiles can fit in memory, the service spawns multiple workers to read and process the files.

## Future improvements
The service can be fine tuned by add knobs around how big our files are and how much memory we want to give to service. We can call
s3's GetObject API with Range header to read only a part of the file and process it. This will help us in reducing the memory footprint.
We can also add a cache layer to cache the results of the search. This will help us in reducing the number of calls to s3 as mostly 
historical files are modified.
This willl however comes with a tradeoff for the latency.