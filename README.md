# go-decouple

Package go-decouple is inspired by the [python-decouple][] package. It
provides a layer above [gotdotenv][] that handles defaults and type
conversion.

[godotenv]: https://github.com/joho/godotenv
[python-decouple]: https://github.com/henriquebastos/python-decouple

For example, if you want to read an integer value from an
environment variable, you can call:

```
decouple.Load()
value, exists = decouple.GetInt("MY_INT_VAR", 0)
```

If `MY_INT_VAR` exists, `value` will be set to the integer value and
`exists` will be `true`. If `MY_INT_VAR` does not exist, `value` will
be set to 0 and `exists` will be `false`. If `MY_INT_VAR` exists but
cannot be converted into the requested type, `value` will be set to
0 and `exists` will be `false`.
