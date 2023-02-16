import threading

class TrafficLight:
    def __init__(self):
        self.lock = threading.Lock()
        self.road = 1

    def carArrived(self, carId, roadId, direction, turnGreen, crossCar):
        with self.lock:
            if (self.road != roadId):
                turnGreen()
                self.road = roadId

            crossCar()