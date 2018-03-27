public class XRomanRule implements RomanRule {
    @Override
    public boolean check(int number) {
        return number >= 10;
    }

    @Override
    public String number() {
        return "X";
    }

    @Override
    public int decNumber() {
        return 10;
    }
}
