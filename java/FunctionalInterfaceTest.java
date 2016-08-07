/**
 * Test java 8 Functional Interface
 * @author wdeqin
 * @version 1.0
 */
public class FunctionalInterfaceTest {
    /**
     * Program entry point
     * @param argv arguement vector from cmd
     */
    public static void main(String[] argv) {
        System.out.println(TakeLambda(new Integer(100), e -> e + ""));
    }

    private static String TakeLambda(Integer e, Converter<Integer, String> c) {
        return c.Convert(e);
    }
}

@FunctionalInterface
interface Converter<F, T> {
    T Convert(F e);
}
