import os, sys
sys.path.insert(0, os.path.abspath('.'))

import json
from typing import Dict, List, Tuple

"""
const (
	MetricsCMPhaseMasterIssue uint8 = iota
	MetricsCMPhaseWorkerPrepareUnconfirm
	MetricsCMPhaseWorkerPrepareConfirm
	MetricsCMPhaseMasterReceiveUnconfirm
	MetricsCMPhaseMasterReceiveConfirm
	MetricsCMPhaseMasterCommitUnconfirm
	MetricsCMPhaseMasterCommitConfirm
	MetricsCMPhaseMasterRollbackUnconfirm
	MetricsCMPhaseMasterRollbackConfirm
	MetricsCMPhaseWorkerCommitUnconfirm
	MetricsCMPhaseWorkerCommitConfirm
	MetricsCMPhaseWorkerRollbackUnconfirm
	MetricsCMPhaseWorkerRollbackConfirm
)
"""

# transactionHash -> phase -> []metrics

class Metric:
    MasterIssue = "MasterIssue"
    WorkerPrepareUnconfirm = "WorkerPrepareUnconfirm"
    WorkerPrepareConfirm = "WorkerPrepareConfirm"
    MasterReceiveUnconfirm = "MasterReceiveUnconfirm"
    MasterReceiveConfirm = "MasterReceiveConfirm"
    MasterCommitUnconfirm = "MasterCommitUnconfirm"
    MasterCommitConfirm = "MasterCommitConfirm"
    MasterRollbackUnconfirm = "MasterRollbackUnconfirm"
    MasterRollbackConfirm = "MasterRollbackConfirm"
    WorkerCommitUnconfirm = "WorkerCommitUnconfirm"
    WorkerCommitConfirm = "WorkerCommitConfirm"
    WorkerRollbackUnconfirm = "WorkerRollbackUnconfirm"
    WorkerRollbackConfirm = "WorkerRollbackConfirm"

    def __init__(self, data: Dict):
        self.transactionHash = data['transactionHash']
        self.cmHash = data['cmHash']
        self.chainId = data['chainId']
        self.height = data['height']
        self.phase = data['phase']
        self.isConfirmed = data['isConfirmed']
        self.byHeader = data['byHeader']
        self.timestamp = data['timestamp']
        self.txHash = data['txHash']
        # TODO: add following fields to MetricsData in types.go
        self.forged = data['forged']
        self.retry = data['retry']
        self.gas = data['gas']
        self.root = data['root']
        self.realRoot = data['realRoot']
        self.fromChainId = data['fromChainId']
        self.fromHeight = data['fromHeight']
    

cmp = lambda x, y: x if x[0] < y[0] else y
fmt2f = lambda x: float(f"{x:.2f}")
avglist = lambda x: sum(x) / len(x) if len(x) > 0 else 0
avg = lambda x: avglist(list(x))

class PhaseLatencyMetric:
    def __init__(self):
        self.m_issue_2_w_ucf_p = 0
        self.w_ucf_p_2_m_cf_recv = 0
        self.w_ucf_p_2_m_cf_c = 0
        self.w_ucf_p_2_m_cf_r = 0
        self.m_cf_c_2_w_cf_c = 0
        self.m_cf_r_2_w_cf_r = 0
        self.overall = 0

class GasMetric:
    def __init__(self) -> None:
        self.m_issue = 0
        self.w_ucf_p = 0
        self.w_cf_p = 0
        self.m_cf_recv = 0
        self.m_cf_c = 0
        self.m_cf_r = 0
        self.w_cf_c = 0
        self.w_cf_r = 0

class TransactionMetric:
    def __init__(self):
        self.MasterIssue: Tuple[int, int] = (0, 0)
        self.WorkerPrepareUnconfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.WorkerPrepareConfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.MasterReceiveUnconfirm: Dict[str, Tuple[int, int]] = {} # fromChainId -> (the earliest timestamp, gas)
        self.MasterReceiveConfirm: Dict[str, Tuple[int, int]] = {} # fromChainId -> (the earliest timestamp, gas)
        self.MasterCommitUnconfirm: Dict[str, Tuple[int, int]] = {} # fromChainId -> (the earliest timestamp, gas)
        self.MasterCommitConfirm: Tuple[int, int] = (0, 0) # (the earliest timestamp, gas)
        self.MasterRollbackUnconfirm: Dict[str, Tuple[int, int]] = {} # fromChainId -> (the earliest timestamp, gas)
        self.MasterRollbackConfirm: Tuple[int, int] = (0, 0) # (the earliest timestamp, gas)
        self.WorkerCommitUnconfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.WorkerCommitConfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.WorkerRollbackUnconfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.WorkerRollbackConfirm: Dict[str, Tuple[int, int]] = {} # currentChainId -> (the earliest timestamp, gas)
        self.latencyMetric = PhaseLatencyMetric()
        self.gasMetric = GasMetric()
    
    def calc_EachPhaseLatency(self):
        self.latencyMetric.m_issue_2_w_ucf_p = avg(value[0]-self.MasterIssue[0] for value in self.WorkerPrepareUnconfirm.values() if int(value[0]) != 0)
        self.latencyMetric.w_ucf_p_2_m_cf_recv = avg(self.MasterReceiveConfirm[key][0]-value[0] for key, value in self.WorkerPrepareUnconfirm.items() if int(value[0]) != 0 and key in self.MasterReceiveConfirm)
        self.latencyMetric.w_ucf_p_2_m_cf_c = avg(self.MasterCommitConfirm[0]-value[0] for value in self.WorkerPrepareUnconfirm.values() if int(value[0]) != 0 and self.MasterCommitConfirm[0] != 0)
        self.latencyMetric.w_ucf_p_2_m_cf_r = avg(self.MasterRollbackConfirm[0]-value[0] for value in self.WorkerPrepareUnconfirm.values() if int(value[0]) != 0 and self.MasterRollbackConfirm[0] != 0)
        self.latencyMetric.m_cf_c_2_w_cf_c = avg(value[0] - self.MasterCommitConfirm[0] for value in self.WorkerCommitConfirm.values() if int(value[0]) != 0 and self.MasterCommitConfirm[0] != 0)
        self.latencyMetric.m_cf_r_2_w_cf_r = avg(value[0] - self.MasterRollbackConfirm[0] for value in self.WorkerRollbackConfirm.values() if int(value[0]) != 0 and self.MasterRollbackConfirm[0] != 0)

    def calc_TotalLatency(self):
        final = 0
        a = fmt2f(self.__avg_timestamp(self.WorkerCommitConfirm))
        b = fmt2f(self.__avg_timestamp(self.WorkerRollbackConfirm))
        if (a == 0.00 and b == 0.00) or (a != 0.00 and b != 0.00):
            raise ValueError("invalid final latency")
        if a == 0.00:
            final = b
        else:
            final = a
        self.latencyMetric.overall = final-float(self.MasterIssue[0])

    def __avg_timestamp(self, obj: Dict[str, Tuple[int, int]]):
        if len(obj) == 0:
            return 0
        return sum([x[0] for x in obj.values()]) / len(obj)
    
    def calc_gas(self):
        self.gasMetric.m_issue = self.MasterIssue[1]
        self.gasMetric.w_ucf_p = avg(value[1] for value in self.WorkerPrepareUnconfirm.values() if int(value[1]) != 0)
        self.gasMetric.w_cf_p = avg(value[1] for value in self.WorkerPrepareConfirm.values() if int(value[1]) != 0)
        self.gasMetric.m_cf_recv = avg(value[1] for value in self.MasterReceiveConfirm.values() if int(value[1] != 0))
        self.gasMetric.m_cf_c = self.MasterCommitConfirm[1]
        self.gasMetric.m_cf_r = self.MasterRollbackConfirm[1]
        self.gasMetric.w_cf_c = avg(value[1] for value in self.WorkerCommitConfirm.values() if int(value[1] != 0))
        self.gasMetric.w_cf_r = avg(value[1] for value in self.WorkerRollbackConfirm.values() if int(value[1] != 0))

    def update_metric(self, metric: Metric):
        if metric.forged:
            return
        if metric.phase == Metric.MasterIssue:
            self.MasterIssue = (metric.timestamp, metric.gas)
            return
        elif metric.phase == Metric.WorkerPrepareUnconfirm:
            if metric.chainId not in self.WorkerPrepareUnconfirm:
                self.WorkerPrepareUnconfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerPrepareUnconfirm[metric.chainId] = cmp(self.WorkerPrepareUnconfirm[metric.chainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.WorkerPrepareConfirm:
            if metric.chainId not in self.WorkerPrepareConfirm:
                self.WorkerPrepareConfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerPrepareConfirm[metric.chainId] = cmp(self.WorkerPrepareConfirm[metric.chainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.MasterReceiveUnconfirm:
            if metric.fromChainId not in self.MasterReceiveUnconfirm:
                self.MasterReceiveUnconfirm[metric.fromChainId] = (metric.timestamp, metric.gas)
            else:
                self.MasterReceiveUnconfirm[metric.fromChainId] = cmp(self.MasterReceiveUnconfirm[metric.fromChainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.MasterReceiveConfirm:
            if metric.fromChainId not in self.MasterReceiveConfirm:
                self.MasterReceiveConfirm[metric.fromChainId] = (metric.timestamp, metric.gas)
            else:
                self.MasterReceiveConfirm[metric.fromChainId] = cmp(self.MasterReceiveConfirm[metric.fromChainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.MasterCommitUnconfirm:
            if metric.fromChainId not in self.MasterCommitUnconfirm:
                self.MasterCommitUnconfirm[metric.fromChainId] = (metric.timestamp, metric.gas)
            else:
                self.MasterCommitUnconfirm[metric.fromChainId] = cmp(self.MasterCommitUnconfirm[metric.fromChainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.MasterCommitConfirm:
            self.MasterCommitConfirm = (metric.timestamp, metric.gas)
        elif metric.phase == Metric.MasterRollbackUnconfirm:
            if metric.fromChainId not in self.MasterRollbackUnconfirm:
                self.MasterRollbackUnconfirm[metric.fromChainId] = (metric.timestamp, metric.gas)
            else:
                self.MasterRollbackUnconfirm[metric.fromChainId] = cmp(self.MasterRollbackUnconfirm[metric.fromChainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.MasterRollbackConfirm:
            self.MasterRollbackConfirm = (metric.timestamp, metric.gas)
        elif metric.phase == Metric.WorkerCommitUnconfirm:
            if metric.chainId not in self.WorkerCommitUnconfirm:
                self.WorkerCommitUnconfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerCommitUnconfirm[metric.chainId] = cmp(self.WorkerCommitUnconfirm[metric.chainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.WorkerCommitConfirm:
            if metric.chainId not in self.WorkerCommitConfirm:
                self.WorkerCommitConfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerCommitConfirm[metric.chainId] = cmp(self.WorkerCommitConfirm[metric.chainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.WorkerRollbackUnconfirm:
            if metric.chainId not in self.WorkerRollbackUnconfirm:
                self.WorkerRollbackUnconfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerRollbackUnconfirm[metric.chainId] = cmp(self.WorkerRollbackUnconfirm[metric.chainId], (metric.timestamp, metric.gas))
        elif metric.phase == Metric.WorkerRollbackConfirm:
            if metric.chainId not in self.WorkerRollbackConfirm:
                self.WorkerRollbackConfirm[metric.chainId] = (metric.timestamp, metric.gas)
            else:
                self.WorkerRollbackConfirm[metric.chainId] = cmp(self.WorkerRollbackConfirm[metric.chainId], (metric.timestamp, metric.gas))
        else:
            raise ValueError(f"Unknown phase: {metric.phase}")

    def to_dict(self):
        return {
            'latency': {
                'overall': self.latencyMetric.overall,
                'm_issue_2_w_ucf_p': self.latencyMetric.m_issue_2_w_ucf_p,
                'w_ucf_p_2_m_cf_recv': self.latencyMetric.w_ucf_p_2_m_cf_recv,
                'w_ucf_p_2_m_cf_c': self.latencyMetric.w_ucf_p_2_m_cf_c,
                'w_ucf_p_2_m_cf_r': self.latencyMetric.w_ucf_p_2_m_cf_r,
                'm_cf_c_2_w_cf_c': self.latencyMetric.m_cf_c_2_w_cf_c,
                'm_cf_r_2_w_cf_r': self.latencyMetric.m_cf_r_2_w_cf_r,
            },
            'gas': {
                'm_issue': self.gasMetric.m_issue,
                'w_ucf_p': self.gasMetric.w_ucf_p,
                'w_cf_p': self.gasMetric.w_cf_p,
                'm_cf_recv': self.gasMetric.m_cf_recv,
                'm_cf_c': self.gasMetric.m_cf_c,
                'm_cf_r': self.gasMetric.m_cf_r,
                'w_cf_c': self.gasMetric.w_cf_c,
                'w_cf_r': self.gasMetric.w_cf_r,
            }
        }

class Data:
    def __init__(self, path):
        self.path = path
        self.data: Dict[str, Dict[str, Dict[str, List[Dict]]]] = json.load(open(path, 'r'))
        self.transactionMetrics: Dict[str, TransactionMetric] = {} # txhash -> TransactionMetric

    def get_data(self):
        return self.data
    
    def filter(self):
        # for each transaction hash (TransactionMetric), update its each metric
        for txhash, item in self.data.items():
            for chainId, metricsInPhase in item.items():
                for phase, metrics in metricsInPhase.items():
                    for metric in metrics:
                        m = Metric(metric)
                        if txhash not in self.transactionMetrics:
                            self.transactionMetrics[txhash] = TransactionMetric()
                        self.transactionMetrics[txhash].update_metric(m)
    
    def deal(self):
        for txhash, transactionMetric in self.transactionMetrics.items():
            transactionMetric.calc_EachPhaseLatency()
            transactionMetric.calc_TotalLatency()
            transactionMetric.calc_gas()
    
    def save(self):
        d = {key: value.to_dict() for key, value in self.transactionMetrics.items()}
        json.dump(d, open("transactionMetrics.json", "w"))

if __name__ == "__main__":
    data = Data("observation/test.json")
    data.filter()
    data.deal()
    data.save()
