{{ $curseasonal:= "[Dropbox](https://www.dropbox.com/sh/v300chsprb653fr/AADux62OjOhIj599_IqBKgFha)" }}
{{ $avatar := "https://www.bungie.net/common/destiny2_content/icons/8ff7269628a34bf7adc71ee0b518f9a2.png" }}

{{$embed := cembed 
    "color" 1772743
    "fields" (cslice 
        (sdict "name" "Bungie Press Kit" "value" $curseasonal "inline" false) 
    ) 
    "thumbnail" (sdict "url" $avatar) 
}}

{{ sendMessage nil $embed }}