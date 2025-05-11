import argparse
import requests
def get_args():
    parser = argparse.ArgumentParser(description="Analysis metrics")
    parser.add_argument("--metrics-server-url", type=str, default="127.0.0.1:8090", help="metrics collector server websocket url")
    parser.add_argument("--logfile", type=str, default="logs/analysis.log", help="log file path")
    args = parser.parse_args()
    return args.metrics_server_url, args.logfile

def main():
    metrics_server_url, logfile = get_args()
    
    url = f"http://{metrics_server_url}/metrics/get"

    response = requests.get(url)
    print(response.json())

if __name__ == "__main__":
    main()
