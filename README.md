# Number manager

This is a simple API that holds value of numbers. It's great for keeping track of builds in CIs. It can set new values, increment or decrement them. The API is pretty simple and doesn't use any kind of authentication.

## Build

With docker:

    docker build -t creckx/numbermanager:latest .

Without docker:

    go get
    go test
    go run *.go
    # or
    go build -o numbermanager *.go

## API

### Registration

	POST /api/number/register
	-- no input --

Returns ID of a new number you need to access its value in other API endpoints.

### Set

	POST /api/number/:id
	{"number": 99}

Sets specific value to the number.

### Get

	GET /api/number/:id

Returns value of the number.

### Increment

	POST /api/number/:id/incr
	-- no input --

Increment the number by one.

### Decrement

	POST /api/number/:id/decr
	-- no input --

Decrement the number by one.