{{ $curseasonal:= "[Season of the Splicer](https://youtu.be/crKGHG3stbY)" }}
{{ $expansion := "[Beyond Light](https://youtu.be/cthLUnEqT5k)" }}
{{ $avatar := (dbGet 0 "currentSeasonIcon").Value }}

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