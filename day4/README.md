# Day4

## Array , Slice, and Map

cara membuat array

```go

// array yang sudah di tentukan jumlah dan isi nya

var colors = [3]string{"red", "yellow","brown"}
```

menambah koleksi array

```go
var colors = [3]string{"red", "yellow","brown"}
colors = append("blue")
```

cara membuat array multi dimensi

```go
multi := [][]int{} // ini slice
multi = append(multi, []int{1, 2})
multi = append(multi, []int{3, 4})

// [[1 2] [3 4]]
```
