/*
 * Test Java 8 lamdba feature
 * @author wdeqin
 * @version 1.0
 */
public class LambdaTest {
    /**
     * Program entry point
     * @param argv arguements vector from cmd
     */
    public static void main(String[] argv) {
        delagate(e -> System.out.println(e * e));
    }

    /** Use lambda function to process element 
     * @param op function takes an int and return nothing
     */
    public static void delagate(Operator op) {
        for (int i = 0; i <= 100; i += 17)
            op.operate(i);
    }
}

/** Interface type to match lambda function */
interface Operator {
    void operate(int e);
}
