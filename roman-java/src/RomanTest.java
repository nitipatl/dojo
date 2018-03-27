import org.junit.Before;
import org.junit.Test;

import java.util.ArrayList;

import static org.junit.Assert.assertEquals;

public class RomanTest {
    Roman roman;
    @Before
    public void Setup() {
        RomanRule iRomanRule = new IRomanRule();
        RomanRule vRomanRule = new VRomanRule();
        RomanRule xRomanRule = new XRomanRule();
        ArrayList<RomanRule> rules = new ArrayList<>();
        rules.add(xRomanRule);
        rules.add(vRomanRule);
        rules.add(iRomanRule);
        roman = new Roman(rules);
    }

    @Test
    public void input1displayI() {
        assertEquals("I", roman.display(1));
    }

    @Test
    public void input2displayII() {
        assertEquals("II", roman.display(2));
    }

    @Test
    public void input3displayIII() {
        assertEquals("III", roman.display(3));
    }

    @Test
    public void input4displayIIII() {
        assertEquals("IIII", roman.display(4));
    }

    @Test
    public void input5displayV() {
        assertEquals("V", roman.display(5));
    }

    @Test
    public void input6displayVI() {
        assertEquals("VI", roman.display(6));
    }

    @Test
    public void input7displayVII() {
        assertEquals("VII", roman.display(7));
    }

    @Test
    public void input8displayVIII() {
        assertEquals("VIII", roman.display(8));
    }

    @Test
    public void input9displayVIIII() {
        assertEquals("VIIII", roman.display(9));
    }

    @Test
    public void input10displayX() {
        assertEquals("X", roman.display(10));
    }

    @Test
    public void input11displayXI() {
        assertEquals("XI", roman.display(11));
    }

}
