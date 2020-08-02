# Control roomba with a Logitech F310
# Requires you to be hooked up to the Roomba's UART port
import numpy as np
import queue
import threading

from decimal import Decimal
from inputs import get_gamepad
from pyroombaadapter import PyRoombaAdapter
from time import sleep

q = queue.Queue()

def process_input():
    while True:
        events = get_gamepad()
        for event in events:
            q.put(event)

if __name__ == "__main__":
    q = queue.Queue()
    PORT = "/dev/ttyUSB0"
    adapter = PyRoombaAdapter(PORT)
    adapter.change_mode_to_full()
    max_speed = Decimal(1.000)
    direction = Decimal(0.00)
    back_up = False
    curr_speed = Decimal(0.000)
    spin_up = False
    vacuum = False
    previous_cmd = None
    index = 0
    thread = threading.Thread(target=process_input)
    thread.start()

    while True:
        print(f"Curr_Speed: {curr_speed}, Direction: {direction}, Spin_up: {spin_up}, back_up: {back_up}")
        if spin_up:
            if curr_speed  < max_speed:
                curr_speed += Decimal(.02)
        elif back_up:
            if curr_speed > -max_speed:
                curr_speed -= Decimal(.02)
        else:
            if curr_speed < Decimal(0.00):
                curr_speed += Decimal(.02)
                if curr_speed > Decimal(0.00):
                    curr_speed = Decimal(0.00)
            elif curr_speed > Decimal(0.00):
                curr_speed -= Decimal(.02)
                if curr_speed < Decimal(0.00):
                    curr_speed = Decimal(0.00)

        if not q.empty():
            event = q.get()

            if event.code == "BTN_TOP":
                if event.state == 1:
                    spin_up = True
                else:
                    spin_up = False

            elif event.code == "BTN_TRIGGER":
                if event.state == 1:
                    back_up = True
                elif event.state == 0:
                    back_up = False
            if event.code == "ABS_HAT0X":
                if event.state == 1:
                    direction = -80
                elif event.state == -1:
                    direction = 80
                elif event.state == 0:
                    direction = 0
            if event.code == "BTN_THUMB" and event.state == 1:
                vacuum = not vacuum
                adapter.send_moters_cmd(vacuum, vacuum, vacuum, vacuum, vacuum)
            if event.code == "BTN_THUMB2":
                adapter.move(0, 0)
                thread.join()
                exit(1)
        if [curr_speed, direction] != previous_cmd:
            adapter.move(float(curr_speed), np.deg2rad(round(float(direction),1)))
            previous_cmd = [curr_speed, direction]
