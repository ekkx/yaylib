from datetime import datetime


def parse_datetime(timestamp: int) -> str:
    return str(datetime.fromtimestamp(timestamp))
