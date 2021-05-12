{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Season of the Splicer" "value" "[Patchnotes](https://www.bungie.net/en/Explore/Detail/News/50339)" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
{{ sendMessage nil $embed }}