# go-leveldb

## key-value database modelling

How to model a key-value database correctly?

## Primary key

Let's try to imagine we have a simple json object that we would like to store in a key-value store:

```json
{
  "id": 1,
  "name": "john"
}
```

We can model it this way:

```
key:value
1:john
2:beta
3:casey
4:john
```

Each row will be an id, and it will have the name as the value. 

```
# Put operation
db.put(1, "john")
db.put(2, "beta")

# Get operation
db.get(1) // Returns "john"
```

As the name dictates, we should only store key-value pairs. It's probably not a good idea to store everything as the key, even though most people do so for quick key lookup:

```
# Sample row, the value is empty and we only have a key
1/john:1

# Put operation
db.put("1/john", 1) // 0 = false, 1 = true

# Get operation (acting as "Has")
db.get("1/john") // 1 = true
```

## Growing records

That's fine if we only have one key and one value. What if we have several?

```
{
  "id": 1,
  "name": "john",
  "age": 20,
  "color": "RED", // Enum RED, GREEN or BLUE
}
```

First idea that crossed our mind could be:

```
1:'{"name": "john", "age": 20, "color": "RED"}'
```

Operations
```
# Put (the value is normally stored as byte)
db.put(1, {"name": "john", "age": 20, "color": "RED"})

# Get
db.get(1)
```

However, this poses some problem:
- we can't perform complex query
- we can't search for the field without querying each row
- we can't perform filter (e.g. by age > 10, color = "RED")

Remember that we are using a key-value store. The keys are normally sorted lexigraphically, and therefore the prefix matters. Instead of using just using one key, we can create composite keys:

```
# With one key
primary-key:primary-value/secondary-key:secondary-value
id/name/age
1/john/20
```

The order of the keys matter. If we use age as the primary key first, then the user will be sorted by age first.

With this, we can perform filters too:

```
# Bad
db.put("1/john/20", "RED")
// Note that in order to perform filtering, putting the unique index in front is meaningless
db.iter(prefix{"1/john"})

# Better
db.put("john/20/1", "RED")
db.iter(prefix{"john"}) // Filter user with name john
```
