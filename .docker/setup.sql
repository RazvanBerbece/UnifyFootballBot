USE unifyfootballdb;

CREATE TABLE IF NOT EXISTS User (
  id                    INT AUTO_INCREMENT NOT NULL,
  discord_tag           VARCHAR(255) NOT NULL,
  userId                VARCHAR(255) NOT NULL,
  mainTeam              VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
INSERT INTO User
  (discord_tag, userId, mainTeam)
VALUES
  ('AntonioZrd#7863', '573659533361020941', 'none'),
  ('fhatti', '253893620837384192', 'none');

CREATE TABLE IF NOT EXISTS FavouriteTeams (
  id                  INT AUTO_INCREMENT NOT NULL,
  userId              VARCHAR(255) NOT NULL,
  favouriteTeam       VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);