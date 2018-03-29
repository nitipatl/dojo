public class StringCalculator {
    public int display(String input) {
        AddRule addRule = new AddRule();
        EmptyRule emptyRule = new EmptyRule();
        OtherRule otherRule = new OtherRule();
        if (emptyRule.check(input)) {
            return emptyRule.result();
        }
        if (addRule.check(input)) {
            return addRule.result(input);
        }
        return otherRule.result(input);
    }
}
