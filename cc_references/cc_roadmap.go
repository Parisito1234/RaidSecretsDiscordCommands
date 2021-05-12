{{ $imageurl:= (dbGet 0 "roadmapURL").Value }}
{{ $avatar := (dbGet 0 "currentSeasonIcon").Value }}
{{$embed := cembed 
    "color" 1772743
    "image" (sdict "url" $imageurl) 
    "footer" (sdict "icon_url" $avatar "text" "Season of the Splicer") 
}}
{{ sendMessage nil $embed }}