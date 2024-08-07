module Concourse.Cli exposing (Cli(..), clis, downloadUrl, id, label)

import Api.Endpoints as Endpoints
import Url.Builder


clis : List Cli
clis =
    [ OSX, OSXArm, Windows, Linux ]


type Cli
    = OSX
    | OSXArm
    | Windows
    | Linux


downloadUrl : Cli -> String
downloadUrl cli =
    let
        platformName =
            case cli of
                OSX ->
                    "darwin"

                OSXArm ->
                    "darwin"

                Windows ->
                    "windows"

                Linux ->
                    "linux"

        archName =
            case cli of
                OSXArm ->
                    "arm"

                _ ->
                    "amd64"
    in
    Endpoints.Cli
        |> Endpoints.toString
            [ Url.Builder.string "arch" archName
            , Url.Builder.string "platform" platformName
            ]


label : Cli -> String
label cli =
    let
        platformName =
            case cli of
                OSX ->
                    "OS X"

                OSXArm ->
                    "OS X arm"

                Windows ->
                    "Windows"

                Linux ->
                    "Linux"
    in
    "Download " ++ platformName ++ " CLI"


id : Cli -> String
id cli =
    case cli of
        OSX ->
            "osx"

        OSXArm ->
            "osx-arm"

        Windows ->
            "windows"

        Linux ->
            "linux"
