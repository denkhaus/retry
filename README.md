# retry - retryable command execution made easy

## install

```sh {"id":"01J354RVXAQS18SM5RXX8HZWGZ"}
go install github.com/denkhaus/retry/cmd/retry@latest
```

## usage

```sh {"id":"01J352QPBBQYH53ZJ0TE9S1XXA"}
Usage of retry:
  -c string
        command to execute (env RETRY_COMMAND)
  -e    exponential backoff: enable (env RETRY_EXPONENTIAL_ENABLED)
  -ei duration
        exponential backoff: initial interval between retries (env RETRY_EXPONENTIAL_INITIAL_INTERVAL) (default 1s)
  -ej float
        exponential backoff: add jitter to the retry interval like Interval = RetryInterval * (1 ± RandomizationFactor) (env RETRY_EXPONENTIAL_RANDOMIZATION_FACTOR)
  -em float
        exponential backoff: retry duration multiplier (env RETRY_EXPONENTIAL_MULTIPLIER) (default 2)
  -et duration
        exponential backoff: maximum total time for retries (env RETRY_EXPONENTIAL_MAX_ELAPSED_TIME) (default 1h0m0s)
  -ex duration
        exponential backoff: maximum interval between retries (env RETRY_EXPONENTIAL_MAX_INTERVAL) (default 10s)
  -i duration
        constant backoff: interval between retries (env RETRY_INTERVAL) (default 1s)
  -l string
        log level - debug, info, warn, error (env RETRY_LOG_LEVEL) (default "info")
  -m uint
        constant/exponential backoff: max retries (env RETRY_RETRIES_MAX) (default 5)
  -v    show version and exit
```

## examples

```sh {"id":"01J35CJPYXA0H0MDVMD5A1VM21"}
retry -e -c "echo retry; exit 1" #with exponential growing interval
retry -i 5s -c "echo retry; exit 1" #with constant interval of 5 seconds
```