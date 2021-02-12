{{ $challenge_1 := sdict 
	"title" "Contender's Ascent II" 
	"description" "Report to the War Table in the H.E.L.M. and complete the \"Challenger's Proving II\" quest." 
	"reward" "•XP  •Empress Lore  •War Table Reputation (Medium)" 
}}
{{ $challenge_2 := sdict 
	"title" "Golden Reaper" 
	"description" "Acquire (200) Cabal Gold by playing strikes, Gambit, Crucible, public events, and more." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_3 := sdict 
	"title" "Crash and Converge" 
	"description" "Smash (10) Tribute Chests and focus (10) Season of the Chosen Engrams using Charges from the Hammer of Proving." 
	"reward" "•XP  •War Table Reputation (Medium)" 
}}
{{ $challenge_4 := sdict 
	"title" "The Bigger They Are…" 
	"description" "Defeat (60) Elite or Boss Cabal anywhere in the system." 
	"reward" "•XP" 
}}
{{ $challenge_5 := sdict 
	"title" "Trials of the Tinker" 
	"description" "Unlock (12) artifact mods." 
	"reward" "•Double XP  •Bright Dust" 
}}
{{ $challenge_6 := sdict 
	"title" "Chosen Cosmonaut" 
	"description" "In the Cosmodrome, earn (10) progress by completing bounties, patrols, public events, and Lost Sectors." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_7 := sdict 
	"title" "Expose to the Elements" 
	"description" "Calibrate elemental weapons in the Cosmodrome. (200)" 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_8 := sdict 
	"title" "Drifter's Chosen" 
	"description" "Earn (250) points by banking Motes, defeating Blockers, and defeating Invaders in Gambit." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_9 := sdict 
	"title" "Entertain Lord Shaxx" 
	"description" "Complete (15) matches in the Showdown Crucible playlist. Earn bonus progress for wins." 
	"reward" "•XP  •Bright Dust" 
}}
{{ $challenge_10 := sdict 
	"title" "Dominance Operandi: Cabal" 
	"description" "Defeat (5) Cabal bosses in strikes." 
	"reward" "•XP  •Bright Dust" 
}}

{{dbSet 0 "reference_challenges_2" (sdict 
	"title" "Seasonal Challenges - Week 2" 
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
{{sendMessage nil (cembed (dbGet 0 "reference_challenges_2").Value)}}