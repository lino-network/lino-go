# lino-go

## Examples

### Broadcast transfer transaction
```
api := api.NewLinoAPIFromConfig()
err := api.Broadcast.Transfer(1, "Lino", "receiver", "10000", "memo", privKeyHex)
```

### Get Validator
```
api := api.NewLinoAPIFromConfig()
validator, err := api.Query.GetValidator("Lino")
```

### More examples are under folder of /lino-go/example/...