defmodule TeatherFacade do
  alias Amplifier, as: Amp
  alias DVDPlayer, as: DVD
  alias PopcornPopper, as: PMaker

  def watchMovie(movie) do
    Amp.turnOnAmp() |> PMaker.turnOnPopcornMaker() |> PMaker.makePopcorn() |> DVD.turnOnMovie(movie)
  end
end
