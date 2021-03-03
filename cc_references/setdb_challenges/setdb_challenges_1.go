{{ $challenge_1 := sdict 
	"title" "Contender's Ascent" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving\" quest. Then defeat (75) Cabal with Rocket Launchers anywhere in the system." 
	"reward" "•XP  •Code Duello  •Empress Lore  •War Table Reputation (Medium)" 
}}
{{ $challenge_2 := sdict 
	"title" "Golden Reaper" 
	"description" "Acquire (200) Cabal Gold by playing strikes, Gambit, Crucible, public events, and more." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_3 := sdict 
	"title" "Crash and Converge" 
	"description" "Smash (5) Tribute Chests and focus (5) Season of the Chosen Engrams using Charges from the Hammer of Proving." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_4 := sdict 
	"title" "Lenses in Focus" 
	"description" "Participate in the Battlegrounds playlist to unlock your first lens." 
	"reward" "•XP" 
}}
{{ $challenge_5 := sdict 
	"title" "Challenger's Aspiration" 
	"description" "Complete (3) weekly playlist challenges." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Icebound" 
	"description" "On Europa, earn (10) progress by completing bounties, patrols, public events, and Lost Sectors." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "Hail of Bullets" 
	"description" "Calibrate Kinetic weapons on Europa. Earn bonus progress in Lost Sectors. (200)" 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_8 := sdict 
	"title" "Dredgin' Up Victory" 
	"description" "Complete (15) Gambit matches. Earn bonus progress for wins." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_9 := sdict 
	"title" "Flourish of Power" 
	"description" "Defeat (50) Guardians in the Mayhem playlist with Super abilities." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_10 := sdict 
	"title" "Dominance Operandi: Fallen" 
	"description" "Defeat (200) Fallen combatants in strikes. Earn bonus progress for defeating tougher combatants." 
	"reward" "•XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_1" (sdict 
	"title" "Seasonal Challenges - Week 1" 
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
		(sdict "name" ($challenge_9.Get "title") "value" (joinStr "" ($challenge_9.Get "description") " **|** `" ($challenge_9.Get "reward") "`") "inline" false) 
		(sdict "name" ($challenge_10.Get "title") "value" (joinStr "" ($challenge_10.Get "description") " **|** `" ($challenge_10.Get "reward") "`") "inline" false) 
	) 
)}}
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_1").Value)}}