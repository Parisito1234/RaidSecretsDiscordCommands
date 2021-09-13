{{/*THRALLPIT REACT*/}}
{{$e := "<:RSStonkCoin:869340420692394095>"}}
{{$key := "RSCoinBalance" }}
{{$pitKey := "RSCoinThrallPit"}}
{{$cooldownKey := "RSCoinThrallPitCooldown"}}
{{$HP := 0}}
{{$Attack := 0}}
{{$XP := 0}}
{{$Coins := 0}}
{{$Kills := 0}}
{{$hiveSlain := 0}}
{{$data := (dbGet $.User.ID $pitKey).Value}}
{{$hiveDead := false}}


{{$hiveType := cslice "Thrall" "Stasis Thrall" "Buff Thrall" "Thrall with a Gun" "Thrall with a Wizard hat" "Treasure Chest" "Mega-Thrall with Laser Eyes" "Ascendant Thrall" "Thrall with a Hive Ghost" "Corrupted Sweeperbot Husk"}}
{{$hiveAtk := cslice 5 10 15 15 20 0 25 30 30 40}}
{{$hiveHP := cslice 15 25 35 45 50 1 75 80 100 200}}
{{$hiveValue := cslice 1 2 3 3 5 69 20 20 20 100}}
{{$hiveImage := cslice
"https://i.imgur.com/Co0tUnD.png"
"https://i.imgur.com/1Bm0boC.png"
"https://i.imgur.com/XWis9xb.png"
"https://i.imgur.com/kEkTNTA.png"
"https://i.imgur.com/I3CsT3u.png"
"https://i.imgur.com/pbDnw7V.png"
"https://i.imgur.com/6k1aSbz.png"
"https://i.imgur.com/lUXEiB3.png"
"https://i.imgur.com/042khm2.png"
"https://i.imgur.com/I3YQNYc.png"
}}

{{$H_Type := 0}}{{$H_Atk := 0}}{{$H_HP := 0}}{{$H_Value := 0}}{{$H_Image := 0}}

{{if eq (printf "%T" $data) "*templates.SDict"}}
	{{$x := $data.Get "X"}}
	{{if ne (toInt $x) 0}}
	{{if eq $x $.Reaction.MessageID}}
	    {{$emoji := $.Reaction.Emoji.Name}}
	    {{$start := $data.Get "Start"}}
	    {{if eq $emoji "☑"}}
            {{if eq $start 1}}
                {{/*NEW GAME*/}}
                {{deleteAllMessageReactions nil $x}}
                {{$data.Set "Start" 0}}
                {{$XP = $data.Get "XP"}}

                {{$level := (toInt (pow $XP 0.4))}}
                {{if gt $level 100}}
                    {{$level = 100}}
                {{end}}

                {{$HP = (add 100 (mult $level 2))}}
                {{$Attack = (add 10 $level)}}
                {{$Coins = 0}}
                {{$Kills = 0}}

            {{else}}
                {{/*CONTINUE*/}}
                {{deleteAllMessageReactions nil $x}}
                {{$HP = $data.Get "HP"}}
                {{$Attack = $data.Get "Attack"}}
                {{$XP = $data.Get "XP"}}
                {{$Coins = $data.Get "Coins"}}
                {{$Kills = $data.Get "Kills"}}
                {{$hiveSlain = $data.Get "hiveSlain"}}
            {{end}}

            {{$Hi := (mult 10 (add $Kills 1))}}
            {{if ge $Hi 100}}{{$Hi = (add 100 (mult 3 (sub $Kills 9)))}}{{end}}
            {{$Lo := (mult 5 (sub $Kills 1))}}
            {{if lt $Lo 0}}{{$Lo = 0}}{{end}}
            {{if gt $Kills 24}}{{$Lo = 145}}{{$Hi = 146}}{{end}}
            {{$H := randInt $Lo $Hi}}

            {{$T := 0}}
            {{if le $H 10}}{{$T = 0}}
            {{else if le $H 25}}{{$T = 1}}
            {{else if le $H 45}}{{$T = 2}}
            {{else if le $H 60}}{{$T = 3}}
            {{else if le $H 68}}{{$T = 4}}
            {{else if eq $H 69}}{{$T = 5}}
            {{else if le $H 80}}{{$T = 6}}
            {{else if le $H 90}}{{$T = 7}}
            {{else if le $H 100}}{{$T = 8}}
            {{else}}{{$T = 9}}{{end}}

            {{$H_Type = (index $hiveType $T)}}
            {{$H_Atk = (index $hiveAtk $T)}}
            {{$H_HP = (index $hiveHP $T)}}
            {{$H_Value = (index $hiveValue $T)}}
            {{$H_Image = (index $hiveImage $T)}}

            {{$fields := (cslice
				        (sdict "name" "Encounter:" "value" (joinStr "" "`" $H_Type "`") "inline" true)
				        (sdict "name" "Potential Loot:" "value" (joinStr " " $H_Value $e) "inline" true)
				        (sdict "name" "Previous Loot:" "value" (joinStr " " $Coins $e) "inline" true))}}
            {{/*sendMessage nil (joinStr " " "Roll:" $H "Type:" $H_Type "Attack:" $H_Atk "HP:" $H_HP "Value:" $H_Value "Player Attack:" $Attack)*/}}
            {{$dead := false}}

            {{/*INSERT ATTACK/DEFENCE STUFF HERE*/}}
            {{- range seq 0 28 -}}{{/* (add 1 (div $H_HP (mult 0.75 $Attack)))*/}}
            {{- $pDMG := 0 -}}
            {{- $hDMG := 0 -}}
            {{- $critbool := "No" -}}
            {{- if eq $dead false -}}

                {{- if and (gt $H_HP 0) (gt $HP 0) -}}
                    {{- $crit := randInt 0 5 -}}

                    {{- if eq $crit 0 -}}
                        {{- $pDMG = mult 2 $Attack -}}
                        {{- $critbool = "<:Crhit:886716842486419466>" -}}
                    {{- else -}}
                        {{- $pDMG = (roundFloor (mult (toFloat $Attack) (div (toFloat (randInt 80 110)) 100.0))) -}}
                        {{- $critbool = "<:SadBach:878032431989555232>" -}}
                    {{- end -}}

                    {{- $H_HP = sub $H_HP $pDMG -}}
                    {{- if lt $H_HP 0 -}}{{- $H_HP = 0 -}}{{- end -}}

                    {{- $fields = $fields.AppendSlice (cslice
                         (sdict "name" (joinStr "" $.User.Username ":") "value" (joinStr "" "Dealt `" $pDMG "` ⚔\nThrall: ❤`" $H_HP "`") "inline" true)
                         
                        ) -}}


                    {{- if le $H_HP 0 -}}
                        {{- $hiveDead = true -}}
                    {{- else -}}
                        {{- $hDMG = (roundFloor (mult (toFloat $H_Atk) (div (toFloat (randInt 80 110)) 100.0))) -}}
                        {{- $HP = sub $HP $hDMG -}}
                    {{- end -}}
                    {{- if le $HP 0 -}}
                        {{- $HP = 0 -}}
                        {{- $dead = true -}}
                    {{- end -}}
                    {{- $fields = $fields.AppendSlice (cslice
                         (sdict "name" (joinStr "" $H_Type ":") "value" (joinStr "" "Dealt `" $hDMG "` ⚔\nYou have ❤`" $HP "`") "inline" true)
                         (sdict "name" "Crit?" "value" $critbool "inline" true)
                        ) -}}

                {{- end -}}

            {{- end -}}
            {{- end -}}

            {{if eq $dead true}}
                {{if eq $hiveDead true}}
                    {{$hiveSlain = add $hiveSlain 1}}
                {{end}}
                {{$fields = $fields.AppendSlice (cslice (sdict "name" "Result:" "value" (joinStr "" "You died to the " $H_Type  "\nCoins lost: " $e " `" $Coins "` \nKills: " $Kills))
                                                        (sdict "name" "Cooldown:" "value" "30m" ))}}
                {{$data.Set "State" "false"}}
                {{dbSetExpire $.User.ID $cooldownKey 1 1800}}

            {{else if eq $hiveDead true}}
                {{$hiveSlain = add $hiveSlain 1}}
                {{$Coins = add $Coins $H_Value}}
                {{$Kills = add $Kills 1}}
                {{$fields = $fields.AppendSlice (cslice (sdict "name" "Result:" "value" (joinStr "" "**You killed the " $H_Type "**\n Current HP: `" $HP "` ❤\n Loot: " $e " `" $Coins "`\nKills: `" $Kills "`") ))}}
                {{addMessageReactions nil $x "☑" "❎"}}
            {{end}}

            {{$embed := cembed
				"title" (joinStr "" "__" $.User.Username "__ has entered the thrallpit")
				"thumbnail" (sdict "url" $H_Image)
				"fields" $fields}}
            {{editMessage nil $x $embed}}

            {{$data.Set "hiveSlain" $hiveSlain}}
            {{$data.Set "HP" $HP}}
            {{$data.Set "Coins" $Coins}}
            {{$data.Set "Attack" $Attack}}
            {{$data.Set "Kills" $Kills}}

            {{dbSet $.User.ID $pitKey $data}}

        {{else if eq $emoji "❎"}}
            {{$data.Set "State" "false"}}
            {{$data.Set "Start" 0}}
            {{if eq (toInt $start) 1}}
                {{/*ABORT*/}}
                {{deleteAllMessageReactions nil $x}}
                {{editMessage nil $x "You walked away from the thrallpit"}}

            {{else}}
                {{/*LEAVE*/}}
                {{deleteAllMessageReactions nil $x}}

                {{$data.Set "hiveSlain" (add ($data.Get "Kills") ($data.Get "hiveSlain"))}}
                {{$data.Set "CoinsEarnt" (add ($data.Get "Coins") ($data.Get "CoinsEarnt"))}}
                {{$data.Set "XP" (add ($data.Get "XP") ($data.Get "Coins"))}}
                {{dbSet $.User.ID $key (add ($data.Get "Coins") (toInt (dbGet $.User.ID $key).Value)) }}

                {{$XP = $data.Get "XP"}}
                {{$L := (toInt (pow $XP 0.4))}}
                {{if gt $L 100}}
                    {{$L = 100}}
                {{end}}
                {{$C := $data.Get "Coins"}}
                {{$H := $data.Get "HP"}}
                {{$K := $data.Get "Kills"}}
                {{$cooldown := (mult 60 (add (mult (log (add $L 5)) (add 7 (mult 0.07 $H))) (div (toFloat (mult $C $L)) 300.0)))}}

                {{$embed := cembed
                    "title" (joinStr "" "__" $.User.Username "__ has left the thrallpit")
                    "fields" (cslice
                        (sdict "name" "Results:" "value" (joinStr "" "Kills: `" $K "`\n" "Loot: " $e "`" $C "`\n" "HP: `" $H "`") "inline" true)
                        (sdict "name" (joinStr "" "XP: " "`" $XP "`") "value" (joinStr "" "Level: " "`" $L "`") "inline" true))
                    "footer" (sdict "text" (joinStr "" "Cooldown:" (div $cooldown 60) "m " (mod $cooldown 60) "s"))
				    }}

                {{dbSetExpire $.User.ID $cooldownKey 1 $cooldown}}


                {{editMessage nil $x $embed}}
                {{/*DO THE MONEY THING HERE AND ADD TO STATS*/}}

            {{end}}
            {{dbSet $.User.ID $pitKey $data}}
        {{end}}
	{{end}}
	{{end}}
{{end}}