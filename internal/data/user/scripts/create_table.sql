CREATE TABLE IF NOT EXISTS User (
  id                    INT AUTO_INCREMENT NOT NULL,
  discord_tag           VARCHAR(255) NOT NULL,
  userId                VARCHAR(255) NOT NULL,
  mainTeam              VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO User
  (discord_tag, userId)
VALUES
  ('AntonioZrd#7863', '573659533361020941', 'none'),
  ('fhatti', '253893620837384192', 'none');