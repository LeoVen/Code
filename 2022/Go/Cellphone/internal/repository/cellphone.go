package repository

import (
	"cellphone/internal/entity"
	"database/sql"

	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	getByIdCellphone           = "SELECT * FROM CELLPHONE WHERE ID = ?;"
	insertSingleCellphone      = "INSERT INTO CELLPHONE (PROVIDER_ID, NUMBER) VALUES (?, ?);"
	serveSingleCellphone       = "SELECT * FROM CELLPHONE WHERE PROVIDER_ID = ? LIMIT 1"
	serveSingleDeleteCellphone = "DELETE FROM CELLPHONE WHERE ID = ?;"
)

type CellphoneRepositoryMock struct {
	db *sql.DB
}

type CellphoneRepositorySql struct {
	db *sql.DB
}

type CellphoneRepositoryGorm struct {
	db *gorm.DB
}

func NewCellphoneRepository(repoType int, conn *sql.DB) (CellphoneRepository, error) {
	switch repoType {
	case REPO_MOCK:
		return newCellphoneRepositoryMock(conn)
	case REPO_SQL:
		return newCellphoneRepositorySql(conn)
	case REPO_GORM:
		return newCellphoneRepositoryGorm(conn)
	}
	return newCellphoneRepositorySql(conn)
}

func newCellphoneRepositorySql(conn *sql.DB) (CellphoneRepository, error) {
	return &CellphoneRepositorySql{conn}, nil
}

func newCellphoneRepositoryMock(conn *sql.DB) (CellphoneRepository, error) {
	return &CellphoneRepositoryMock{conn}, nil
}

func newCellphoneRepositoryGorm(conn *sql.DB) (CellphoneRepository, error) {
	gormDB, err := gorm.Open(mysqlGorm.New(mysqlGorm.Config{
		Conn: conn,
	}))

	if err != nil {
		return nil, err
	}

	return &CellphoneRepositoryGorm{gormDB}, nil
}

func (self *CellphoneRepositorySql) GetById(id int) (interface{}, error) {
	row := self.db.QueryRow(getByIdCellphone, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var phone entity.Cellphone
	if err := row.Scan(&phone.Id, &phone.Number, &phone.ProviderId); err != nil {
		return nil, err
	}

	return phone, nil
}

// Retrieves a semi-random phone number and deletes it from the database
func (self *CellphoneRepositorySql) ServeSingleFromProvider(providerId int) (*entity.Cellphone, error) {
	row := self.db.QueryRow(serveSingleCellphone, providerId)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity entity.Cellphone

	err := row.Scan(&entity.Id, &entity.ProviderId, &entity.Number)

	if err != nil {
		return nil, err
	}

	tx, err := self.db.Begin()

	_, err = tx.Exec(serveSingleDeleteCellphone, entity.Id)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (self *CellphoneRepositorySql) InsertSingle(cellphone *entity.Cellphone) error {
	tx, err := self.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(insertSingleCellphone, cellphone.ProviderId, cellphone.Number)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (self *CellphoneRepositoryMock) GetById(id int) (interface{}, error) {
	row := self.db.QueryRow(getByIdCellphone, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var phone entity.Cellphone
	if err := row.Scan(&phone.Id, &phone.Number, &phone.ProviderId); err != nil {
		return nil, err
	}

	return phone, nil
}

func (self *CellphoneRepositoryMock) ServeSingleFromProvider(providerId int) (*entity.Cellphone, error) {
	row := self.db.QueryRow(serveSingleCellphone, providerId)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var entity entity.Cellphone

	err := row.Scan(&entity.Id, &entity.ProviderId, &entity.Number)

	if err != nil {
		return nil, err
	}

	tx, err := self.db.Begin()

	_, err = tx.Exec(serveSingleDeleteCellphone, entity.Id)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (self *CellphoneRepositoryMock) InsertSingle(cellphone *entity.Cellphone) error {
	tx, err := self.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(insertSingleCellphone, cellphone.ProviderId, cellphone.Number)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (self *CellphoneRepositoryGorm) GetById(id int) (interface{}, error) {
	var result entity.Cellphone
	tx := self.db.Model(&entity.Cellphone{}).First(&result, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (self *CellphoneRepositoryGorm) ServeSingleFromProvider(providerId int) (*entity.Cellphone, error) {
	var result entity.Cellphone

	self.db.Model(&entity.Cellphone{}).Where(&entity.Cellphone{
		ProviderId: providerId,
	}).First(&result)

	if self.db.Error != nil {
		return nil, self.db.Error
	}

	tx := self.db.Begin()
	defer func() {
		if tx.Error == nil {
			tx.Commit()
		}
	}()

	if tx.Error != nil {
		return nil, tx.Error
	}

	self.db.Model(&entity.Cellphone{}).Delete(&entity.Cellphone{}, 0)

	if self.db.Error != nil {
		return nil, self.db.Error
	}

	return &result, nil
}

func (self *CellphoneRepositoryGorm) InsertSingle(cellphone *entity.Cellphone) error {
	tx := self.db.Model(&entity.Cellphone{}).Begin()
	defer func() {
		if tx.Error == nil {
			tx.Commit()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	tx.Create(cellphone)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
