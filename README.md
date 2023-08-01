1. Remove any existing `go.mod` and `go.sum` files

2. Set your `GO111MODULE` to auto using

```bash
go env -w GO111MODULE=auto
```

3. `go mod init <module_name>`

4. import the necessary external packages in your go files

5. `go mod tidy`

6. the above command will install all the dependencies

> IMPORTANT: If you open a folder containing multiple separate GO projects in vs code then you will error that `unable to import the package`. That is because Go is getting confused looking at multiple go.mod files at the same time. To fix this issue, open a single Go project in a single vs code instance
