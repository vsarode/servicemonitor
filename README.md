# servicemonitor
Service to monitor service APIS, along with deatails like response time and status.

**Steps to test**

``` go test ./... -v ```

**Steps to run**

```1) docker build --tag <tag_name> .```

```2)  docker run -p 8080:8080 -v <config_file_path>:$WORKDIR/config/local.json <tag_name> ```

visit http://localhost:8080/
