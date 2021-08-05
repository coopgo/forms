# formsvc from coopgo/forms

`formsvc` is an example form service using [coopgo/forms](https://github.com/coopgo/forms) library.

## REST examples :

curl -d '{ "type": "Structured", "schema": { "value1": { "type": "Text" }, "value2": { type: "Parent", "child": { "value_child": { "type": "Boolean" } }  } } }' -X POST http://localhost:8080/forms/

## UI

This is intended to be a minmalistic web service example.

We implement a very basic example UI targetting the REST API using [Svelte](https://svelte.dev/) framework.
To run `formsvc` with this UI, you must add the -ui flag to the command line.

If you want something more efficient for form responses visualization, you could plug something like [NocoDB](https://nocodb.com/) to the database (if you backend is PostgreSQL)

