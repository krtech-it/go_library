package book

import (
	"go_library/internal/infrastructure/db/models"
	"go_library/internal/infrastructure/repository/common"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks(page, pageSize int) ([]models.Book, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(book *models.Book) error
	CheckAuthorByID(id string) error
	CheckBookName(name string) error
	UpdateBook(id string, book *models.Book) error
	DeleteBook(id string) error
	GetCountBooks() (int, error)
	GetUser(userId string) (*models.User, error)
}

type bookRepository struct {
	br *common.BaseRepository
}

func (r *bookRepository) DeleteBook(id string) error {
	return r.br.Db.Where(map[string]interface{}{"id": id}).Delete(&models.Book{}).Error
}

func (r *bookRepository) GetAllBooks(page, pageSize int) ([]models.Book, error) {
	var books []models.Book
	err := r.br.Paginate(page, pageSize).Preload("Author").Find(&books).Error
	return books, err
}

func (r *bookRepository) GetBookByID(id string) (models.Book, error) {
	var book models.Book
	err := r.br.Db.Preload("Author").Preload("Genres").Where(map[string]interface{}{"id": id}).Find(&book).Error
	return book, err
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.br.Db.Create(book).Error
}

func (r *bookRepository) CheckAuthorByID(id string) error {
	var author models.Author
	return r.br.Db.Where(map[string]interface{}{"id": id}).First(&author).Error
}

func (r *bookRepository) CheckBookName(name string) error {
	var book models.Book
	return r.br.Db.Where(map[string]interface{}{"title": name}).First(&book).Error
}

func (r *bookRepository) UpdateBook(id string, book *models.Book) error {
	return r.br.Db.Model(models.Book{}).Where(map[string]interface{}{"id": id}).Updates(book).Error
}

func (r *bookRepository) GetCountBooks() (int, error) {
	var count int64
	err := r.br.Db.Model(&models.Book{}).Count(&count).Error
	return int(count), err
}

func (r *bookRepository) GetUser(userId string) (*models.User, error) {
	var user models.User
	err := r.br.Db.Preload("Author").Where(map[string]string{"id": userId}).First(&user).Error
	return &user, err
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{br: common.NewBaseRepository(db)}
}
