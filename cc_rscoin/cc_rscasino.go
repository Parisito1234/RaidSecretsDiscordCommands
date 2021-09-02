{{$embed := cembed
	"title" "__RSCoin Commands__"
	"fields" (cslice
	(sdict "name" "Balance" "value" "`-rscoin <bal|balance>`" "inline" true )
	(sdict "name" "Pay" "value" "`-rscoin <give|pay> <amount>`" "inline" true )
	(sdict "name" "Leaderboard" "value" "`-rscoin <top|list>`" "inline" true )
	(sdict "name" "Claim (One-time)" "value" "`-rscoin claim`" "inline" true )
	(sdict "name" "__Games__" "value" "These do not start with `-rscoin`" "inline" false )
	(sdict "name" "Roulette" "value" "`-roulette <tile> <amount>`" "inline" true )
	(sdict "name" "Coin Flip" "value" "`-coinflip <amount>`" "inline" true )
	(sdict "name" "Blackjack" "value" "`-blackjack <bet>`" "inline" true )
	(sdict "name" "Duel" "value" "`-duel <@user> <bet>`" "inline" true ))
}}
{{ sendMessage nil $embed }}