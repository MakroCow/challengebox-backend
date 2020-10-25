package databaseService

// Runs all functions to initialize Database Tables and introduces Fake Data
func InitializeDatabaseTables() error {
	err := CreateChallengeTable()
	if err != nil {
		return err
	}
	err = CreateTagTable()
	if err != nil {
		return err
	}
	err = CreateTagChallengeTable()
	if err != nil {
		return err
	}
	err = InitializeTestData()
	if err != nil {
		return err
	}
	return nil
}

func InitializeTestData() error {
	err := createFakeChallenges()
	if err != nil {
		return err
	}
	err = createFakeTags()
	if err != nil {
		return err
	}
	err = crateFakeTagChallengeConnection()
	if err != nil {
		return err
	}
	return nil
}

func crateFakeTagChallengeConnection() error {
	return nil
}

func createFakeTags() error {
	return nil
}

//Creates the Tag Table in Database
func CreateTagTable() error {
	sqlStatement := `create table tag(
    					id   serial  not null
        					constraint tag_pk
            				primary key,
    					name varchar not null)`
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `alter table tag owner to postgres`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `create unique index tag_id_uindex on tag (id)`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}
	return nil
}

// Creates the Tag-Challenge Table in Database
func CreateTagChallengeTable() error {
	sqlStatement := `create table tag_challenge(
    id           serial  not null
        		 constraint tag_challenge_pk
           		 primary key,
    challenge_id integer not null,
    tag_id       integer not null)`
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `alter table tag_challenge owner to postgres`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `create unique index tag_challenge_id_uindex on tag_challenge (id)`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}
	return nil
}

// Creates the Challenge Table in Database
func CreateChallengeTable() error {
	sqlStatement := `create table challenge 
			(id serial not null constraint challenge_pk primary key, 
			 title varchar not null, 
			 description varchar not null, 
			 sport_value integer not null, 
			 intelligence_value integer not null, 
			 culinary_value integer, 
			 social_value integer, 
			 selfcare_value  integer)`
	_, err := DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `alter table challenge owner to postgres`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}

	sqlStatement = `create unique index challenge_id_uindex on challenge (id);`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		return error(err)
	}
	return nil
}

// Fills Challenge Table with fake Challenges
func createFakeChallenges() error {
	sqlStatement := `INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (9, 'Bath', 'Take a bath', 2, 2, 0, 0, 20);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (1, 'Kuchen backen', 'Schoko doer Vanille', 3, 12, 1, 1, 1);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (10, 'Read in the Bath', 'Take a book in the bath.', 2, 10, 0, 3, 20);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (8, 'Run', 'Go for a run for at least 20 minutes or 3km. ', 20, 0, 0, 0, 10);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (2, 'Tarzan Call', 'Go to an elevated area to call out a tarzan call. ', 2, 5, 0, 10, 3);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (7, 'Walk', 'Go for a walk for at last 30 minutes.', 10, 2, 0, 0, 10);
					 INSERT INTO public.challenge (id, title, description, sport_value, intelligence_value, culinary_value, social_value, selfcare_value) VALUES (6, 'Walk', 'Go for a walk for at last 15 minutes.', 5, 1, 0, 0, 5);`

	_, err := DB.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}
