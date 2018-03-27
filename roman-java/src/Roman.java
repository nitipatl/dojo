import java.util.ArrayList;

public class Roman {
    ArrayList<RomanRule> rules;
    public Roman(ArrayList<RomanRule> rules) {
        this.rules = rules;
    }

    public String display(int number) {

        String output = "";

        for (RomanRule rule: rules) {
            if (rule.check(number)) {
                output += rule.number();
                number -= rule.decNumber();
            }
        }
        return output;
    }


}
