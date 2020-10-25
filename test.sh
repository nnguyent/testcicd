#!/bin/bash

# go test -v
# go test -v -run TestSum
# go test -v -cover

#---------------
## go help test
#	# -c
#	#     Compile the test binary to pkg.test but do not run it
#	#     (where pkg is the last element of the package's import path).
#	#     The file name can be changed with the -o flag.
## go test -c -coverpkg="github.com/nnguyent/testcicd"
# go test -c -coverpkg="./..."
# ./testcicd.test -test.coverprofile=coverage.out # -test.run "^TestRunMain$"
# go tool cover -html=coverage.out -o coverage.html

#---------------
rm -f coverage*
# go test -v -coverprofile=coverage.out

# Run unittest of all packages 
# go test -race -covermode atomic -coverprofile=coverage.out ./...

# Run unittest of package SilentiumIO
go test -race -covermode atomic -coverprofile=coverage.out ./SilentiumIO
go tool cover -html=coverage.out -o coverage.html

#---------------
# Run unittest of package fmt. And list coverage percent of functions in package fmt.
# go test -race -covermode=count -coverprofile=count.out fmt
# go tool cover -func=count.out