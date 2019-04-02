# redirect
> a simple tool built to redirect http requests in a 5MB docker image


## Usage

### As a binary
```
go get github.com/moqmar/redirect
TO=https://example.org PERMANENT=1 PREFIX=/home/ redirect
```

### With Docker
```
docker run -d -p 8080:80 -e "TO=https://example.org" -e "PERMANENT=1" -e "PREFIX=/home/" momar/redirect
```
This will redirect e.g. http://example.com/home/whatever.txt to https://example.com/whatever.txt


## Environment
Variable | Description
-------- | -----------
`TO` | Redirection target/prefix. Required.
`PERMANENT` | Use `301 Moved Permanently` instead of `302 Found` if this is set to something other than an empty string, `0`, `false` or `no`.
`PREFIX` | Prefix regular expression to remove from the path - a leading slash is required, a trailing slash is recommended in most cases. To completely ignore the path, set this to `.*` - the default behaviour is to keep the full path.
`HOST` & `PORT` | The hostname and port to listen on.


## Examples
Environment | Request | Target
`TO=https://example.org/hello PERMANENT=1 PREFIX=/world/` | http://example.com/world/test | 301 https://example.org/hello/test
`TO=https://example.org/hello PERMANENT=1 PREFIX=/world/` | http://example.com/whatever/test | 301 https://example.org/hello/whatever/test
`TO=https://example.org/hello PERMANENT=1 PREFIX=/[^/]*/` | http://example.com/whatever/test | 301 https://example.org/hello/test
`TO=https://example.org/hello PERMANENT=1 PREFIX=/world` | http://example.com/worldtest | 301 https://example.org/hello/test
`TO=https://example.org/hello PERMANENT=0` | http://example.com/whatever/test | 302 https://example.org/hello/whatever/test
`TO=https://example.org/hello PERMANENT=0 PREFIX=.*` | http://example.com/whatever/test | 302 https://example.org/hello
