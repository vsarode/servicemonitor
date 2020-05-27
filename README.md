# servicemonitor
Service to monitor service APIS, along with deatails like response time and status.

**Steps to test**

``` go test ./... -v ```

**Steps to run**

```1) docker build --tag <tag_name> .```

```2) docker run -p 8080:8080 -d -e cron="5" <tag_name> ```

visit http://localhost:8080/
