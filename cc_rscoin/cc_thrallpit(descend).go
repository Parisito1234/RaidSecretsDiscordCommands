{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$key := "RSCoinBalance" }}

{{$pitKey := "RSCoinThrallPit"}}
{{$cooldownKey := "RSCoinThrallPitCooldown"}}

{{$cooldown := dbGet $.User.ID $cooldownKey}}
{{if eq (toInt $cooldown.Value) 1}}
    {{$time := $cooldown.ExpiresAt}}
    {{sendMessage nil (joinStr "" "You are exhausted from your previous encounters in the thrall pit, and will be exhausted for the next " (humanizeDurationSeconds (toDuration ($time.Sub currentTime))))}}
{{else}}

{{/*
    Dict:
    State: true = currently active, false = not active
    X: message ID if active game
    Attack: character stats
    Defence: same as above
    Luck: same as above
    HP: current player hitpoints
    Coins: current active coin pool
    Kills: current hive killed
    hiveSlain: historical stats
    CoinsEarnt: historical stats
    XP: xp
*/}}

{{$data := (dbGet $.User.ID $pitKey).Value}}
{{$State := "true"}}

{{if eq (printf "%T" $data) "*templates.SDict"}}
    	{{$State = $data.Get "State"}}
{{else}} {{/*NEW PLAYER*/}}
        {{$x := sendMessageRetID nil "New Player detected, use -descend to begin your adventure"}}
        {{$data = sdict
        "State" "false"
        "Luck" 0
        "hiveSlain" 0
        "CoinsEarnt" 0
        "XP" 1
        }}
        {{dbSet $.User.ID $pitKey $data}}
        {{$State := "false"}}
{{end}}

{{if eq $State "false"}}
    {{$data := (dbGet $.User.ID $pitKey).Value}}
    {{$XP := $data.Get "XP"}}
    {{$level := (toInt (pow $XP 0.4))}}

    {{$embed := cembed
        "title" (joinStr "" "__" $.User.Username "__ wishes to enter the thrallpit")
        "fields" (cslice
            (sdict "name" "XP:" "value" (toString $XP) "inline" true)
            (sdict "name" "Level:" "value" (toString $level) "inline" true)
            (sdict "name" "Options:" "value" "☑ to enter the Thrallpit™ or venture deeper in.\n❎ to walk away or to leave with the loot you've obtained"))
        }}

    {{$x := sendMessageRetID nil $embed}}
    {{addMessageReactions nil $x ":ballot_box_with_check:" ":negative_squared_cross_mark:"}}
    {{$data.Set "State" "true"}}
    {{$data.Set "X" $x}}
    {{$data.Set "Start" 1}}
    {{dbSet $.User.ID $pitKey $data}}




{{else}}
    {{sendMessage nil "You are currently in the thrall pit, quit tryna break my code"}}

{{end}}


{{end}}
