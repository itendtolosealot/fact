package fact
import "testing"

func TestFactnegative(t *testing.T){
	number:= -10
    err_got, _ := Factorial(number)
    if err_got == nil {
        t.Errorf("Negative number %d did not get validation error", number)
    }
}

func TestFactzero(t *testing.T){
	number:= 0
    err_got, val := Factorial(number)
    if err_got != nil || val != 1{
        t.Errorf("Exoected error == nil and value = 1. Got err: %v, val: %d", err_got, val)
    }
}

func TestFactLarge(t *testing.T){
	number:= 1000
    err_got, _:= Factorial(number)
    if err_got == nil {
        t.Errorf("Exoected validation error, but got err: %v ", err_got)
    }
}

func TestFactCorrect(t *testing.T){
	number:= 10
    err_got, val:= Factorial(number)
    if err_got != nil || val != 3628800{
        t.Errorf("Exoected validation error, but got err: %v ", err_got)
    }
}