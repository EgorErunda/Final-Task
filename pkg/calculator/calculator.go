package calculator

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func stringToFloat64(str string) float64 {
// 	degree := float64(1)
// 	var res float64 = 0
// 	var invers bool = false
// 	for i := len(str); i > 0; i-- {
// 		if str[i-1] == '-' {
// 			invers = true
// 		} else {
// 			res += float64(9-int('9'-str[i-1])) * degree
// 			degree *= 10
// 		}
// 	}
// 	if invers {
// 		res = 0 - res
// 	}
// 	return res
// }

// func isSign(value rune) bool {
// 	return value == '+' || value == '-' || value == '*' || value == '/'
// }

// func Calc(expression string) (float64, error) {
// 	if len(expression) < 3 {
// 		return 0, ErrInvalidExpression
// 	}
// 	//////////////////////////////////////////////////////////////////////////////////////////////////////
// 	var res float64
// 	var b string
// 	var c rune = 0
// 	var resflag bool = false
// 	var isc int
// 	var countc int = 0
// 	//////////////////////////////////////////////////////////////////////////////////////////////////////
// 	for _, value := range expression {
// 		if isSign(value) {
// 			countc++
// 		}
// 	}
// 	//////////////////////////////////////////////////////////////////////////////////////////////////////
// 	if isSign(rune(expression[0])) || isSign(rune(expression[len(expression)-1])) {
// 		return 0, ErrInvalidExpression
// 	}
// 	for i, value := range expression {
// 		if value == '(' {
// 			isc = i
// 		}
// 		if value == ')' {
// 			calc, err := Calc(expression[isc+1 : i])
// 			if err != nil {
// 				return 0, ErrInvalidExpression
// 			}
// 			calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
// 			i2 := i
// 			i -= len(expression[isc:i+1]) - len(calcstr)
// 			expression = strings.Replace(expression, expression[isc:i2+1], calcstr, 1) // Меняем скобки на результат выражения в них
// 		}
// 	}
// 	if countc > 1 {
// 		for i := 1; i < len(expression); i++ {
// 			value := rune(expression[i])

// 			//Умножение и деление
// 			if value == '*' || value == '/' {
// 				var imin int = i - 1
// 				if imin != 0 {
// 					for !isSign(rune(expression[imin])) && imin > 0 {
// 						imin--
// 					}
// 					imin++
// 				}
// 				var imax int = i + 1
// 				if imax == len(expression) {
// 					imax--
// 				} else {
// 					for !isSign(rune(expression[imax])) && imax < len(expression)-1 {
// 						imax++
// 					}
// 				}
// 				if imax == len(expression)-1 {
// 					imax++
// 				}
// 				calc, err := Calc(expression[imin:imax])
// 				if err != nil {
// 					return 0, ErrInvalidExpression
// 				}
// 				calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
// 				i -= len(expression[isc:i+1]) - len(calcstr) - 1
// 				expression = strings.Replace(expression, expression[imin:imax], calcstr, 1) // Меняем скобки на результат выражения в них
// 			}
// 			if value == '+' || value == '-' || value == '*' || value == '/' {
// 				c = value
// 			}
// 		}
// 	}
// 	//////////////////////////////////////////////////////////////////////////////////////////////////////
// 	for _, value := range expression + "s" {
// 		switch {
// 		case value == ' ':
// 			continue
// 		case value > 47 && value < 58: // Если это цифра
// 			b += string(value)
// 		case isSign(value) || value == 's': // Если это знак
// 			if resflag {
// 				switch c {
// 				case '+':
// 					res += stringToFloat64(b)
// 				case '-':
// 					res -= stringToFloat64(b)
// 				case '*':
// 					res *= stringToFloat64(b)
// 				case '/':
// 					res /= stringToFloat64(b)
// 				}
// 			} else {
// 				resflag = true
// 				res = stringToFloat64(b)
// 			}
// 			b = strings.ReplaceAll(b, b, "")
// 			c = value

// 			/////////////////////////////////////////////////////////////////////////////////////////////
// 		case value == 's':
// 		default:
// 			return 0, fmt.Errorf("not correct input")
// 		}
// 	}
// 	return res, nil
// }
import (
	"errors"
	"strconv"
	"unicode"
)

// вычисляет значение строкового выражения
func Calc(expression string) (float64, error) {
	expression = removeSpaces(expression)
	if !isBalanced(expression) {
		return 0, ErrInvalidExpression
	}
	rpn, err := shuntingYard(expression)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// удаляет все пробелы из строки
func removeSpaces(s string) string {
	var result string
	for _, char := range s {
		if !unicode.IsSpace(char) {
			result += string(char)
		}
	}
	return result
}

// проверяет кол-во скобок в строке
func isBalanced(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// преобразование выражения в обратную польскую нотацию
func shuntingYard(expression string) ([]interface{}, error) {
	var output []interface{}
	var operators []rune

	for i := 0; i < len(expression); {
		char := rune(expression[i])

		if unicode.IsDigit(char) || char == '.' {
			var number string
			for ; i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.'); i++ {
				number += string(expression[i])
			}
			num, err := strconv.ParseFloat(number, 64)
			if err != nil {
				return nil, ErrInvalidExpression
			}
			output = append(output, num)
		} else if isOperator(char) {
			for len(operators) > 0 && isOperator(operators[len(operators)-1]) &&
				precedence(operators[len(operators)-1]) >= precedence(char) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, char)
			i++
		} else if char == '(' {
			operators = append(operators, char)
			i++
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				return nil, ErrInvalidExpression
			}
			operators = operators[:len(operators)-1]
			i++
		} else {
			return nil, ErrInvalidExpression
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == '(' || operators[len(operators)-1] == ')' {
			return nil, ErrInvalidExpression
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func precedence(char rune) int {
	switch char {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func evaluateRPN(rpn []interface{}) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		switch token := token.(type) {
		case float64:
			stack = append(stack, token)
		case rune:
			if len(stack) < 2 {
				return 0, ErrInvalidExpression
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case '+':
				result = a + b
			case '-':
				result = a - b
			case '*':
				result = a * b
			case '/':
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			default:
				return 0, ErrInvalidExpression
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}

	return stack[0], nil
}
