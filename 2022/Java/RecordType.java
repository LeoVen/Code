public class RecordType {
    record SimpleData(String coord, String address, int number) {
        SimpleData {
            // Constructor
        }

        public String getFoo() {
            // Method
            return String.format("%s %s %d", coord, address, number);
        }
    }

    public static void main(String[] args) {
        var data1 = new SimpleData("1122", "Addler St.", 222);
        var data2 = new SimpleData("1122", "Addler St.", 222);
        var data3 = new SimpleData("1122", "Hoover St.", 222);
        var data4 = data3;

        System.out.println(data1 == data2);
        System.out.println(data2 == data3);
        System.out.println(data3 == data4);

        System.out.println(data1.equals(data2));
        System.out.println(data2.equals(data3));

        System.out.println(data1.getFoo());
    }
}
