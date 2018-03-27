package com.company;

public class FizzRule implements Rule {
    public boolean check(int number) {
        return number % 3 == 0;
    }

    public String display() {
        return "Fizz";
    }
}

