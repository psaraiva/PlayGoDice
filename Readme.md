# Play Go Dice
This is game of dice using terminal.

## Deploy (local)
- `git clone`
- `cd PlaySlimDice`
- `go run .`

## Request to play

Description:

- `Face` => `6`.
- `Quantity` => "Enter an integer:"

Tested with 999,999, output example:

```
(...)
[999999] Dice: 6
---------
.:Table:.
---------
Index: 1 Qtd: 167068
Index: 2 Qtd: 166914
Index: 3 Qtd: 166789
Index: 4 Qtd: 166161
Index: 5 Qtd: 166282
Index: 6 Qtd: 166785
```
# Test Unit
- `go test`
- `go test -v`
- `go test -cover`
