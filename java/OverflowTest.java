public class OverflowTest {
    public static void main(String[] args) {
        int max_int32 = (1 << 31) - 1;

        System.out.println(max_int32);
        max_int32 += 1;
        System.out.println(max_int32);

        long max_int64 = (1 << 63) - 1;
        System.out.println(max_int64);
        System.out.println((1 << 63) - 1);
    }
}
