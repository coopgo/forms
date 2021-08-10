# coopgo/forms

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/coopgo/forms.svg)](https://github.com/coopgo/forms)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/coopgo/forms)

`coopgo/forms` is a web service written in [Golang](https://golang.org/) to handle forms data using différent data formats, protocols and backends

Communication protocols (transports) :

- [X] HTTP API (REST) 
- [ ] gRPC

Multiple transports can be run at the same time.

Different data backends / storages can be set :

- [X] Memory (temporary storage, for development/testing purpose)
- [X] [PostgreSQL](https://www.postgresql.org/) database (might also work with [CockroachDB](https://github.com/cockroachdb/cockroach) but not tested yet as we use the [pgx](https://github.com/jackc/pgx) driver compatible with both PostgreSQL and CockroachDB). PostgreSQL backend relies heavily on database tables/schemas generated on the fly (each table stores data from one form, even for `unstructured` forms, without defined schemas), so that we can easily plug another tool to visualize data like [Apache Superset](https://github.com/apache/superset), [Metabase](https://github.com/metabase/metabase) or [NocoDB](https://www.nocodb.com/).
- [X] [Kantree](https://kantree.io) (one way) (TODO : add some documentation about this)
- [ ] Email with SMTP (one way)
- [ ] Webhooks (one way)

"One way" backends do not store the form states. They could be better considered as "connectors" to other services, but are designed as the other backends in the code (they implement the same interface `Backend`).

At least one storage has to be set at the service level globally. These global storages cannot be "one way". Supported "global" storages are at the moment : Memory (temporary storage : only for development) and PostgreSQL

Multiple additional storages can be set for each form individually.

Our priority is to support memory (for easy development purpose), PostgreSQL, [Kantree](https://kantree.io) and Email (because that's our internal needs at [COOPGO](https://coopgo.fr)), but feel free to contribute something else (other databases, connectors to other services like chat apps, webhooks, etc...) if you want it first.

We have different types of forms :

- [X] Structured forms : forms with a defined schema. Every response must match the given schema. Structured forms schemas are defined using the [JSON Type Definition standard](https://jsontypedef.com/).
- [X] Unstructured forms : forms accepting any input

## Project Status

This project is still in development. It should be kind of functional using the PostgreSQL backend for simple use cases. Efforts still need to be made to support more complex forms with PostgreSQL and to implement other backends.

## Contributing

We welcome any contributions following theses guidelines :
- Write simple, clear and maintainable code and avoid technical debt. 
- Leave the code cleaner than when you started.
- Refactoring existing code for better performance, better readability or better testing wins over creating a new feature.

If you want to contribute, you can fork the repository and create a pull request.

## Bug report

For reporting a bug, you can open an issue using the **Bug Report** template. Try to write a bug report that is easy to understand and explain how to reproduce the bug. 
Do not duplicate an existing issue and keep each issue specific to an individual bug.

## License

`coopgo/forms` is under the Apache 2.0 license. Please refer to the [LICENSE](LICENSE) file for details.
