/*
 * File: InterfaceDefaultTest.java
 * Author: wdeqin
 * Time: 2015/5/23 14:40:52
 * Purpose: Java 8 Default interface method test
 */

public class InterfaceDefaultTest {
    public static void main(String[] args) {
        Formula c = new Calculator();
        System.out.println(c.sqr(c.sqrt(19)));
    }
}

interface Formula {
    double sqr(double e);
    default double sqrt(double e) {
        return Math.sqrt(e);
    }
}

class Calculator implements Formula {
    @Override
    public double sqr(double e) {
        return e * e;
    }
}
