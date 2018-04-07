gRPC-Examples
==============

Heartbeat
------------

- Build .proto

    ```
    $ cd heartbeat/proto
    $ make build
    ```

- Build and run the server

  ```
  $ cd server
  $ make build && ./heartbeat
  ```

- Run the client

  ```
  $ cd client
  $ go run main.go -s elastic
  ```