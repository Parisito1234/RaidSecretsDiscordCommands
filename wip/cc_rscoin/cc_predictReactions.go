{{ $balanceKey := "RSCoinBalance" }}
{{ $gameKey := "RSCasinoPrediction" }}
{{ $e := "<:RSStonkCoin:869340420692394095>" }}

{{ $x := .Reaction.MessageID }}
{{ $gameState := (dbGet $x $gameKey).Value }}
{{ $user := .User }}

{{ if eq (printf "%T" $gameState) "*templates.SDict" }}
{{ end }}