module Main exposing (main)

import Html exposing (Html, a, div, h1, li, text, ul)
import Html.Attributes exposing (href, style)


main : Html msg
main =
    div [] [ header, content ]


header : Html msg
header =
    h1 [] [ text "Useful Links" ]


listElement : String -> String -> Html msg
listElement link title =
    li []
        [ a
            [ href link ]
            [ text title ]
        ]


content : Html msg
content =
    div []
        [ ul []
            [ listElement
                "https://elm-lang.org"
                "homepge"
            , listElement
                "https://elm-lang.org"
                "iwomepage"
            , listElement
                "https://elm-lang.org"
                "aa"
            ]
        ]
