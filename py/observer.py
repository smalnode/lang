# design pattern: observer

class weatherstat():
    def __init__(self):
        self._temperature = 0
        self._observers = set()

    def addobserver(self, obfunc):
        self._observers.add(obfunc)

    def removeobserver(self, obfunc):
        self._observers.remove(obfunc)

    def settemp(self, temp):
        self._temperature = temp
        for f in self._observers:
            f(self._temperature)

def printtemp(temp):
    print("Temperature changed, current = %.1f" % (temp))

def lowtemp(temp):
    if temp < 16:
        print("Low temperature(%.1f), take care you self" % (temp))

ws = weatherstat()
ws.addobserver(printtemp)
ws.addobserver(lowtemp)

ws.settemp(23)
ws.settemp(27)
ws.removeobserver(printtemp)
ws.settemp(12)

