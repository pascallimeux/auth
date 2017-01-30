#!/bin/bash
LOGDIR="/var/log/auth"
USER="pascal"
if [ ! -d "$LOGDIR" ]; then
	echo "Create log directory"
    sudo mkdir -p $LOGDIR
fi
    sudo chown -R $USER $LOGDIR
echo "AUTHGO process started."
if [ "$1" ==  "init" ]
 then
     rm auth.log
 	 echo "start auth init"
     ./auth-go init &
 else
 	 echo "start auth"
     ./auth &
fi

read -rst 0.5
tail -f auth.log

