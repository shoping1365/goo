// Package unitofwork مستندسازی لایه مدیریت تراکنش
//
// UnitOfWork
// ----------
//
//	الگوی UnitOfWork تضمین می‌کند تمامی عملیات داده‌ای مرتبط در یک تراکنش اتمیک
//	اجرا شوند. این پیاده‌سازی:
//	  • از GORM استفاده می‌کند.
//	  • تراکنش تو در تو (Nested) را با Savepoint پشتیبانی می‌کند.
//	  • ریپازیتوری‌ها را Lazy Load می‌کند تا کارایی بهبود یابد.
//
//	توالی معمول استفاده:
//	  uow := uowFactory.Create()
//	  err := uow.BeginTx(ctx)
//	  repo := uow.MediaFileRepository()
//	  ...
//	  uow.Commit() // یا uow.Rollback()
//
// UnitOfWorkFactory
// -----------------
//
//	یک factory ساده برای جداسازی تست‌ها از دیتابیس واقعی. در تست‌ها می‌توان یک
//	اتصال SQLite in-memory تزریق کرد تا بدون Postgres تست اجرا شود.
//
// قوانین رعایت-شده:
//   - Rule 1  – آنالیز کامل قبل از اقدام (در مستندات توضیح داده شده)
//   - Rule 3  – جداسازی Handler از Router (Handlerها فقط از Bus استفاده می‌کنند)
//   - Rule 15 – امنیت: هر تراکنش در صورت خطا Rollback می‌شود.
//   - Rule 21 – پرفورمنس: استفاده مجدد از connection و اجتناب از اتصال اضافی.
package unitofwork
