# !/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

check_session() {
    local session_name="$1"
    if screen -ls | grep -q "$session_name"; then
        echo "session is running: $session_name"
        return 0
    fi
    echo "cannot find session: $session_name"
    return 2
}

# Get debug parameter with default value false
DEBUG=${1:-false}

# bash ./evaluation/onestep/3/start_avalanche.sh

echo ">>> >>> >>> start metrics"
screen -L -S session_metrics -dm ./metrics/run_metrics.sh
sleep 2s
check_session "session_metrics"
if [ $? -eq 2 ]; then
    echo "Terminating script due to missing session: session_metrics"
    exit 1
fi

# echo ">>> >>> >>> start analysis"
# screen -L -S session_analysis -dm ./analysis/run_analysis.sh 1 test evaluation/configs/config1.yaml,evaluation/configs/config2.yaml,evaluation/configs/config3.yaml false
# sleep 2s
# check_session "session_analysis"

echo ">>> >>> >>> start deploy"
./scripts/deploy.sh --config=evaluation/configs/config1.yaml &
./scripts/deploy.sh --config=evaluation/configs/config2.yaml &
./scripts/deploy.sh --config=evaluation/configs/config3.yaml &
wait

echo ">>> >>> >>> start init values"
./scripts/init_values.sh --values="100,100,100" --config=evaluation/configs/config1.yaml &
./scripts/init_values.sh --values="100,100,100" --config=evaluation/configs/config2.yaml &
./scripts/init_values.sh --values="100,100,100" --config=evaluation/configs/config3.yaml &
wait

echo ">>> >>> >>> start eICN"
screen -L -S session_eICN_1 -dm go run main.go --config=evaluation/configs/config1.yaml --log=logs/1.log --debug=$DEBUG
screen -L -S session_eICN_2 -dm go run main.go --config=evaluation/configs/config2.yaml --log=logs/2.log --debug=$DEBUG
screen -L -S session_eICN_3 -dm go run main.go --config=evaluation/configs/config3.yaml --log=logs/3.log --debug=$DEBUG
sleep 3s
check_session "session_eICN_1"
if [ $? -eq 2 ]; then
    echo "Terminating script due to missing session: session_eICN_1"
    exit 1
fi
check_session "session_eICN_2"
if [ $? -eq 2 ]; then
    echo "Terminating script due to missing session: session_eICN_2"
    exit 1
fi
check_session "session_eICN_3"
if [ $? -eq 2 ]; then
    echo "Terminating script due to missing session: session_eICN_3"
    exit 1
fi

echo ">>> >>> >>> start regist"
./scripts/regist_eICN.sh --config=evaluation/configs/config1.yaml --target-config=evaluation/configs/config2.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config1.yaml --target-config=evaluation/configs/config3.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config2.yaml --target-config=evaluation/configs/config1.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config3.yaml --target-config=evaluation/configs/config1.yaml

# echo ">>> >>> >>> start cross send"
# sleep 3s
# ./scripts/cross_send.sh --chain-ids="2,3" --value="100" --app-identifier="State" --app-value-id="0" --config=evaluation/configs/config1.yaml

# echo ">>> >>> >>> start cross send workload"
# sleep 3s
# ./scripts/cross_send_workload.sh --chain-ids="2,3" --app-identifier="State" --write-conflict-rate=0 --transaction-number=5 --worker-cfg-files=evaluation/configs/config2.yaml,evaluation/configs/config3.yaml --config=evaluation/configs/config1.yaml

echo ">>> >>> >>> Finished!"