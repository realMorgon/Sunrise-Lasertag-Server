class Device:

    def __init__ (self, time, type):
        self.time = time
        self.type = type
        print(f"Created new device with time {time} and type {type}")

    def get_time(self):
        return self.time

    def get_type(self):
        return self.type
