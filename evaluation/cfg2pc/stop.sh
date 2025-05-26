# !/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

stop_session() {
    local session_name="$1"
    if screen -ls | grep -q "$session_name"; then
        screen -S "$session_name" -X quit
        echo "会话已停止: $session_name"
    else
        echo "未找到会话: $session_name"
    fi
}

stop_session "session_geth_1"
stop_session "session_geth_2"
stop_session "session_geth_3"

stop_session "session_eICN_1"
stop_session "session_eICN_2"
stop_session "session_eICN_3"

stop_session "session_metrics"