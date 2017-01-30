#!/bin/bash
PID=`pidof auth-go`
if [ -n "$PID" ]
then
   kill -9 $PID
   echo "AUTHGO process stopped"
else
   echo "Could not send SIGTERM to kill AUTHGO, probably it does not work:" >&2
fi