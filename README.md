# go-pubsub-debug
Google Cloud Pub/Sub debug tool written in Go.

This is a simple binary that subscribes to a given topic in Google Cloud
Pub/Sub and prints the message it receives.

## Running/Configuration
Make sure the `GCLOUD_PROJECT` environment variable is set (and that you
have application default credentials set up).

``` bash
$ go build .
$ ./pubsub-debug -topic=my-test-topic -subscription-name=my-test-sub
```
