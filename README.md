# grpc-test

gRPC test service

problem repro for protovalidate  RegEx `\S` failure

Using matches with a RegEx `\S` not whitespace fails with a compile expression failure while using the semantic equivalent `[^\t\n\f\r ]` works correctly.

According to the RegEx documentation CEL, this  should be supported as part of the supported RE2 syntax (https://github.com/google/re2/wiki/Syntax)

Using `\S` failure

```
    (buf.validate.field) = {
        required: true,
        cel: {
            id: "hello_request_name"
            message: "cannot contain any spaces or other whitespace characters"
            expression: "this.matches('^[^\\S]+$')"
        }
        string: {
            max_len: 256
        }
    }
```

Failure output

```
ERROR:
  Code: Unknown
  Message: compilation error: failed to compile expression hello_request_name: ERROR: <input>:1:14: Syntax error: token recognition error at: ''^[^\S'
 | this.matches('^[^\S]+$')
 | .............^
ERROR: <input>:1:20: Syntax error: mismatched input ']' expecting {'[', '{', '(', ')', '.', '-', '!', 'true', 'false', 'null', NUM_FLOAT, NUM_INT, NUM_UINT, STRING, BYTES, IDENTIFIER}
 | this.matches('^[^\S]+$')
 | ...................^
ERROR: <input>:1:22: Syntax error: token recognition error at: '$'
 | this.matches('^[^\S]+$')
 | .....................^
ERROR: <input>:1:23: Syntax error: token recognition error at: '')'
 | this.matches('^[^\S]+$')
 | ......................^
ERROR: <input>:1:25: Syntax error: mismatched input '<EOF>' expecting {'[', '{', '(', '.', '-', '!', 'true', 'false', 'null', NUM_FLOAT, NUM_INT, NUM_UINT, STRING, BYTES, IDENTIFIER}
 | this.matches('^[^\S]+$')
 | ........................^
make: *** [run-client-failed] Error 66
```

Semantic equivalent:

```
    (buf.validate.field) = {
        required: true,
        cel: {
            id: "hello_request_name"
            message: "cannot contain any spaces or other whitespace characters"
            expression: "this.matches('^[^\\t\\n\\f\\r ]+$')"
        }
        string: {
            max_len: 256
        }
    }
```

Expected output

```
ERROR:
  Code: Unknown
  Message: validation error:
 - name: cannot contain any spaces or other whitespace characters [hello_request_name]
make: *** [run-client-success] Error 66
```