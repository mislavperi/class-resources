# Assign them to variables
double = fn x -> x * 2 end
IO.inspect(double.(2))
# Pass them as arguments
IO.inspect(Enum.map(1..4, fn x -> x * x end))
