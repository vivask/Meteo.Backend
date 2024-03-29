# flactool
[![Build Status](https://travis-ci.org/vivask/flactool-ip.svg?branch=main)](https://travis-ci.org/vivask/flactool)
[![GitHub release](https://img.shields.io/github/v/release/vivask/flactool.svg)](https://github.com/vivask/flactool/releases/latest)
![GitHub](https://img.shields.io/github/license/vivask/flactool.svg)

Multithreaded, batch tool for converting, concatenating and splitting audio files in flac, ape, wav formats


To use the program, you must first install the following packages:
- ffmpeg;
- Monkey's Audio Codec; https://github.com/fernandotcl/monkeys-audio;
- shntool;
- sox.

## Uasage:

    flactool [OPTION] 
-  -c convert dsf/wav/ape files in dir to flac
-  -C concat all flac files in dir to one flac file
-  -s split flac or ape files in dir
-  -d "path"
-  -f "file"
-  -h help
-  -p num core, default 4 (default 4)
-  -r remove source after operation
-  -v verbose

## Examples:
1. Convert all dsf/wav/ape files from ~/apedir (with subdirectories) to flac

    flactool -d ~/apedir -c 

2. All ape and wav files from the current directory (with subdirectories) are split with cue (if there is a cue file with a name similar to ape or wav) with conversion to flac. With the subsequent removal of the original ape and wav files

    flactool -s -r

3. Convert all dsf/wav/ape files from current dir (with subdirectories) to flac then concatenate the converted

    flactool -c -C


### Build 
To build from source code you need run the following commands:

    git clone https://github.com/vivask/flactool.git
    cd flactool
    install.sh