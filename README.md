# BitArray

Provides a structure for efficiently representing arrays of k-bit elements.

```
a := bitarray.New(3, 6) // Makes an array of 3 items of 6 bits each
a.Set(0, 0x2b)
a.Set(1, 0x1c)
a.Set(2, 0x3a)

a.Get(1) // == 0x1c
```
