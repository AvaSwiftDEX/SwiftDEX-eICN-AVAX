# !/bin/bash

check_session() {
    local session_name="$1"
    if screen -ls | grep -q "$session_name"; then
        echo "session is running: $session_name"
        return 2
    fi
    echo "cannot find session: $session_name"
    return 0
}

echo ">>> >>> >>> start geth"
screen -L -S session_geth_1 -dm ../SuperRunner-go-ethereum/build/bin/geth --dev --dev.period 3 --keystore ./node/keystore --allow-insecure-unlock --http --http.api eth,web3,net,miner,txpool,admin --ws --ws.api eth,web3,net --http.port 8545 --ws.port 8546
screen -L -S session_geth_2 -dm ../SuperRunner-go-ethereum/build/bin/geth --dev --dev.period 3 --keystore ./node/keystore --allow-insecure-unlock --http --http.api eth,web3,net,miner,txpool,admin --ws --ws.api eth,web3,net --http.port 8555 --ws.port 8556
screen -L -S session_geth_3 -dm ../SuperRunner-go-ethereum/build/bin/geth --dev --dev.period 3 --keystore ./node/keystore --allow-insecure-unlock --http --http.api eth,web3,net,miner,txpool,admin --ws --ws.api eth,web3,net --http.port 8565 --ws.port 8566
sleep 2s
check_session "session_geth_1"
check_session "session_geth_2"
check_session "session_geth_3"

echo ">>> >>> >>> start metrics"
screen -L -S session_metrics -dm ./metrics/run_metrics.sh

echo ">>> >>> >>> start deploy"
./scripts/deploy.sh --config=evaluation/configs/config1.yaml
./scripts/deploy.sh --config=evaluation/configs/config2.yaml
./scripts/deploy.sh --config=evaluation/configs/config3.yaml

echo ">>> >>> >>> start eICN"
screen -L -S session_eICN_1 -dm go run main.go --config=evaluation/configs/config1.yaml --log=logs/1.log
screen -L -S session_eICN_2 -dm go run main.go --config=evaluation/configs/config2.yaml --log=logs/2.log
screen -L -S session_eICN_3 -dm go run main.go --config=evaluation/configs/config3.yaml --log=logs/3.log
sleep 3s
check_session "session_eICN_1"
check_session "session_eICN_2"
check_session "session_eICN_3"

echo ">>> >>> >>> start regist"
./scripts/regist_eICN.sh --config=evaluation/configs/config1.yaml --target-config=evaluation/configs/config2.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config1.yaml --target-config=evaluation/configs/config3.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config2.yaml --target-config=evaluation/configs/config1.yaml
./scripts/regist_eICN.sh --config=evaluation/configs/config3.yaml --target-config=evaluation/configs/config1.yaml

echo ">>> >>> >>> start cross send"
sleep 3s
./scripts/cross_send.sh --chain-ids="2,3" --value="100" --config=evaluation/configs/config1.yaml

echo ">>> >>> >>> Finished!"