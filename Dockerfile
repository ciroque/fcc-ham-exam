############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git tree
WORKDIR $GOPATH/src/fcc-ham-exam
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN go build -o /go/bin/fcc-ham-exam

RUN tree 
COPY ./data/sources/*.json /go/bin/fcc-ham-exam/data/sources/

############################
# STEP 2 build a small image
############################
FROM alpine:3.10

# Copy our static executable.
COPY --from=builder /go/bin/fcc-ham-exam /go/bin/fcc-ham-exam
COPY --from=builder /go/bin/fcc-ham-exam/data/source/*.json /go/bin/fcc-ham-exam/data/sources/

# Run the hello binary.
ENTRYPOINT ["/go/bin/fcc-ham-exam"]
