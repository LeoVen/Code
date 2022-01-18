class Equals {
    public static void main(String[] args) {
        var one = new String("Hello, World");
        var two = new String("Hello, World");

        if (one == two) {
            System.out.println("One == two");
        } else if (one.equals(two)) {
            System.out.println("one.Equals(two)");
        }
    }
}
