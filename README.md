# go-keep-alive-http

This is an example on how HTTP keep-alive work.

Example includes Golang server and client to reproduce the following network errors:
– EOF
– http: server closed idle connection
– read tcp: read: connection reset by peer

Also you can see http.Post and http.Post behaving differently due to Post not having idempotency property but Get having it.

Article with explanation: https://techgeorgii.com/go-golang-fix-network-issues-connection-reset-by-peer-http-server-closed-idle-connection-eof/



