{{ /* Trigger: Regex: \A */ }}
{{ /* Works with sticky command in the RaidSecretsDiscordCommands repository */ }}
{{ /* Channel/User Restrictions: None */ }}
{{ /* Output errors as command response: DISABLED -- High-volume servers can cause database executions to overload, leading to command failure, so disable this. */ }}
{{ /* This command was created by Black Wolf and modified by Parisito */ }}

{{/*
	Trigger type: Regex
	Trigger: \A

	Copyright (c): Black Wolf, 2021
	License: MIT
	Repository: https://github.com/BlackWolfWoof/yagpdb-cc/
*/}}
{{/*NOTE: YOU NEED THE 'sticky message 2 command.gtmpl' CODE FOR THIS TO WORK*/}}
{{$whitelist := true}}
{{/*This enables the whitelist of channels where it should run*/}}
{{$channels := cslice CHANNEL_ID_1 CHANNEL_ID_2 }}
{{/*Replace these channel id's with your channels if you have enabled the whitelist*/}}


{{if not $whitelist}}{{$channels = cslice .Channel.ID}}{{end}}
{{if and (dbGet 0 "stickymessage") (dbGet .Channel.ID "smcooldown"|not) (in $channels .Channel.ID)}}
	{{$db := (dbGet 0 "stickymessage").Value}}
	{{$message := cembed "author" (sdict "name" "Reminder!" "icon_url" "https://cdn.discordapp.com/emojis/587253903121448980.png") "color" $db.color "description" $db.message "image" (sdict "url" (or $db.img ""))}}

	{{if $db := dbGet .Channel.ID "smchannel"}}
		{{deleteMessage nil (toInt $db.Value) 0}}
	{{end}}
	{{$id := sendMessageRetID nil $message}}
	{{dbSet .Channel.ID "smchannel" (str $id)}}
	{{dbSetExpire .Channel.ID "smcooldown" true (toInt (toDuration $db.cooldown).Seconds)}}
{{end}}