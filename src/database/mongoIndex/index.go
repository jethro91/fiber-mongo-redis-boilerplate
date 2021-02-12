package mongoIndex

func CreateMongoIndexes() error {
	ch1 := createUserIndex()
	ch2 := createPasswordResetIndex()
	ch3 := createRoleIndex()

	var err error
	err = <-ch1
	if err != nil {
		return err
	}
	err = <-ch2
	if err != nil {
		return err
	}
	err = <-ch3
	if err != nil {
		return err
	}

	return err
}
