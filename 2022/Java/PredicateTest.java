import java.util.List;
import java.util.function.Predicate;

public class PredicateTest {
    public static void main(String[] args) {
        var numbers = List.of(1, 2, 3, 4, 5, 6, 7, 8);

        Predicate<Integer> pred = number -> number % 2 == 0;

        numbers.stream().filter(pred).forEach(System.out::println);
        numbers.stream().filter(Predicate.not(pred)).forEach(System.out::println);
    }
}
