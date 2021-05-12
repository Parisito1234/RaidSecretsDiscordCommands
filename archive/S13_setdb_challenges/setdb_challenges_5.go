{{ $challenge_1 := sdict 
	"title" "Contender's Ascent V" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving V\" quest." 
	"reward" "•XP  •Empress Lore  •War Table Reputation (Large)" 
}}
{{ $challenge_2 := sdict 
	"title" "Golden Reaper" 
	"description" "Acquire (500) Cabal Gold by playing strikes, Gambit, Crucible, public events, and more." 
	"reward" "•XP  •War Table Reputation (Large)" 
}}
{{ $challenge_3 := sdict 
	"title" "Graven Scrawl" 
	"description" "Investigate (3) cryptic notes left in the captain's log." 
	"reward" "•XP" 
}}
{{ $challenge_4 := sdict 
	"title" "Lenses in Focus" 
	"description" "Go to the Prismatic Recaster and discover how to unlock (5) more lenses." 
	"reward" "•XP" 
}}
{{ $challenge_5 := sdict 
	"title" "Apex Armorer" 
	"description" "Masterwork a piece of armor." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Salvager's Salvo Armament" 
	"description" "Acquire the Seasonal Ritual weapon." 
	"reward" "•4x XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "In It for Infamy" 
	"description" "Earn (5) Infamy ranks in Gambit." 
	"reward" "•4x XP  •Bright Dust" 
}}
{{ $challenge_8 := sdict 
	"title" "Ultimate Champion" 
	"description" "Defeat (60) Champions in any Nightfall: The Ordeal strikes. Earn bonus progress at higher difficulty tiers." 
	"reward" "•4x XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_5" (sdict 
	"title" "Seasonal Challenges - Week 5" 
	"color" 1772743
	"fields" (cslice 
		(sdict "name" ($challenge_1.Get "title") "value" (joinStr "" ($challenge_1.Get "description") " **|** `" ($challenge_1.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_2.Get "title") "value" (joinStr "" ($challenge_2.Get "description") " **|** `" ($challenge_2.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_3.Get "title") "value" (joinStr "" ($challenge_3.Get "description") " **|** `" ($challenge_3.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_4.Get "title") "value" (joinStr "" ($challenge_4.Get "description") " **|** `" ($challenge_4.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_5.Get "title") "value" (joinStr "" ($challenge_5.Get "description") " **|** `" ($challenge_5.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_6.Get "title") "value" (joinStr "" ($challenge_6.Get "description") " **|** `" ($challenge_6.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_7.Get "title") "value" (joinStr "" ($challenge_7.Get "description") " **|** `" ($challenge_7.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_8.Get "title") "value" (joinStr "" ($challenge_8.Get "description") " **|** `" ($challenge_8.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_5").Value)}}