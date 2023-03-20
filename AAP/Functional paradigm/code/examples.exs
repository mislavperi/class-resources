def Immutable do
  list = [1, 2, 3, 4]
  List.delete_at(list, -1)
  list ++ [1]
  IO.inspect(list)
end

defmodule Example do
  pure_f = fn n -> n + 2 end
  IO.inspect("I am a side effect")
end

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

defmodule StringList do
  def upcase([]), do: []
  def upcase([first | rest]), do: [String.upcase(first) | upcase(rest)]
end

defmodule Main do
  def main do
    # set = %MySet{}
    # set = MySet.push(set, "apple")
    # new_set = %MySet{}
    # new_set = MySet.push(new_set, "pie")
    # IO.inspect(MySet.push(set, "apple"))
    # IO.inspect(MySet.push(new_set, "apple"))
    # IO.inspect(Enum.map(["apple", "dog"], &String.upcase/1))
    IO.inspect(StringList.upcase(["dogs", "hot dogs", "bananas"]))
  end
end

Main.main()
