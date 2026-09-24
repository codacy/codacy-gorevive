## inefficient-map-lookup

_Go version_: 1.0.

**_Typed_**

_Description_: This rule identifies code that iteratively searches for a key in a map.

This inefficiency is usually introduced when refactoring code from using a slice to a map.
For example if during refactoring the `elements` slice is transformed into a map.

```diff
-       elements             []string
+       elements             map[string]float64
```

and then a loop over `elements` is changed in an obvious but inefficient way:

```diff
-       for _, e := range elements {
+       for e := range elements {
                if e == someStaticValue {
                        // do something
                }
        }
```

Configuration example:

```go
aMap := map[string]bool{}{}
aValue := false

// Inefficient map lookup
for k := range aMap {
	if k == aValue {
		// do something
	}
}

// Simpler and more efficient version
if _, ok := aMap[aValue]; ok {
	// do something
}
```

_Configuration_: N/A

