# Chapter 14. Here There Be Dragons: Reflect, Unsafe, and Cgo

## Reflection Lets Us Work with Types at Runtime

-> Curious how many of the reflection uses will disappear with generics
-> It will probably remain in the program boundaries

### Types, Kinds, and Values

Type

```go
var x int
xt := reflect.TypeOf(x)
fmt.Println(xt.Name())     // returns int

f := Foo{}
ft := reflect.TypeOf(f)
fmt.Println(ft.Name())     // returns Foo

xpt := reflect.TypeOf(&x)
fmt.Println(xpt.Name())    // returns an empty string
```

slice or pointer types don't have names.

Kind:
> if you define a struct named `Foo`, the kind is `reflect.Struct` and the type is `“Foo”`.

```go
type Foo struct {
    A int    `myTag:"value"`
    B string `myTag:"value2"`
}

var f Foo
ft := reflect.TypeOf(f)
for i := 0; i < ft.NumField(); i++ {
    curField := ft.Field(i)
    fmt.Println(curField.Name, curField.Type.Name(),
        curField.Tag.Get("myTag"))
}
```

#### Values

`vValue := reflect.ValuesOf(v)`

### Making New Values

-> Not convinced I will ever need it.


### Use Reflection to Check If an Interface's Value is `nil`

```go
func hasNoValue(i interface{}) bool {
    iv := reflect .ValueOf(i)
    if !iv.IsValid() {
        return true
    }

    switch iv.Kind() {
    case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
        return iv.IsNil()
    default:
        return falset
    }
}
```

### Use Reflection to Write a Data Marshaller

CSV 

### Build Functions with Reflection to Automate Repetitive Tasks

```go
func MakeTimedFunction(f interface{}) interface{} {
    ft := reflect.TypeOf(f)
    fv := reflect.ValueOf(f)
    wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
        start := time.Now()
        out := ft.Call(in)
        end := time.Now()
        fmt.Println(end.Sub(start))
        return out
    })
    return wrapperF.Interface()
}
```

### You Can Build Structs with Reflection, but Don't

### Reflection Can't Make Methods

### Only Use Reflection If It's Worthwhile

#### Use unsafe to Convert External Binary Data



## unsafe is Unsafe

## Cgo is for Integration, Not Performance


-> This chapter didn't seem practical for me. I'm not going to remember those until I need to use them. 
-> And I don't think I will need to use them in the first place...
-> But, if I need then I can go back to this chapter or to docs. 

