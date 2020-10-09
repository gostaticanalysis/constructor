package a

/* Struct literal */
func Create() (v struct{}) { return } // OK it is not a constructor like function

func CreatePtr() (v *struct{}) { return } // OK it is not a constructor like function

/* Named struct */
type T1 struct{}

func NewT1() (t *T1) { return } // OK

func CreateT1() (t *T1) { return } // want `name of a constructor like function must begin "New"`

func CreateT_1() (t *T1) { return } // want `name of a constructor like function must begin "New"`

func NoPtrT_1() (t T1) { return } // want `name of a constructor like function must begin "New"`

func ErrT_1() (t *T1, e error) { return } // want `name of a constructor like function must begin "New"`

func NoPtrErrT_1() (t T1, e error) { return } // want `name of a constructor like function must begin "New"`

func ErrMultiRetT_1() (t *T1, n int, e error) { return } // OK it is not a constructor like function

func Err() (e error) { return } // OK it is not a constructor like function

func unexportT_1() (t *T1) { return } // OK it is not a constructor like function

func (t *T1) Clone() (_t *T1) { return } // OK it is a method

func (t T1) CloneNoPtr() (_t T1) { return } // OK it is a method

/* Non struct */
type T2 int

func CreateT2() (t *T2) { return } // OK T2 is not struct

func CreateT_2() (t *T2) { return } // OK T2 is not struct

func NoPtrT_2() (t T2) { return } // OK T2 is not struct

func ErrT_2() (t *T2, e error) { return } // OK T2 is not struct

func NoPtrErrT_2() (t T2, e error) { return } // OK T2 is not struct
