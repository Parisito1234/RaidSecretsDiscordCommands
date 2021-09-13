{{dbSetExpire (toInt (index .Args 1)) "RSCoinThrallPitCooldown" 1 1}}
{{$pitKey := "RSCoinThrallPit"}}
{{$data := (dbGet $.User.ID $pitKey).Value}}
{{$data.Set "State" "false"}}
{{dbSet $.User.ID $pitKey $data}}
