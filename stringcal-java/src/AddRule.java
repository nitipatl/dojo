public class AddRule {
    public int result(String input) {
        String[] out = input.split(",");
        int sum = 0;
        if (out.length > 0) {
            for (String number: out) {
                sum += Integer.valueOf(number);
            }
        }
        return sum;
    }

    public boolean check(String input) {
        return input.contains(",");
    }
}
