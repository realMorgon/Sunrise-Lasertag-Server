import time

from flask import Flask, request

from data.device import Device
from data.player import Player

app = Flask(__name__)

devices = {}
players = {}
startTime = round(time.time() * 1000)

print(f"Started Server at {startTime} milliseconds after boot")


def calculate_time(client_id, client_time):
    return devices.get(client_id).get_time() + client_time


@app.route("/api/test")
def api_test():
    print("Anfrage erhalten!")
    return ("Anfrage erhalten!", 200)


@app.post("/api/sync")
def sync_device():
    client_time = int(request.args.get("time"))
    client_id = str(len(devices))
    client_type = request.args.get("type")

    device = Device(client_time, client_type)

    devices.update({client_id: device})
    print(devices)

    if client_type == "Vest":
        players.update({client_id: Player(1)})
        print(players)

    return (client_id, 200)


@app.get("/api/data")
def get_data():
    device_id = request.args.get("id")

    device_time = int(devices.get(device_id).get_time())
    device_type = devices.get(device_id).get_type()

    response = f"ID: {device_id} </n>Time: {device_time} </n>Type: {device_type}"

    return (response, 200)


@app.post("/api/hit/send")
def send_hit():
    client_id = request.args.get("id")
    client_time = int(request.args.get("time"))

    gametime = calculate_time(client_id, client_time)
    print(f"ID: {client_id} schießt bei {gametime} mills")

    response = f"Du hast ID: {client_id} und TIME: {client_time} (Clienttime) als Schuss gesendet"
    return (response, 200)


@app.post("/api/hit/recieve")
def recieve_hit():
    client_id = request.args.get("my_id")
    hitter_id = request.args.get("their_id")
    client_time = int(request.args.get("time"))

    gametime = calculate_time(client_id, client_time)
    print(f"ID: {client_id} wurde von ID: {hitter_id} bei {gametime} mills getroffen")

    players.get(client_id).deal_damage(1)

    response = f"Du hast EIGENE_ID: {client_id}, FREMDE_ID: {hitter_id} und TIME: {client_time} (Clienttime) als Treffer gesendet"
    return (response, 200)
