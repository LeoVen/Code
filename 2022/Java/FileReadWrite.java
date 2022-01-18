import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class FileReadWrite {
    public static void main(String[] args) throws IOException {
        var path = Paths.get("FileReadWrite.java");
        var result = Files.readString(path);

        System.out.println(result);

        Files.writeString(path, result + "\n// Hello There");
    }
}

// Hello There
// Hello There
// Hello There