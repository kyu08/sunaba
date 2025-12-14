module Main exposing (..)

import Http
import Json.Decode exposing (Decoder, field, string)


type Msg
    = GotGif (Result Http.Error String)


getRandomGif : Cmd Msg
getRandomGif =
    Http.get
        { url = "https://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag=cat"
        , expect = Http.expectJson GotGif gifDecoder
        }


gifDecoder : Decoder String
gifDecoder =
    field "data" (field "image_url" string)
