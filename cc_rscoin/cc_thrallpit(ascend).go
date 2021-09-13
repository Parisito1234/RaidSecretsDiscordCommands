{{$pitKey := "RSCoinThrallPit"}}
{{$data := (dbGet $.User.ID $pitKey).Value}}
{{if eq ($data.Get "State") "true"}}
    {{$data.Set "State" "false"}}
    {{dbSet $.User.ID $pitKey $data}}
    {{dbSetExpire $.User.ID "RSCoinThrallPitCooldown" 1 1800}}
    {{sendMessage nil "You left the thrallpit and forfeit all reward"}}
    {{deleteMessage nil ($data.Get "X")}}
{{end}}
