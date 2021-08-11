{{ /* Trigger: Command: twab */ }}

{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}
{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Bungie News page" "value" "[Check out the page for all TWABs here!](https://www.bungie.net/en/News/Index?tag=news-destiny&page=0)" "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}
{{ sendMessage nil $embed }}