# turner-defaultbackend

A trivial web service that returns _501 Not Implemented_ for all requests except of health checks, which get _200 OK_.

Configuration via the following env variables (all are optional):

* `PORT` (default: 8001) - the port to listen on
* `HEALTHCHECK`: (default: /healthz) - the URL that should return _200 OK_
* `PRODUCT` (default: "") - for inclusion in the body of the HEALTHCHECK response
* `ENVIRONMENT` (default: "") - for inclusion in the body of the HEALTHCHECK response
* `ENABLE_LOGGING` - set to `true` to log requests to stdout
