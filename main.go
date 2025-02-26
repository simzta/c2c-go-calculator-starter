//Defining main package
package main

//import fmt
import "fmt"

//variables
var schoolYear = 3   //int

//structure
type Student struct{
  name string;
  major string;
  GPA float32;
  GradStatus bool;
  numClasses int
  class string
  }

//function
func classification(schoolYear int) string{
  switch schoolYear{   //switch
    case 1:
      return "Freshman"
    case 2:
      return "Sophomore"
    case 3:
      return "Junior"
    case 4:
      return "Senior"
    default:
      return "Not an valid year."
  }
}

//main function
func main(){

  var  student1 Student
  student1.name ="Felicia Forester"
  student1.major = "CS Engineering"
  student1.numClasses = 5
  student1.GPA = 4.00
  student1.GradStatus = false
  student1.class = classification(4)

  if student1.class == "Senior"{
    fmt.Println(student1)
    fmt.Println("Congrats!")

  } else {
    fmt.Println(student1)

  }

}
