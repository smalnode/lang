import math

class Shape():
    numOfInstance = 0

    @classmethod
    def __init__(self):
        Shape.numOfInstance += 1

    @classmethod
    def getPerimeter(self):
        return -1
    
    @classmethod
    def getProportion(self):
        return -1

    @classmethod
    def getVolume(self):
        return -1

    @staticmethod
    def getNumOfInstance():
        return Shape.numOfInstance

class Rectangle(Shape):
    @classmethod
    def __init__(self, length, width):
        Shape.__init__()
        self.length = length
        self.width = width

    @classmethod
    def getPerimeter(self):
        return 2 * (self.length + self.width)

    @classmethod
    def getProportion(self):
        return self.length * self.width

class Circle(Shape):
    @classmethod
    def __init__(self, radius):
        Shape.__init__()
        self.radius = radius

    @classmethod
    def getPerimeter(self):
        return 2 * math.pi * self.radius

    @classmethod
    def getProportion(self):
        return math.pi * self.radius * self.radius
        

def main():
    rec = Rectangle(10, 20)
    cir = Circle(math.pi)
    print "rec.getPerimeter() = ", rec.getPerimeter()
    print "rec.getPorportion() = ", rec.getProportion()
    print "rec.getVolume() = ", rec.getVolume()
    print "type(rec) = ", type(rec)
    print "type(Rectangle) = ", type(Rectangle)
    print "isinstance(rec, Rectangle) = ", isinstance(rec, Rectangle)
    print "isinstance(rec, Shape) = ", isinstance(rec, Shape)
    print "issubclass(Rectangle, Shape) = ", issubclass(Rectangle, Shape)
    print "issubclass(Shape, Shape) = ", issubclass(Shape, Shape)
    print "issubclass(Shape, Rectangle) = ", issubclass(Shape, Rectangle)
    print "Shape.getNumOfInstance() = ", Shape.getNumOfInstance()
    print "cir.getProportion() = ", cir.getProportion()

if __name__ == "__main__":
    main()
