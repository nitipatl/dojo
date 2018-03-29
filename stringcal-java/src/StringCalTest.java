import org.junit.Test;

import static junit.framework.TestCase.assertEquals;

public class StringCalTest {
    @Test
    public void input1return1() {
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("1");
        assertEquals(1, display);
    }

    @Test
    public void input2return2() {
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("2");
        assertEquals(2, display);
    }

    @Test
    public void input1comma2return3() {
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("1,2");
        assertEquals(3, display);
    }

    @Test
    public void input1comma3return4() {
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("1,3");
        assertEquals(4, display);
    }

    @Test
    public void input1comma2comma3return6 (){
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("1,2,3");
        assertEquals(6, display);
    }

    @Test
    public void inputEmptyreturn0 (){
        StringCalculator stringCalulator = new StringCalculator();
        int display = stringCalulator.display("");
        assertEquals(0, display);
    }

}
