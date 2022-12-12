# go-keep-alive-http

This is an example on how HTTP keep-alive work. You can see http.Post and http.Post behaving differently due to Post not having idempotency property but Get having it.

Example includes Golang server and client to reproduce the following network errors:

- EOF
- http: server closed idle connection
- read tcp: read: connection reset by peer

To fix these errors go to article: https://techgeorgii.com/go-golang-fix-network-issues-connection-reset-by-peer-http-server-closed-idle-connection-eof/






