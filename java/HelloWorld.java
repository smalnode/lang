/**
 * Java Hello world
 * @author wdeqin
 * @version 1.0
 */
public class HelloWorld {
    /**
     * Program entry point
     * @param argv Arguements vector from cmd
     */
    public static void main(String[] argv) {
        System.out.println("Hello, " + (argv.length > 0 ? argv[0] : "world") + "!");
    }
}
