package mlservice

// ResponseDocRegister представляет информацию о документе
type ResponseDocRegister struct {
	// Uuid - Идентификатор документа в ИНИТ ПРО
	Uuid string `json:"uuid,omitempty"`
	// Time - Время запроса
	Time string `json:"timestamp,omitempty"`
	// Err - Ошибка
	Err ResponseErr `json:"error,omitempty"`
	// Status - Статус запроса
	Status string `json:"status,omitempty"`
}

// ResponseCheckOperationResult представляет результат фискализации
type ResponseCheckOperationResult struct {
	// Uuid - Идентификатор документа в ИНИТ ПРО
	Uuid string `json:"uuid,omitempty"`
	// Status - Статус запроса
	Status string `json:"status,omitempty"`
	// Payload - Данные фискализации
	Payload FnPayload `json:"payload,omitempty"`
	// Warnings - Предупреждения
	Warnings Service `json:"warnings,omitempty"`
	// Time - Время запроса
	Time string `json:"timestamp,omitempty"`
	// Err - Ошибка
	Err ResponseErr `json:"error,omitempty"`
}

// ResponseErr представляет ошибку
type ResponseErr struct {
	// ErrorID - Идентификатор ошибки
	ErrorID string `json:"error_id"`
	// Code - Код ошибки
	Code uint `json:"code"`
	// Text - Текст ошибки
	Text string `json:"text"`
	// Type - Тип ошибки
	Type string `json:"type"`
}

// Receipt представляет чек
type Receipt struct {
	// Client - Данные клиента
	Client Client `json:"client"`
	// Company - Данные компании
	Company Company `json:"company"`
	// Items - Список товаров
	Items []Item `json:"items"`
	// Payments - Список платежей
	Payments []Payment `json:"payments"`
	// Vats - Список НДС
	Vats []Vat `json:"vats"`
	// Total - Общая сумма
	Total float64 `json:"total"`
}

// Client представляет данные клиента
type Client struct {
	// Email - Электронная почта клиента
	Email string `json:"email"`
	// Phone - Телефон клиента
	Phone string `json:"phone"`
}

// Company представляет данные компании
type Company struct {
	// Email - Электронная почта компании
	Email string `json:"email"`
	// Sno - Система налогообложения
	Sno string `json:"sno"`
	// Inn - ИНН компании
	Inn string `json:"inn"`
	// PaymentAddress - Адрес оплаты
	PaymentAddress string `json:"payment_address"`
}

// Item представляет товар
type Item struct {
	// Name - Название товара
	Name string `json:"name"`
	// Price - Цена товара
	Price float64 `json:"price"`
	// Quantity - Количество товара
	Quantity float64 `json:"quantity"`
	// Sum - Сумма за товар
	Sum float64 `json:"sum"`
	// MeasurementUnit - Единица измерения товара
	MeasurementUnit string `json:"measurement_unit,omitempty"`
	// PaymentMethod - Способ оплаты
	PaymentMethod string `json:"payment_method,omitempty"`
	// PaymentObject - Объект оплаты
	PaymentObject string `json:"payment_object,omitempty"`
	// Vat - НДС
	Vat Vat `json:"vat"`
}

// Vat представляет НДС
type Vat struct {
	// Type - Тип НДС
	Type string `json:"type"`
	// Sum - Сумма НДС (опционально)
	Sum float64 `json:"sum,omitempty"`
}

// Payment представляет оплату
type Payment struct {
	// Type - Тип оплаты
	Type int `json:"type"`
	// Sum - Сумма оплаты
	Sum float64 `json:"sum"`
}

// Service представляет сервисные данные
type Service struct {
	// CallbackUrl - URL для обратного вызова
	CallbackUrl string `json:"callback_url"`
}

// FnPayload представляет данные фискализации
type FnPayload struct {
	// FiscalReceiptNumber - Номер фискального чека
	FiscalReceiptNumber int `json:"fiscal_receipt_number"`
	// ShiftNumber - Номер смены
	ShiftNumber int `json:"shift_number"`
	// ReceiptDatetime - Дата и время чека
	ReceiptDatetime string `json:"receipt_datetime"`
	// Total - Общая сумма
	Total float64 `json:"total"`
	// FnNumber - Номер фискального накопителя
	FnNumber string `json:"fn_number"`
	// EcrRegistrationNumber - Регистрационный номер кассы
	EcrRegistrationNumber string `json:"ecr_registration_number"`
	// FiscalDocumentNumber - Номер фискального документа
	FiscalDocumentNumber int `json:"fiscal_document_number"`
	// FiscalDocumentAttribute - Атрибут фискального документа
	FiscalDocumentAttribute int `json:"fiscal_document_attribute"`
	// FnsSite - Сайт ФНС
	FnsSite string `json:"fns_site"`
}

// AgentInfo представляет информацию об агенте
type AgentInfo struct {
	// Type - Тип агента
	Type string `json:"type,omitempty"`
	// PayingAgent - Информация о платёжном агенте
	PayingAgent PayingAgent `json:"paying_agent,omitempty"`
	// ReceivePaymentsOperator - Оператор по приёму платежей
	ReceivePaymentsOperator Operator `json:"receive_payments_operator,omitempty"`
	// MoneyTransferOperator - Оператор по переводу денежных средств
	MoneyTransferOperator MoneyTransferOperator `json:"money_transfer_operator,omitempty"`
	// SupplierInfo - Информация о поставщике
	SupplierInfo SupplierInfo `json:"supplier_info,omitempty"`
}

// PayingAgent представляет информацию о платёжном агенте
type PayingAgent struct {
	// Operation - Операция агента
	Operation string `json:"operation,omitempty"`
	// Phones - Телефоны агента
	Phones []string `json:"phones,omitempty"`
}

// Operator представляет оператора
type Operator struct {
	// Phones - Телефоны оператора
	Phones []string `json:"phones,omitempty"`
}

// MoneyTransferOperator представляет оператора по переводу денежных средств
type MoneyTransferOperator struct {
	// Phones - Телефоны оператора
	Phones []string `json:"phones,omitempty"`
	// Name - Имя оператора
	Name string `json:"name,omitempty"`
	// Address - Адрес оператора
	Address string `json:"address,omitempty"`
	// Inn - ИНН оператора
	Inn string `json:"inn,omitempty"`
}

// SupplierInfo представляет информацию о поставщике
type SupplierInfo struct {
	// Phones - Телефоны поставщика
	Phones []string `json:"phones,omitempty"`
}

// Correction представляет данные коррекции
type Correction struct {
	Company        Company        `json:"company"`           // Данные компании
	CorrectionInfo CorrectionInfo `json:"correction_info"`   // Данные о коррекции
	Payments       []Payment      `json:"payments"`          // Список платежей
	Vats           []Vat          `json:"vats"`              // Список НДС
	Cashier        string         `json:"cashier,omitempty"` // ФИО кассира
}

// CorrectionInfo представляет информацию о коррекции
type CorrectionInfo struct {
	Type       string `json:"type"`        // Тип коррекции
	BaseDate   string `json:"base_date"`   // Дата документа основания
	BaseNumber string `json:"base_number"` // Номер документа основания
	BaseName   string `json:"base_name"`   // Описание коррекции
}
