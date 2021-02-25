package i18n

import (
	"errors"
	"fmt"
	"strings"
)

const (
	LanguageEnglish = "EN"
	LanguageRussian = "RU"

	// page titles
	keyTitleLogin          Key = "title.login"
	keyTitleRegister       Key = "title.register"
	keyTitleDashboard      Key = "title.dashboard"
	keyTitleCustomers      Key = "title.customers"
	keyTitleCreateCustomer Key = "title.create-customer"

	// buttons
	keyButtonLogin          Key = "button.login"
	keyButtonRegister       Key = "button.register"
	keyButtonLogout         Key = "button.logout"
	keyButtonAddCustomer    Key = "button.add-customer"
	keyButtonCreateCustomer Key = "button.create-customer"
	keyButtonEditCustomer   Key = "button.edit-customer"
	keyButtonSaveCustomer   Key = "button.save-customer"
	keyButtonCancel         Key = "button.cancel"

	// form fields
	keyFieldEmail        Key = "field.email"
	keyFieldPassword     Key = "field.password"
	keyFieldSearch       Key = "field.search"
	keyFieldFirstName    Key = "field.first-name"
	keyFieldLastName     Key = "field.last-name"
	keyFieldBirthDate    Key = "field.birth-date"
	keyFieldGender       Key = "field.gender"
	keyFieldGenderMale   Key = "field.gender-male"
	keyFieldGenderFemale Key = "field.gender-female"
	keyFieldAddress      Key = "field.address"

	// links
	keyLinkLogin    Key = "link.login"
	keyLinkRegister Key = "link.register"

	// table columns
	keyColumnName      Key = "column.name"
	keyColumnEmail     Key = "column.email"
	keyColumnBirthDate Key = "column.birth-date"
	keyColumnGender    Key = "column.gender"
	keyColumnAddress   Key = "column.address"
	keyColumnJanuary   Key = "column.january"
	keyColumnFebruary  Key = "column.february"
	keyColumnMarch     Key = "column.march"
	keyColumnApril     Key = "column.april"
	keyColumnMay       Key = "column.may"
	keyColumnJune      Key = "column.june"
	keyColumnJuly      Key = "column.july"
	keyColumnAugust    Key = "column.august"
	keyColumnSeptember Key = "column.september"
	keyColumnOctober   Key = "column.october"
	keyColumnNovember  Key = "column.november"
	keyColumnDecember  Key = "column.december"

	// error message
	KeyErrorEmailTaken         Key = "error.email-taken"
	KeyErrorCredentialsInvalid Key = "error.credentials-invalid"
	KeyErrorDataCollision      Key = "error.data-collision"

	// other
	keyNoSearchResults Key = "message.no-search-results"
)

type Key string

var (
	translations = map[Key]map[string]string{
		// page titles
		keyTitleLogin: {
			LanguageEnglish: "Sign in",
			LanguageRussian: "Вход",
		},
		keyTitleRegister: {
			LanguageEnglish: "Create account",
			LanguageRussian: "Создание аккаунта",
		},
		keyTitleDashboard: {
			LanguageEnglish: "Dashboard",
			LanguageRussian: "Главная",
		},
		keyTitleCustomers: {
			LanguageEnglish: "Customers",
			LanguageRussian: "Клиенты",
		},
		keyTitleCreateCustomer: {
			LanguageEnglish: "New customer",
			LanguageRussian: "Новый клиент",
		},

		// buttons
		keyButtonLogin: {
			LanguageEnglish: "Sign in",
			LanguageRussian: "Войти",
		},
		keyButtonRegister: {
			LanguageEnglish: "Create account",
			LanguageRussian: "Создать аккаунт",
		},
		keyButtonLogout: {
			LanguageEnglish: "Sign out",
			LanguageRussian: "Выйти",
		},
		keyButtonAddCustomer: {
			LanguageEnglish: "Add",
			LanguageRussian: "Создать",
		},
		keyButtonCreateCustomer: {
			LanguageEnglish: "Create",
			LanguageRussian: "Создать",
		},
		keyButtonEditCustomer: {
			LanguageEnglish: "Edit",
			LanguageRussian: "Редактировать",
		},
		keyButtonSaveCustomer: {
			LanguageEnglish: "Save",
			LanguageRussian: "Сохранить",
		},
		keyButtonCancel: {
			LanguageEnglish: "Cancel",
			LanguageRussian: "Отмена",
		},

		// form fields
		keyFieldEmail: {
			LanguageEnglish: "Email",
			LanguageRussian: "Имейл",
		},
		keyFieldPassword: {
			LanguageEnglish: "Password",
			LanguageRussian: "Пароль",
		},
		keyFieldSearch: {
			LanguageEnglish: "Search",
			LanguageRussian: "Поиск",
		},
		keyFieldFirstName: {
			LanguageEnglish: "First name",
			LanguageRussian: "Имя",
		},
		keyFieldLastName: {
			LanguageEnglish: "Last name",
			LanguageRussian: "Фамилия",
		},
		keyFieldBirthDate: {
			LanguageEnglish: "Birth date",
			LanguageRussian: "День рождения",
		},
		keyFieldGender: {
			LanguageEnglish: "Gender",
			LanguageRussian: "Пол",
		},
		keyFieldGenderMale: {
			LanguageEnglish: "Male",
			LanguageRussian: "Мужской",
		},
		keyFieldGenderFemale: {
			LanguageEnglish: "Female",
			LanguageRussian: "Женский",
		},
		keyFieldAddress: {
			LanguageEnglish: "Address",
			LanguageRussian: "Адрес",
		},

		// links
		keyLinkLogin: {
			LanguageEnglish: "Already have an account?",
			LanguageRussian: "Уже есть аккаунт?",
		},
		keyLinkRegister: {
			LanguageEnglish: "Don't have an account? Create one now.",
			LanguageRussian: "Нет аккаунта? Создать учетную запись.",
		},

		// table columns
		keyColumnName: {
			LanguageEnglish: "Name",
			LanguageRussian: "Имя",
		},
		keyColumnEmail: {
			LanguageEnglish: "Email",
			LanguageRussian: "Имейл",
		},
		keyColumnBirthDate: {
			LanguageEnglish: "Birth date",
			LanguageRussian: "День рождения",
		},
		keyColumnGender: {
			LanguageEnglish: "Gender",
			LanguageRussian: "Пол",
		},
		keyColumnAddress: {
			LanguageEnglish: "Address",
			LanguageRussian: "Адрес",
		},

		keyColumnJanuary: {
			LanguageEnglish: "January",
			LanguageRussian: "Январь",
		},
		keyColumnFebruary: {
			LanguageEnglish: "February",
			LanguageRussian: "Февраль",
		},
		keyColumnMarch: {
			LanguageEnglish: "March",
			LanguageRussian: "Март",
		},
		keyColumnApril: {
			LanguageEnglish: "April",
			LanguageRussian: "Апрель",
		},
		keyColumnMay: {
			LanguageEnglish: "May",
			LanguageRussian: "Май",
		},
		keyColumnJune: {
			LanguageEnglish: "June",
			LanguageRussian: "Июнь",
		},
		keyColumnJuly: {
			LanguageEnglish: "July",
			LanguageRussian: "Июль",
		},
		keyColumnAugust: {
			LanguageEnglish: "August",
			LanguageRussian: "Август",
		},
		keyColumnSeptember: {
			LanguageEnglish: "September",
			LanguageRussian: "Сентябрь",
		},
		keyColumnOctober: {
			LanguageEnglish: "October",
			LanguageRussian: "Октябрь",
		},
		keyColumnNovember: {
			LanguageEnglish: "November",
			LanguageRussian: "Ноябрь",
		},
		keyColumnDecember: {
			LanguageEnglish: "December",
			LanguageRussian: "Декабрь",
		},

		// error messages
		KeyErrorEmailTaken: {
			LanguageEnglish: "Email taken",
			LanguageRussian: "Имейл занят",
		},
		KeyErrorCredentialsInvalid: {
			LanguageEnglish: "Credentials invalid",
			LanguageRussian: "Неверный email или пароль",
		},
		KeyErrorDataCollision: {
			LanguageEnglish: "Data collision",
			LanguageRussian: "Коллизия данных",
		},

		// other
		keyNoSearchResults: {
			LanguageEnglish: "No customers",
			LanguageRussian: "Нет клиентов",
		},
	}

	errLangNotSupported = errors.New("target translation language not supported")
	errKeyNotDefined    = errors.New("translation key not defined")
	errLangNotDefined   = errors.New("language not defined for key")
)

func Translate(key Key, language string) (string, error) {
	if strings.TrimSpace(language) == "" || (language != LanguageEnglish && language != LanguageRussian) {
		return "", fmt.Errorf("%w: %s", errLangNotSupported, language)
	}

	m, ok := translations[key]
	if !ok {
		return "", fmt.Errorf("%w: %s", errKeyNotDefined, key)
	}

	res, ok := m[language]
	if !ok {
		return "", fmt.Errorf("%w: %s %s", errLangNotDefined, language, key)
	}

	return res, nil
}
