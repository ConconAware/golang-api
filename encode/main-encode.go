package encode

import (
	"fmt"
)

var secretMessage = "My name is JELE"

func MainEncode() {
	fmt.Println(" ------ encode -------")
	//Symmetric Encryption AES
	fmt.Println("Symmetric encode:")
	SymmetricMain()
	//Asymmetric RSA
	fmt.Println()
	fmt.Println("\n Asymmetric encode:")
	AsymmetricMain()

}
