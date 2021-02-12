{{ $challenge_1 := sdict 
	"title" "Contender's Ascent III" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving III\" quest." 
	"reward" "•XP  •Empress Lore  •War Table Reputation (Medium)" 
}}
{{ $challenge_2 := sdict 
	"title" "Golden Reaper" 
	"description" "Acquire (300) Cabal Gold by playing strikes, Gambit, Crucible, public events, and more." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_3 := sdict 
	"title" "Crash and Converge" 
	"description" "Smash (15) Tribute Chests and focus (1) Tier 3 Umbral Engrams." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_4 := sdict 
	"title" "<Classified> Captain's Log" 
	"description" "Leave no stone unturned aboard the Glykon." 
	"reward" "•Classified" 
}}
{{ $challenge_5 := sdict 
	"title" "Failsafe Forward" 
	"description" "On Nessus, earn (10) progress by completing bounties, patrols, public events, and Lost Sectors." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Challenger's Apogee" 
	"description" "Reach Power Level 1300 by earning powerful rewards and Prime Engrams." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "Algorithmic Precision" 
	"description" "Calibrate weapons with precision final blows on Nessus. (100)" 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_8 := sdict 
	"title" "Primeval Entourage" 
	"description" "Defeat (100) Taken in Gambit. Earn bonus progress for defeating tougher combatants." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_9 := sdict 
	"title" "Iron Sharpens Iron" 
	"description" "Complete (15) Iron Banner matches. Earn bonus progress for wins." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_10 := sdict 
	"title" "Challenge Our Foes" 
	"description" "Complete (5) playlist strikes." 
	"reward" "•XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_3" (sdict 
	"title" "Seasonal Challenges - Week 3" 
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
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_3").Value)}}