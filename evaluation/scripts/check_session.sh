# !/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

local session_name="$1"
if screen -ls | grep -q "$session_name"; then
    echo "session is running: $session_name"
    exit 0
fi
echo "cannot find session: $session_name"
exit 2