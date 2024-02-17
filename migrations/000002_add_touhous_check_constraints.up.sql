ALTER TABLE
  touhous
ADD
  CONSTRAINT abilities_length_check CHECK (
    array_length(abilities, 1) BETWEEN 1
    and 10
  )