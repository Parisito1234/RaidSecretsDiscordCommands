{{ $imageurl:= "https://www.bungie.net/pubassets/pkgs/148/148944/ae_season13_cal_EN.jpg?cv=3983621215&av=2328450057" }}
{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/f9ad6ac2aecaa89b2b4e075cc6e8b89f.png" }}
{{$embed := cembed 
    "color" 1772743
    "image" (sdict "url" $imageurl) 
    "footer" (sdict "icon_url" $avatar "text" "Season of the Chosen") 
}}
{{ sendMessage nil $embed }}