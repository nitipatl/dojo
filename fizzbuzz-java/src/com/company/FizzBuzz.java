package com.company;

import java.util.ArrayList;

public class FizzBuzz {
    ArrayList<Rule> rules;
    public FizzBuzz(ArrayList<Rule> rules) {
        this.rules = rules;
    }
    public String display(int number) {
        for (Rule rule: rules) {
            if (rule.check(number)) {
                return rule.display();
            }
        }
        return String.valueOf(number);
    }
}
