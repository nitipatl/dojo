public class IRomanRule implements RomanRule {
    private int _number;
    public boolean check(int number) {
        this._number = number;
        return number > 0;
    }

    public String number() {
        String Ioutput = "";
        for (int i = 0; i < _number; i++) {
            Ioutput += "I";
        }
        return Ioutput;
    }

    public int decNumber() {
        return 0;
    }
}
