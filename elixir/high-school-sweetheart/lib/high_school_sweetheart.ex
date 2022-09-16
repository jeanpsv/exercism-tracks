defmodule HighSchoolSweetheart do
  def first_letter(name) do
    name
    |> String.trim()
    |> String.first()
  end

  def initial(name) do
    name
    |> first_letter()
    |> String.upcase()
    |> Kernel.<>(".")
  end

  def initials(full_name) do
    [first, last] =
      full_name
      |> String.split()
      |> Enum.map(&initial/1)
    "#{first} #{last}"
  end

  def pair(full_name1, full_name2) do
    f1 =
      full_name1
      |> initials()

    f2 =
      full_name2
      |> initials()

    """
         ******       ******
       **      **   **      **
     **         ** **         **
    **            *            **
    **                         **
    **     #{f1}  +  #{f2}     **
     **                       **
       **                   **
         **               **
           **           **
             **       **
               **   **
                 ***
                  *
    """
  end
end
