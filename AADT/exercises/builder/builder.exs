defmodule MealBuilder do
  def init_burger,
    do: %Burger{}

  def init_soup,
    do: %Soup{}

  def init_past,
    do: %Pasta{}

  def init_meal,
    do: %Meal{}

  def add_burger(meal) do
    burger = init_burger() |> add_bun |> add_meat |> add_toppings("mayo") |> add_meal(meal)
  end

  def add_pasta(meal) do
    pasta = init_past() |> add_pasta_sauce |> add_pasta_cheese |> add_meal(meal)
  end

  def add_soup(meal) do
    soup = init_soup() |> add_soup_type("tomato") |> add_meal(meal)
  end

  def add_soup_type(soup, type),
    do: %{soup | type: type}

  def add_pasta_sauce(pasta),
    do: %{pasta | sauce: true}

  def add_pasta_cheese(pasta),
    do: %{pasta | cheese: true}

  def add_meal(newMeal, meal),
    do: %{meal | meals: [newMeal | meal.meals]}

  def add_bun(burger),
    do: %{burger | bun: true}

  def add_meat(burger),
    do: %{burger | meat: true}

  def add_toppings(burger, newTopping),
    do: %{burger | toppings: [newTopping | burger.toppings]}
end
