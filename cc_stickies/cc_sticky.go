{{ /* Trigger: Command: Sticky */ }}
{{ /* Works with sticky_trigger command in the RaidSecretsDiscordCommands repository */ }}
{{ /* Channel/User Restrictions: Moderator role required */ }}
{{ /* This command was created by Black Wolf and modified by Parisito */ }}

{{/*
	Trigger type: Command
	Trigger: sticky

	Enable sticky message: sm This is a test
	Disable sticky message: sm
	If you still don't get how it works look at thi gif: https://i.imgur.com/ohRubPw.gif

	Copyright (c): Black Wolf, 2021
	License: MIT
	Repository: https://github.com/BlackWolfWoof/yagpdb-cc/
*/}}
{{/*NOTE: YOU NEED THE 'sticky message 2.gtmpl' CODE FOR THIS TO WORK*/}}
{{$perms := "ManageMessages"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}
{{$cooldown := "1h"}}{{/*FORMAT EXAMPLE: "10s" "20m10s" "1mo10d" "3w2h"*/}}
{{/*The default cooldown is how long the sticky message shouldn't be sent in the same channel*/}}


{{if (in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms)}}
	{{if .StrippedMsg}}
		{{$color := 0}}{{$old := 0}}{{range .Guild.Roles}}{{if and (in $.Member.Roles .ID) (ne .Color 0) (gt .Position $old)}}{{$old = .Position}}{{$color = .Color}}{{end}}{{end}}
		
		{{$img := ""}}{{$text := .StrippedMsg}}
		{{with reFindAllSubmatches `(?:(?P<TxtSnip1>(?:.*[\r\n]?){0,}))?(?:-img\s(?P<Link>(?:https?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+))(?P<TxTSnip2>(?:.*[\r\n]?){0,})` $text}}
			{{$img = index . 0 2}}
			{{$text = print (index . 0 1) (index . 0 3)}}
		{{end}}

		{{with reFindAllSubmatches `\A((?:.|[\r\n])*)(-d\s(?P<Duration>(?:(?:\d+)?(?:months?|mo|minutes?|s|seconds?|m|hours?|h|days?|d|weeks?|w|years?|y|permanent|p)){1,}))(\s(?:.|[\r\n])*)?\z` $text}}
			{{$cooldown = index . 0 3}}
			{{$text = print (index . 0 1) (index . 0 4)}}
		{{end}}

		{{dbSet 0 "stickymessage" (sdict "message" $text "author" .User.String "color" $color "cooldown" $cooldown "img" $img)}}
		{{sendMessage nil "The sticky message was enabled for opt-in channels and saved!"}}
	{{else}}
		{{dbDel 0 "stickymessage"}}
		{{$del := 9}}{{if .IsPremium}}{{$del = 49}}{{end}}
		{{range dbTopEntries "smchannel" $del 0}}
			{{deleteMessage .UserID (toInt .Value) 0}}
			{{dbDel .UserID "smchannel"}}
		{{end}}
		{{sendMessage nil "Sticky messages for opt-in channels are now disabled and deleted again."}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}