# Play Go Dice
This is game of dice using terminal.

## Deploy (local)
- `git clone`
- `cd PlayGoDice`
- `go run .`

## Request to play

Description:

- `Face` => `6`.
- `Quantity` => "Enter an integer:"

Tested with 999999, output example:

```
(...)
Dice: 6
---------
.:Table:.
---------
Face: 1 Qtd: 166289
Face: 2 Qtd: 166826
Face: 3 Qtd: 166650
Face: 4 Qtd: 166847
Face: 5 Qtd: 166229
Face: 6 Qtd: 167158
Done! It took 26.033195287 seconds!
```

# Test Unit
- `go test`
- `go test -v`
- `go test -cover`
