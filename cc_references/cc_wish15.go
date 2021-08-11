{{ /* Trigger: Command: wish15 */ }}

{{ $status:= "Currently exploring new seasonal lore and story." }}
{{ $redditurl:= "[Megathread](https://www.reddit.com/r/raidsecrets/comments/9z1bw4/wish_15_megathreadfor_serious_discussion/)" }}
{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/fc5791eb2406bf5e6b361f3d16596693.png" }}

{{$embed := cembed 
    "title" "__Wish 15__ status and links"
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Status:" "value" $status "inline" false) 
        (sdict "name" "Discord channel:" "value" "<#539600273728077824>" "inline" true) 
        (sdict "name" "Reddit" "value" $redditurl "inline" true) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
    "author" (sdict "name" "RaidSecrets") 
}}

{{ sendMessage nil $embed }}