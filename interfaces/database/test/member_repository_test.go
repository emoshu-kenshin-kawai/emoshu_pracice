package database

import (
	"emoshu_practice/domain"
	"emoshu_practice/infrastructure"
	"emoshu_practice/interfaces/database"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDummyHandler(db *gorm.DB) database.DBHandler {
	dbHandler := new(infrastructure.DBHandler)
	dbHandler.DB = db
	return dbHandler
}

func newDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}
	mockDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName:                "mysql",
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		}},
		&gorm.Config{})
	return mockDB, mock, err
}

func TestNew(t *testing.T) {
	mockDB, mock, err := newDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
	}
	member := &domain.Member{
		ID:                 1,
		No:                 "1",
		ProfileImg:         "http://hoge.png",
		FullName:           "emoshu company",
		KanaName:           "emoshu company",
		Motto:              "頑張ります",
		Biography:          "",
		StartDate:          time.Now(),
		EndDate:            nil,
		EmploymentStatusID: 1,
		StatusID:           1,
	}
	rows := sqlmock.NewRows([]string{"id", "no", "profileImg", "fullName", "kanaName", "motto", "biography", "startDate", "endDate"}).AddRow(member.ID, member.No, member.ProfileImg, member.FullName, member.KanaName, member.Motto, member.Biography, member.StartDate, member.EndDate)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "members" ("id", "no", "profileImg", "fullName", "kanaName", "motto", "biography", "startDate", "endDate") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`)).WillReturnRows(rows)
	mock.ExpectCommit()

	repo := &database.MemberRepository{DBHandler: newDummyHandler(mockDB)}
	m, _ := repo.New(*member)
	assert.Equal(t, m.ID, member.ID)
	assert.Equal(t, m.No, member.No)
	assert.Equal(t, m.ProfileImg, member.ProfileImg)
	assert.Equal(t, m.FullName, member.FullName)
	assert.Equal(t, m.KanaName, member.KanaName)
	assert.Equal(t, m.Motto, member.Motto)
	assert.Equal(t, m.Biography, member.Biography)
	assert.Equal(t, m.StartDate, member.StartDate)
	assert.Equal(t, m.EndDate, member.EndDate)
	assert.Equal(t, m.EmploymentStatusID, member.EmploymentStatusID)
	assert.Equal(t, m.StatusID, member.StatusID)
}

func TestFindById(t *testing.T) {

	mockDB, mock, err := newDbMock()
	if err != nil {
		t.Errorf("failed to initialize mock DB: %v", err)
	}
	member := &domain.Member{
		ID:                 1,
		No:                 "1",
		ProfileImg:         "http://hoge.png",
		FullName:           "emoshu company",
		KanaName:           "emoshu company",
		Motto:              "頑張ります",
		Biography:          "",
		StartDate:          time.Now(),
		EndDate:            nil,
		EmploymentStatusID: 1,
		EmploymentStatus: domain.EmploymentStatus{
			ID:               0,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			EmploymentStatus: "",
		},
		StatusID: 1,
		Status: domain.Status{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Status:    "",
		},
	}
	rows := sqlmock.
		NewRows([]string{"id", "no", "profileImg", "fullName", "kanaName", "motto", "biography", "startDate", "endDate", "employmentStatusID", "employmentStatus", "statusID", "status"}).
		AddRow(member.ID, member.No, member.ProfileImg, member.FullName, member.KanaName, member.Motto, member.Biography, member.StartDate, member.EndDate, member.EmploymentStatusID, member.EmploymentStatus, member.StatusID, member.Status)
	// mock.ExpectBegin()
	// mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "members" ("id", "no", "profileImg", "fullName", "kanaName", "motto", "biography", "startDate", "endDate") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`)).WillReturnRows(rows)
	// mock.ExpectCommit()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "members" WHERE id = $1 ORDER BY "id" LIMIT 1`)).WithArgs(member.ID).WillReturnRows(rows)

	repo := &database.MemberRepository{DBHandler: newDummyHandler(mockDB)}
	m, err := repo.FindById("1")
	fmt.Println(m, "membr")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// assert.NotEqual(t, len(res), nil)

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("test find member: %v", err)
	// }

}

// ID                 uint             `gorm:"primarykey" json:"id"`
// CreatedAt          time.Time        `gorm:"autoCreateTime:milli" json:"created_at"`
// UpdatedAt          time.Time        `gorm:"autoUpdateTime:milli" json:"updated_at"`
// No                 string           `json:"no"`
// ProfileImg         string           `gorm:"not null" json:"profile_img"`
// FullName           string           `gorm:"not null" json:"full_name"`
// KanaName           string           `json:"kana_name"`
// Motto              string           `json:"motto"`
// Biography          string           `json:"biography"`
// StartDate          time.Time        `json:"start_date"`
// EndDate            *time.Time       `json:"end_date"`
// EmploymentStatusID uint             `json:"employment_status_id"`
// EmploymentStatus   EmploymentStatus `json:"employment_status"`
// StatusID           uint             `json:"status_id"`
// Status             Status           `json:"status"`
