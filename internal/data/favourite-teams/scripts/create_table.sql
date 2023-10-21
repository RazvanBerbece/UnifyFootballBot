CREATE TABLE IF NOT EXISTS FavouriteTeams (
  id                  INT AUTO_INCREMENT NOT NULL,
  userId              VARCHAR(255) NOT NULL,
  favouriteTeam       VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);