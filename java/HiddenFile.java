import java.io.File;

/**
 * Find hiddent by pass method as a refreence
 * @author wdeqin
 * @version 1.0
 */
public class HiddenFile {
    /**
     * Program entry point
     * @param argv arguements vector from cmd
     */
    public static void main(String[] argv) {
        File[] hiddenFiles = new File(".").listFiles(File::isHidden);
        System.out.println("Hidden files in cwd:");
        if (hiddenFiles.length <= 0) {
            System.out.println("None!!!");
        } else {
            for (File f: hiddenFiles) {
                System.out.println(f);
            }
        }
    }
}
