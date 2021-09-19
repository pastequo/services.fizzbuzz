# services.fizzbuzz

Run `make` for more information about the Makefile commands

## Dev Setup

 - Run `make tools` to install tools locally in the repo. Not needed for compilation. Tools are locally versioned using `gex`.
 - Run `make build.vendor build.local serve.local` to start server locally
 - Run `make build.vendor build.docker server.docker` to start the docker image in background
 - Run `make run.XXX` to request the server. Run targets, since they are not generic, are not part of the Makefile template. They are gathered into `custom.mk`. More parameters are exposed in the file to play with the service, eg `make run.fizzbuzz LIMIT=50`

## How to

 - Find information related to the Makefile: `make` or `make help`
 - Format code: `make check.fmt`
 - Format import: `make check.imports`
 - Check license usage: `make check.licenses`
 - Code linter : `make check.lint`
 - Execute TU: `make check.test`

## Notes
 - The clean architecture (by Uncle Bob) is not the best fit for this kind of service.
   However, depending on how the stat route is considered, it might be possible to see the algorithm parameters as entities.
   Add calling the algorithm is then equivalent to posting a new object parameter.
   It depends on the functional vision.
 - Related to the stat, at first sight, it seems like a metrics route. This is why I added prometheus.
   I added a stat route to justify the clean archiecture design and also because functionaly it fits more the usecase.
   But the metrics route fills the needs as well
 - Talking about this stat route, the number of call is persisted in memory, which has several drawbacks:
   - if the µservice is restarted, the numbers are reinitialised
   - if the µservice is replicated, the counts are not global
   Those drawbacks could be mitigated by using a database or the prometheus metrics ;)
 - I wrote a sample of deployment file for kustomize. I believe they belong to a different repo
 - No authN / authZ / tracing / rate limiting had been implemented for very different reasons
