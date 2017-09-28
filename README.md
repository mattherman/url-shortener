A URL shortener service written in Go. Uses Redis for persistence. Written as a (mini) hackathon project in order to play with docker-compose and AWS.

Can be run locally using the following command:

```docker-compose up```

This will start the web client on port 8080. Use the web client to create shortened URLs. You can also POST directly to the `/create/{alias}` endpoint. If you do not include an alias, an auto-generated value will be used.
