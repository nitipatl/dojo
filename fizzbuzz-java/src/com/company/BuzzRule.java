package com.company;

public class BuzzRule implements Rule {
    public boolean check(int number) {
        return number % 5 == 0;
    }

    public String display() {
        return "Buzz";
    }
}

