import java.util.ArrayList;
import java.util.List;

public class CopyOf {
    public static void main(String[] args) {
        List<String> names = new ArrayList<>();

        names.add("Isaac");
        names.add("Albert");
        names.add("Marie");

        var result = List.copyOf(names);

        for (String name : result) {
            System.out.println(name);
        }
    }
}
