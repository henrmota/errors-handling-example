# Errors example
Just an example of how I build errors with type and context just exporting an error type.

Handling errors is hard, because there is some contextual information we want to save. 

Besides that, errors can be created at every layer of an application. 
Given that we want to handle this exporting the minimum types possible.

In this example I wrapped the library from https://github.com/pkg/errors, extending the functionalities.

Creating a typed errors is as easy:

```GO
  errors.BadRequest.New("error parsing the input information")
```

You can create an untyped error as easy as:

```GO
  errors.New("an untyped error")
```

Adding a new context to an existing error:

```GO
  errors.AddContext(err, "field", "message")
```

In the top layer when you decide to log or return a web response:

```GO
  errors.GetType(err) == errors.BadRequest // true
  errors.GetContext(err) // map[string]string{"field": "field", "message": "message"}
```

To add new error type is just adding a new constant to errors

```GO
const (
	NoType = ErrorType(iota)
	BadRequest
	NotFound
  //ADD here
)
```
