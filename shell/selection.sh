#!/usr/bin/env bash
directorylist="Finished $(ls /)"
PS3='Directory to process? '
until [ "$directory" == "Finished" ]; do
    printf "%b" "\a\n\nSelect a directory to process:\n" >&2
    select directory in $directorylist; do
        if [ "$directory" = "Finished" ]; then
            echo "Finished processing directories."
            break
        elif [ -n "$directory" ]; then
            echo "You chose number $REPLY, processing $directory ..."
            break
        else
            echo "Invalid selection!"
        fi
    done
done
