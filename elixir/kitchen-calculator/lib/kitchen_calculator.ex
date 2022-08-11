defmodule KitchenCalculator do
  def get_volume({_, volume}) do
    volume
  end

  def to_milliliter({:cup, volume}), do: {:milliliter, volume * 240}
  def to_milliliter({:fluid_ounce, volume}), do: {:milliliter, volume * 30}
  def to_milliliter({:teaspoon, volume}), do: {:milliliter, volume * 5}
  def to_milliliter({:tablespoon, volume}), do: {:milliliter, volume * 15}
  def to_milliliter({:milliliter, volume}), do: {:milliliter, volume}

  def from_milliliter(volume_pair, unit) do
    {_, volume} = to_milliliter(volume_pair)
    {_, size} = to_milliliter({unit, 1})
    {unit, volume / size}
  end

  def convert(volume_pair, unit) do
    {_, volume} = to_milliliter(volume_pair)
    from_milliliter({:milliliter, volume}, unit)
  end
end
