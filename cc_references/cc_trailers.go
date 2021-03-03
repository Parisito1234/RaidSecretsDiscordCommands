{{ $curseasonal:= "[Season of the Chosen](https://youtu.be/_XlFBp_ZmyE)" }}
{{ $expansion := "[Beyond Light](https://youtu.be/cthLUnEqT5k)" }}
{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/f9ad6ac2aecaa89b2b4e075cc6e8b89f.png" }}
 
{{$embed := cembed 
    "title" "Currently relevant Destiny 2 trailers:"
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Current Season" "value" $curseasonal "inline" false) 
        (sdict "name" "Current Expansion" "value" $expansion "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
    "author" (sdict "name" "RaidSecrets") 
}}
 
{{ sendMessage nil $embed }}