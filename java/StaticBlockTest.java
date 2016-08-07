/**
 * Test java static block
 * @author wdeqin 
 * @version 1.0
 */

class StaticBlock {
    public static int cnt = 0;
    static {
        cnt++; 
        System.out.println("This is static block of java class!");
    }
}

public class StaticBlockTest {
    public static void main(String[] argv) {
        System.out.println("Create StaticBlock object 1!");
        StaticBlock sb1 = new StaticBlock();
        System.out.println("Create StaticBlock object 2!");
        StaticBlock sb2 = new StaticBlock();
        if (StaticBlock.cnt == 2) {
            System.out.println("Static block exectues every creation of obj!");
        } else {
            System.out.println("Static block exectues only once!");
        }
    }
}

