// TestLiteIDE
package main

import (
	"fmt"
	//"golang.org/x/tour/pic"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y int
}

func passage(x int) {
	x += 3
}
func passage_ptr(x *int) {
	*x += 3
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v",
		s, len(x), cap(x), x)
	if x == nil {
		fmt.Println(" nil!")
	} else {
		fmt.Println()
	}
}

func Pic(dx, dy int) [][]uint8 {

	ret := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//ret[y][x] = uint8((x+y)/2)
			ret[y][x] = uint8((x * y))
			//ret[y][x] = uint8(math.Pow(float64(x), float64(y)))
		}
	}

	return ret

}

// Pour les maps
func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	decoupe := strings.Fields(s)

	for _, mot := range decoupe {
		if _, exist := ret[mot]; exist {
			ret[mot] += 1
		} else {
			ret[mot] = 1
		}
	}

	return ret
}

// Pour fonction closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Pour les methodes
func (v *Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pour les interfaces
type Abser interface {
	Abs() float64
}

type Person struct {
	Name string
	Age  int
}

// Interface stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Pour les errors.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string { // Fonction repondant à l'interface Error
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")

	val := 8
	passage(val)
	fmt.Println(val)

	ptr := &val
	passage_ptr(ptr)
	fmt.Println(val)

	/// Struct ///
	v1 := Vertex{}
	v2 := Vertex{1, 2}
	v3 := Vertex{X: 33}
	v4 := &Vertex{Y: 6}

	fmt.Println(v1, v2, v3, v4, *v4)

	/// ¤¤ Conteneurs ¤¤ ///
	// -- Array --
	// constante size array
	var tab [10]int
	fmt.Println(tab, len(tab))

	// -- Slice --
	// tableau a taille variable
	tab2 := []int{1, 2, 3, 45, 6, 789}
	fmt.Println(tab2)

	for i := 0; i < len(tab2); i++ {
		fmt.Print(tab2[i], " ")
	}
	fmt.Println()

	// Il est possible de slicer les tableaux et les slices
	// missing low index implies 0
	fmt.Println("tab2[:3] ==", tab2[:3])

	// missing high index implies len(tab2)
	fmt.Println("tab2[4:] ==", tab2[4:])

	fmt.Println(tab2[4:][0])

	a := make([]int, 5) // len(a)=5
	printSlice("a", a)
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)
	var z []int // Slice vide
	printSlice("z", z)

	// Append to slice
	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)

	b = append(b, 1)
	printSlice("b", b)
	z = append(z, 1)
	printSlice("z", z)

	// Parcours de slice avec for
	// Fonctionne avec les slices et maps
	for i, v := range tab2 {
		fmt.Printf("%d : %d\n", i, v)
	}
	// Pour n'avoir que l'index
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	// Pour ne pas s'occuper de l'index
	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	//pic.Show(Pic)

	// -- Map --
	// conteneur clef -> valeur
	var m map[string]Vertex     // A ce moment la map est 'nil'.
	m = make(map[string]Vertex) // Elle doit etre alloue avec make
	// ou directement m:= make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40, -74}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	// Peut etre instancie directement
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": Vertex{
			37, -122,
		},
	}
	fmt.Println(m2)

	// If the top-level type is just a type name, you can omit it from the
	// elements of the literal.
	var m3 = map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println(m3, len(m3))

	// Suppresion d'un element
	delete(m2, "Google")

	// Test si un element est present
	v, ok := m3["Google"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = m3["Chopek"]
	fmt.Println("The value:", v, "Present?", ok)

	// Attention si un element n'est pas dans une map il y aura quand meme
	// une variable de retour representant un objet nul.

	// Exmple avec fonction WordCount
	phrase := "I am learning Go!"
	fmt.Println(phrase)
	fmt.Println(WordCount(phrase))

	phrase = "I ate a donut. Then I ate another donut."
	fmt.Println(phrase)
	fmt.Println(WordCount(phrase))

	// ¤¤ Fonctions ¤¤
	// Les fonctione peuvent etre des valeurs
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	// Elles peuvent etre des closures
	// A closure is a function value that references variables from outside its
	// body. The function may access and assign to the referenced variables;
	// in this sense the function is "bound" to the variables.
	// cf adder, the adder function returns a closure.
	// Each closure is bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// ¤¤ Methodes ¤¤
	// Il est possible d'ajouter des méthodes aux classes
	// Go does not have classes. However, you can define methods on struct types.
	// The method receiver appears in its own argument list between
	// the func keyword and the method name.
	// cf fonction (method) Abs
	vMeth := &Vertex{3, 4}
	fmt.Println(vMeth.Abs())
	// You can declare a method on any type that is declared in your package,
	// not just struct types.
	// However, you cannot define a method on a type from another package
	// (including built in types).
	// cf autre fonction Abs
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Si la methode s'attache a un pointeur de strucure il est possible de
	// modifier ses parametres
	// cf func Scale
	fmt.Println(vMeth)
	vMeth.Scale(3)
	fmt.Println(vMeth)

	// ¤¤ Interface ¤¤
	// Permet une sorte de polymorphisme
	// An interface type is defined by a set of methods.
	// A value of interface type can hold any value that implements those methods.
	// cf Abser
	var abser Abser
	abser = f  // a MyFloat implements Abser
	abser = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//abser = v
	fmt.Println(abser)

	// Pour la description des elements il est possible d'utilsier l'interface
	// Stringer utilise par les print de fmt.
	// Pour cela il faut redefinir String pour un objet ex avec Personne
	p1 := Person{"Arthur Dent", 42}
	p2 := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(p1, p2)

	// ¤¤ Errors ¤¤
	// Les errors sont des interfaces :
	//type error interface {
	//Error() string
	//}

	err := run()
	if err != nil {
		fmt.Println(err)
	}

}
