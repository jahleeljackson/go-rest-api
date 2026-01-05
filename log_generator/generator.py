from datetime import datetime, timedelta
from zoneinfo import ZoneInfo 
from ipaddress import IPv4Address
import random
import httpx


NUM_CLIENTS = 1000 

logs = ""

tz = ZoneInfo("UTC")
now = datetime.now(tz)

for i in range(NUM_CLIENTS):
    rand_delta = timedelta(days=random.randint(0, 3), hours=random.randint(0, 24), minutes=random.randint(0, 60))
    now += rand_delta
    id_num = random.randrange(100_000, 1_000_000)
    ip = IPv4Address(id_num)
    logs += f"{now} {id_num} {ip}\n"




data = {
    "title": "log_1",
    "content": logs
}

# Add logs
# response = httpx.post(url="http://localhost:8080/logs", json=data)

# Get all logs
# response = httpx.get(url="http://localhost:8080/logs")

# Get log by id
# response = httpx.get(url="http://localhost:8080/logs/1")

# print(response.text)