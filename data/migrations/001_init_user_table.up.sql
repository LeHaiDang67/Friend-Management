CREATE TABLE IF NOT EXISTS users
(
  email text NOT NULL,
  friends text[],
  subscription text[],
  blocked text[],
  CONSTRAINT users_pkey PRIMARY KEY (email)
);
