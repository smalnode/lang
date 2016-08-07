/**
 * Test java non-abstract class vs abstract class
 * @author 
 * @version 1.0
 */
public class NonAbstractClassTest {
    public static void main(String[] argv) {
        GeometricObject[] geos = new GeometricObject[4];
        geos[0] = new Circle();
        geos[1] = new Rectangle();
        geos[2] = new Rectangle("black", true, 2, 4);
        geos[3] = new Circle("blue", false, 10.0);

        for(GeometricObject geo : geos) {
            System.out.println(geo + "\n area: " + geo.getArea() 
                    + "\n perimeter: " + geo.getPerimeter()
                    + "\n class: " + geo.getClass());
        }
    }
}

class GeometricObject {
    private String color;
    private boolean filled;
    private java.util.Date dateCreated;

    protected GeometricObject() {
        this.color = "white";
        this.filled = false;
        this.dateCreated = new java.util.Date();
    }

    protected GeometricObject(String color, boolean filled) {
        this.color = color;
        this.filled = filled;
        this.dateCreated = new java.util.Date();
    }

    public String getColor() {
        return this.color;
    }

    public void setColor(String color) {
        this.color = color;
    }

    public boolean getFilled() {
        return this.filled;
    }

    public void setFilled(boolean filled) {
        this.filled = filled;
    }

    public java.util.Date getDateCreated() {
        return this.dateCreated;
    }

    public double getArea() {
        return -1.0;
    }

    public double getPerimeter() {
        return -1.0;
    }

    @Override
    public String toString() {
        return "created on " + dateCreated + "\ncolor: " + color 
            + " and filled: " + filled;
    }
}

class Circle extends GeometricObject {
    private double radius;

    public Circle() {
        super();
        this.radius = 1.0;
    }

    public Circle(String color, boolean filled, double radius) {
        super(color, filled);
        this.radius = radius;
    }

    public double getRadius() {
        return this.radius;
    }

    public void setRadius(double radius) {
        this.radius = radius;
    }

    public double getArea() {
        return Math.PI * radius * radius;
    }

    public double getPerimeter() {
        return 2.0 * Math.PI * radius;
    }
}

class Rectangle extends GeometricObject {
    private double length;
    private double height;

    public Rectangle() {
        super();
        this.length = 1.0;
        this.height = 1.0;
    }

    public Rectangle(String color, boolean filled, 
            double length, double height) {
        super(color, filled);
        this.length = length;
        this.height = height;
    }

    public double getLength() {
        return this.length;
    }

    public void setLength(double length) {
        this.length = length;
    }

    public double getHeight() {
        return this.height;
    }

    public void setHeight(double height) {
        this.height = height;
    }

    public double getArea() {
        return length * height;
    }

    public double getPerimeter() {
        return 2 * (length + height);
    }
}
