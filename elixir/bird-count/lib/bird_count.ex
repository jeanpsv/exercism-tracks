defmodule BirdCount do
  def today([]), do: nil
  def today([head | _tail]), do: head

  def increment_day_count([]), do: [1]
  def increment_day_count([head | tail]), do: [head + 1 | tail]

  def has_day_without_birds?([]), do: false
  def has_day_without_birds?([head | tail]) do
    not_visited?(head) || has_day_without_birds?(tail)
  end

  defp not_visited?(0), do: true
  defp not_visited?(_), do: false

  def total([]), do: 0
  def total([head | tail]) do
    head + total(tail)
  end

  defp busy?(head) when head >= 5, do: 1
  defp busy?(_), do: 0

  def busy_days([]), do: 0
  def busy_days([head | tail]), do: busy?(head) + busy_days(tail)
end
