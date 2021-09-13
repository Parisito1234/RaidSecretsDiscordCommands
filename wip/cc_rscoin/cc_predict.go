{{ $perms := "ManageServer"}}
{{ $key := "RSCoinBalance" }}
{{ $coinIcon := "<:RSStonkCoin:869340420692394095>" }}
{{ $args := parseArgs 1 "Syntax is `< start | stop > <option 1> <option 2>`"
	(carg "string" "action - start, stop")
	(carg "string" "Options for predictions") 
	(carg "string" "Options for predictions") }}
{{ $action := lower ($args.Get 0) }}
{{ $lotteryPool := toInt ((dbGet 204255221017214977 $key).Value) }}

{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{ if eq $action "start" }}
		{{/* Start a new prediction with 2 more arguments */}}
		{{ if $args.IsSet 2 }}
			{{ $option1 := $args.Get 1 }}
			{{ $option2 := $args.Get 2 }}

			{{$embed := cembed 
			    "color" 1772743
			    "fields" (cslice 
			        (sdict "name" ":one:" "value" $option1 "inline" false) 
			        (sdict "name" ":two:" "value" $option2 "inline" false) 
			    ) 
			    "author" (sdict "name" "RaidSecrets Prediction") 
			}}
			{{ $x := sendMessageRetID  nil $embed }}
			{{$cid := .Message.ChannelID}}
			{{ addMessageReactions $cid $x ":one:" ":two:"}}
			{{ dbSet 0 (joinStr "" "prediction_" (str $cid)) (str $x) }}
		{{ else }}
			{{ sendMessage nil "You need 2 options to run this command." }}
		{{ end }}

	{{ else if eq $action "stop" }}
		{{/* Stop currently running prediction, clear from DB table */}}
		{{ if $args.IsSet 1 }}
			{{ $cid := $args.Get 1 }}
			{{ $dbKey := (joinStr "" "prediction_" $cid ) }}
			{{ $messageID := toInt ((dbGet 0 $dbKey).Value) }}
			{{ if $messageID}} 
				{{ sendMessage nil $messageID }}
				{{ $message := (getMessage $cid $messageID) }}
				{{ range $message.Reactions }}
					{{ .Emoji.Name }} {{ .Emoji.ID }} 
				{{ end }}
				{{ dbDel 0 $dbKey }}
				{{ deleteAllMessageReactions $cid $messageID }}
				{{ addReactions "👍" }}
			{{ else }}
				{{ sendMessage nil "There's currently not a prediction running in that channel." }}
			{{ end }}
		{{ else }}
			{{ sendMessage nil "Usage: `predict stop channelID`"}}
		{{ end }}
	{{ else if eq $action "list" }}
		{{ range (dbTopEntries "prediction_%" 10 0) }}
			`{{.UserID}}` **:** `{{.Key}}` **:** `{{toInt (.Value)}}`
		{{ end }}
	{{ end }}
	
{{ else }}
	{{ sendMessage nil "You don't have permission to do that." }}
	{{ deleteTrigger 10 }}
	{{ deleteResponse 10 }}
{{ end }}