Use test;

CREATE TABLE `user_movies`(
   user_name VARCHAR(40),
   movie VARCHAR(255),
   rating INT,
   PRIMARY KEY(user_name,movie)
);

INSERT INTO user_movies value ('Pompi','Gladiator',4);
INSERT INTO user_movies value ('Pompi','Dil Se',4);