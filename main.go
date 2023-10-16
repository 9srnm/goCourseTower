package main

import ("fmt"
		"math")

func main() {
	var ans string
	var a, b, c float64

	fmt.Print("Введите коэффициенты квадратного уравнения через пробел: ")
	fmt.Scanf("%f %f %f", &a, &b, &c)

	D := math.Pow(b, 2) - 4 * a * c

	if a == 0 {
		if D > 0 {
			ans = "У уравнения одно решение: " + fmt.Sprintf("%.2f", -c / b)
		} else if c == 0 {
			ans = "У уравнения бесконечное количество решений"
		} else {
			ans = "У уравнения нет решений"
		}
	} else {
		if D > 0 {
			ans = "У уравнения два решения: " + fmt.Sprintf("%.2f", (-b - math.Sqrt(D)) / (2 * a)) + " и " + fmt.Sprintf("%.2f", (-b + math.Sqrt(D)) / (2 * a))
		} else if D == 0 {
			ans = "У уравнения одно решение: " + fmt.Sprintf("%.2f", -b / (2 * a))
		} else {
			ans = "У уравнения нет решений"
		}
	}

	fmt.Println(ans)
}