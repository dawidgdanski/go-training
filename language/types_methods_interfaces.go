package language

import (
	"fmt"
	"net/http"
)

/*
*
https://news.ycombinator.com/item?id=35784598
*/
func init() {
	// comparing interface implementations
	comparingInterfaceExample()
	// Runtime error - slices are not comparable
	// DoublerCompare(dis, dis2)

	// empty interface
	emptyInterfaceExample()

	typeAssertionExample1()

	typeAssertionExample2()

	typeAssertionExample3(1)

}

func typeAssertionExample1() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
}

func typeAssertionExample2() {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(MyInt)
	if ok {
		fmt.Println("Success:", i2)
	} else {
		fmt.Println(fmt.Errorf("failure casting %d to %v", i, i))
	}
}

func typeAssertionExample3(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt")
	case bool:
		fmt.Println("bool")
	default:
		panic(fmt.Sprintf("Unknow type: %s", j))
	}
}

func comparingInterfaceExample() {
	var di DoubleInt = 10
	di.Double()
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	// var dis2 = DoubleIntSlice{1, 2, 3}

	DoublerCompare(&di, &di2)
	DoublerCompare(&di, dis)
}

func emptyInterfaceExample() {
	var i interface{} // or var i any
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Alice", "Bot"}
	fmt.Println(i)
}

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

type MyInt int

//Function meeting interface

type DataStore interface {
	UserNameForID(userId string) (string, bool)
	SayHello(id string) (string, error)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (log LoggerAdapter) Log(message string) {
	log(message)
}

type Controller struct {
	logger    Logger
	dataStore DataStore
}

func (controller Controller) SayHello(writer http.ResponseWriter, request *http.Request) {
	controller.logger.Log("In SayHello")
	userID := request.URL.Query().Get("user_id")
	message, err := controller.dataStore.SayHello(userID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	} else {
		_, err := writer.Write([]byte(message))
		if err != nil {
			return
		}
	}
}
