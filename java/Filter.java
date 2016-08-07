import java.util.List;
import java.util.LinkedList;
import java.util.ArrayList;

/**
 * Java 8 filter accept function as parameter
 * @author wdeqin
 * @version 1.0
 */
public class Filter {
    /**
     * Filter
     * @param l List to filt
     * @param p Predicate item in list is accepted or not
     * @return A list after filt
     */
    public static<T> List<T> filter(List<T> l, Predicate<T> p) {
        LinkedList<T> r = new LinkedList<T>();
        for (T t: l) {
            if (p.accept(t)) {
                r.add(t);
            }
        }
        return r;
    }

    interface Predicate<T> {
        boolean accept(T t);
    }

    /**
     * Program entry point
     * @param argv arguements vector from cmd
     */
    public static void main(String[] argv) {
        ArrayList<Integer> a = new ArrayList<Integer>(10);
        for (int i = 0; i != 10; ++i) {
            a.add(i);
        }

        System.out.println(filter(a, t -> t % 2 == 0));
    }

}
