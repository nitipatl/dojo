public class VRomanRule implements RomanRule {
    public boolean check(int number) {
        return number >= 5;
    }

    public String number() {
        return "V";
    }

    public int decNumber() {
        return 5;
    }
}
