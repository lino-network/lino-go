# lino-go

## Examples
### Broadcast transfer transaction
```
transferTx := model.TransferToAddressMsg{
  Sender:       "Lino",
  ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
  Amount:       "8888888",
}
broadcast.BroadcastTransaction(transferTx, privKeyBytes)
  
```

### Get Validators
```
res, _ := query.GetValidator("Lino")
output, _ = json.MarshalIndent(res, "", "  ")
fmt.Println(string(output))
  
```
