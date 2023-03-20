defmodule Hof do
  def tripler(value, function) do
    3 * function.(value)
  end
end

IO.inspect(Hof.tripler(2, fn x -> x * x end))
