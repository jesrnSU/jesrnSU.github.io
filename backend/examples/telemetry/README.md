# Original telemetry prototype

The previous `backend/main.go` is preserved here for reference, with formatting
and a file comment added. It is separate from the new API skeleton.

To run it intentionally, start the existing development database and run from `backend/`:

```sh
go run ./examples/telemetry
```

It uses port 8080, so stop the new API first. It retains the original hardcoded
database credentials, request-time table creation, and telemetry endpoint.
It writes to the database and is a local learning example, not deployment code.
