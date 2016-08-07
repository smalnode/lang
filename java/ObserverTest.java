/**
 * Test Observer design pattern
 * @author wdeqin
 * @version 1.0
 */
import java.util.LinkedList;

public class ObserverTest {
    /**
     * Program entry point
     * @param argv arugements vector from cmd
     */
    public static void main(String[] argv) {
        WeatherStation ws = new WeatherStation();
        ws.setTemperature(17.5);
        PhoneWeatherForecast p = new PhoneWeatherForecast();
        TVWeatherForecast t = new TVWeatherForecast();
        ws.addObserver(p);
        ws.addObserver(t);
        ws.setTemperature(18.5);
        ws.setTemperature(18.5);
        ws.setTemperature(27.5);
        ws.setTemperature(11.5);
        ws.setTemperature(11.5);
    }
}

/**
 * Obserse interface
 * @author wdeqin
 * @version 1.0
 */
interface Observer {
    void temperatureChanged(double temperature);
}

/**
 * Object that was observed
 * @author wdeqin
 * @version 1.0
 */
class WeatherStation {
    private double __temperature;
    private LinkedList<Observer> __observers;

    public WeatherStation() {
        this.__temperature = 0.0;
        this.__observers = new LinkedList<Observer>();
    }

    public void setTemperature(double temperature) {
        if (temperature == this.__temperature) return;
        this.__temperature = temperature;
        for (Observer ob: this.__observers) {
            ob.temperatureChanged(this.__temperature);
        }
    }

    public void addObserver(Observer ob) {
        this.__observers.addFirst(ob);
    }

}

/**
 * Observer phone
 * @author wdeqin
 * @version 1.0
 */
class PhoneWeatherForecast implements Observer {
    private double __temperature;

    /** Phone observer with temperature zero */
    public PhoneWeatherForecast() {
        this.__temperature = 0.0;
    }

    /** Method that object with call when it changed */
    public void temperatureChanged(double temperature) {
        this.__temperature = temperature;
        System.out.println("Phone Wehater Forecast, temperature: " 
                + this.__temperature);
    }
}

/**
 * Observer tv
 * @author wdeqin
 * @version 1.0
 */
class TVWeatherForecast implements Observer {
    private double __temperature;

    /** Tv observer with temperature zero */
    public TVWeatherForecast() {
        this.__temperature = 0.0;
    }

    /** Method that object with call when it changed */
    public void temperatureChanged(double temperature) {
        this.__temperature = temperature;
        if (this.__temperature < 15.0) {
            System.out.println("TV Wehater Forecast, low temperature, : " 
                    + this.__temperature);
        }
    }
}
