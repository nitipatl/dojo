package com.company;

import org.junit.Before;
import org.junit.Test;

import java.util.ArrayList;

import static org.junit.Assert.assertEquals;

public class FizzBuzzTest {
    FizzBuzz fizzBuzz;
    @Before
    public void SetUp() {
        Rule fizzRule = new FizzRule();
        Rule buzzRule = new BuzzRule();
        Rule fizzBuzzRule = new FizzBuzzRule(fizzRule, buzzRule);

        ArrayList<Rule> rules = new ArrayList<>();
        rules.add(fizzBuzzRule);
        rules.add(fizzRule);
        rules.add(buzzRule);

        fizzBuzz = new FizzBuzz(rules);
    }
    @Test
    public void Input1output1() {
        assertEquals("1", fizzBuzz.display(1));
    }
    @Test
    public void Input2output2() {
        assertEquals("2", fizzBuzz.display(2));
    }
    @Test
    public void Input3outputFizz() {
        assertEquals("Fizz", fizzBuzz.display(3));
    }
    @Test
    public void Input5outputBuzz() {
        assertEquals("Buzz", fizzBuzz.display(5));
    }
    @Test
    public void Input6outputFizz() {
        assertEquals("Fizz", fizzBuzz.display(6));
    }
    @Test
    public void Input15outputFizzBuzz() {
        assertEquals("FizzBuzz", fizzBuzz.display(15));
    }

}
