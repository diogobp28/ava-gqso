import java.util.*;

public class soma_calc {
    
    public static void main(String args[]) {

        double num1, num2;

        Scanner opcao = new Scanner(System.in);
        
        System.out.println("Digite os números a serem somados:\n");
        
        num1 = opcao.nextDouble();
        System.out.println("+");
        num2 = opcao.nextDouble();
            
        double soma = num1+num2;
        
        System.out.println("="+soma);    
    }
}