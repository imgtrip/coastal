#!/usr/bin/env bash

# https://stackoverflow.com/questions/9893667/is-there-a-way-to-write-a-bash-function-which-aborts-the-whole-execution-no-mat
trap "exit 1" TERM
export TOP_PID=$$


LINUX="Linux"
MAC="Mac"
WINDOWS="Windows"

currentOS(){
    unameOut="$(uname -s)"
    case "${unameOut}" in
        Linux*)     machine=LINUX;;
        Darwin*)    machine=MAC;;
        CYGWIN*)    machine=Cygwin;;
        MINGW*)     machine=MinGw;;
        *)          machine="UNKNOWN ${unameOut}"
    esac

    if [ "$machine" == LINUX ] || [ "$machine" == MAC ];then
        echo "$machine"
    fi

    echo "$WINDOWS"
}

isWindows(){
    current=$(currentOS)
    if [ "$current" == "$WINDOWS" ];then
        echo true
    else
        echo false
    fi
}

suffix(){
    if $(isWindows) ;then
        echo "$1.exe"
    else
        echo "$1"
    fi
}

getRootDir(){
    # https://stackoverflow.com/questions/3466166/how-to-check-if-running-in-cygwin-mac-or-linux
    unameOut="$(uname -s)"
    case "${unameOut}" in
        Linux*)     machine=Linux;;
        Darwin*)    machine=Mac;;
        CYGWIN*)    machine=Cygwin;;
        MINGW*)     machine=MinGw;;
        *)          machine="UNKNOWN ${unameOut}"
    esac

    ROOT_DIR=""
    if [ "$machine" == "Linux" ] || [ "$machine" == "Mac" ]; then
        ROOT_DIR="$(dirname $(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd ))"
    elif [ "$machine" == "Cygwin" ] || [ "$machine" == "MinGw" ]; then
        ROOT_DIR="$(dirname $(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd -W))"
    fi

    if [ -z "$ROOT_DIR" ]; then
        printf '%s\n' "ERROR:no adapter for the current system:${machine}" >&2
        printf '%s\n' "exit 1" >&2
        kill -s TERM ${TOP_PID}
        exit 1
    fi

    echo "$ROOT_DIR"
}




