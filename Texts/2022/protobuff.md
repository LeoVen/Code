# Protocol Buffers

> Protocol buffers provide a language-neutral, platform-neutral, extensible mechanism for serializing structured data in a forward-compatible and backward-compatible way. It’s like JSON, except it's smaller and faster, and it generates native language bindings.

## Sources

* [Overview](https://developers.google.com/protocol-buffers/docs/overview)
* [proto3 Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)

## Benefits

* Compact data storage
* Fast parsing
* Availability in many programming languages
* Optimized functionality through automatically-generated classes

## Workflow

![Protobuff Diagram](https://developers.google.com/protocol-buffers/docs/images/protocol-buffers-concepts.png)

## Syntax (proto3)

A field of a `message` is made up of the following components:

> You can also create your own composite data types by defining messages that are, themselves, data types that you can assign to a field.

* A field rule (optional):
  * `singular`: a `message` can have zero or one of this field, and no more than that
  * `repeated`: a `message` this field can repeat any number of times, including zero
* A field type (required)
* A field name (required)
* A field number (required)

```proto
/* SearchRequest represents a search query, with pagination options to
 * indicate which results to include in the response. */
message SearchRequest {
  // The query to be run at the database
  singular string query = 1;
  // Pagination's page number
  singular int32 page_number = 2;
  // Number of results to return per page
  singular int32 result_per_page = 3;
}
```

## Reserved Fields

Reserve field numbers or names for guaranteeing that future updates in the `message` do not break serialization and deserialization. If a field is removed or updated, reserving can assure that the number or name isn't reused, which can cause issues.

```proto
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
  ...
}
```

## Default Values

* `string`: empty string
* `bytes`: empty bytes
* `bool`: false
* Numeric: zero
* Enums: First defined enum value
* Message Fields: the field is not set. The exact value is language-dependent
* Repeated field: empty list
