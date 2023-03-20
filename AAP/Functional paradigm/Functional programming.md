- Approach to software development that uses pure functions to create maintainable software
- Reasoning: limitations of Imperative languages (mutating values in concurrency can be very dangerous)
- Mutating the variable can cause hard to detect bugs (what happens if in the middle of an operation the variable changes)
- Functions are basic building blocks
- Values are immutable
- Main principle, immutability, functions and declarative code
- Idea stems from mathematics
###### Working with Immutable data
- In functional programming, all values are immutable, removing the need for locking mechanisms
- Operations on variable generate new values, allowing the compiler to run operations in parallel
###### Building programs with functions
- Functions are the primary building blocks of our program
- Short and expressive, they do operations on data
- Complexity of apps is reduced when
  1. Values are immutable
  2. The result is affected only by function arguments
  3. No side-effects, no effects are generated beyond the value it returns
- We call these functions pure function
`fn (n) -> n + 2 end` pure function that adds two
###### Using values explicitly 
- always pass values explicity between functions, making it clear what the input and outputs are
```elixir
defmodule MySet do
  defstruct items: []

  def push(set = %{items: items}, item) do
    if Enum.member?(items, item) do
      set
    else
      %{set | items: items ++ [item]}
    end
  end
end

defmodule Main do
  def main do
    set = %MySet{}
    set = MySet.push(set, "apple")
    new_set = %MySet{}
    new_set = MySet.push(new_set, "pie")
    IO.inspect(MySet.push(set, "apple"))
    IO.inspect(MySet.push(new_set, "apple"))
  end
end

Main.main()
```
- Operations and data are detached
###### Using functions as arguments
- Functions can and should be used as arguments
`IO.inspect(Enum.map(["apple", "dog"], &String.upcase/1))`
###### Declaring code
- Functional programming is declarative
- We focus on what is necessary to solve a problem, describing a dataflow
```js
let list = ["dogs", "hotdogs"]

function upcase(list) {
  let newList = []
  list.forEach(item => {
    newList.push(item.toUpperCase())
  })
  return newList
}

console.log(upcase(list))
```
- Using the imperative mindset, a control flow structure like *forEach* is needed to navigate through the list. In the callback function provided to *forEach* we need to push the list item to a new list
- Here is the same example using declarative mindset
```elixir
defmodule StringList do
  def upcase([]), do: []
  def upcase([first | rest]), do: [String.upcase(first) | upcase(rest)]
end
```
- Even more simplified way of transforming the example is using higher-ordering functions
```elixir
list = ["dogs", "cats"]
Enum.map(list, &String.upcase/1)
```
- 