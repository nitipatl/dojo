package com.company;

public class FizzBuzzRule implements Rule {
    private Rule fizzRule;
    private Rule buzzRule;
    public FizzBuzzRule(Rule fizzRule, Rule buzzRule) {
        this.fizzRule = fizzRule;
        this.buzzRule = buzzRule;
    }

    public boolean check(int number) {
        return fizzRule.check(number) && buzzRule.check(number);
    }

    public String display() {
        return "FizzBuzz";
    }
}

