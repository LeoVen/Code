import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class StringApi {
    public static void main(String[] args) throws IOException {
        System.out.println("   ".isBlank());
        System.out.println("".isEmpty());

        var path = Paths.get("StringApi.java");
        var result = Files.readString(path);

        result.lines().forEach(str -> System.out.println("--line--> " + str));
    }
}
