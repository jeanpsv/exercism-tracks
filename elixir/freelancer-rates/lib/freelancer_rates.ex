defmodule FreelancerRates do
  @work_hours 8.0
  @work_days_in_month 22

  def daily_rate(hourly_rate) do
    hourly_rate * @work_hours
  end

  def apply_discount(before_discount, discount) do
    1
    |> Kernel.-(discount / 100)
    |> Kernel.*(before_discount)
  end

  def monthly_rate(hourly_rate, discount) do
    hourly_rate
    |> daily_rate()
    |> Kernel.*(@work_days_in_month)
    |> apply_discount(discount)
    |> ceil()
    |> trunc()
  end

  def days_in_budget(budget, hourly_rate, discount) do
    daily_cost =
      hourly_rate
      |> daily_rate()
      |> apply_discount(discount)

    budget
    |> Kernel./(daily_cost)
    |> Float.floor(1)
    |> Float.ceil(1)
  end
end
