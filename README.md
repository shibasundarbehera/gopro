# gopro

go mod init example.com/goAPI                 # Step 1: Initialize Go module
go clean -modcache                                    # Step 2: Wipe corrupt or stuck module cache
go get github.com/gin-gonic/gin@v1.10.1   # Step 3: Explicitly install Gin
go mod tidy                                                   # Step 4: Cleanup go.mod/go.sum properly
go run .                                                          # Step 5: Run your app!
