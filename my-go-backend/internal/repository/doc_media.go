// Package repository مستندسازی لایه ریپازیتوری رسانه
//
// این بسته حاوی ریپازیتوری‌های دسترسی به داده برای موجودیت‌های رسانه است:
//   - MediaFileRepository
//   - MediaVersionRepository
//   - CompressionJobRepository
//   - ...
//
// قوانین کلیدی رعایت شده:
//
//	Rule 0 – Repository Pattern + UnitOfWork
//	------------------------------------------------
//	هیچ منطق تجاری در ریپازیتوری‌ها وجود ندارد؛ تمام تراکنش‌ها توسط UnitOfWork
//	تزریق‌شده مدیریت می‌شود. هر ریپازیتوری *فقط* کوئری‌های داده را اجرا می‌کند.
//
//	Rule 9 – منع مقادیر جادویی
//	----------------------------------
//	مقادیر ثابت (وضعیت job، فرمت فایل، ...) در فایل constants.go نگهداری می‌شود.
//
// امضای متدهای MediaFileRepositoryInterface
//
//	---------------------------------------
//	• Create / Update / Delete – عملیات CRUD پایه
//	• GetBy*                 – کوئری‌های تک مقداری با ایندکس مناسب
//	• BulkCreate/Update      – برای درج دسته‌ای
//	• GetPaginated           – صفحه‌بندی با فیلتر پیشرفته (ساخته در لایه QueryHandler)
//
// نحوهٔ تزریق در UnitOfWork
//
//	------------------------
//	در پیاده‌سازی UnitOfWork، ریپازیتوری‌ها *Lazy-loaded* هستند تا در صورت عدم نیاز
//	ایجاد نشوند:
//
//	   func (u *unitOfWorkImpl) MediaFileRepository() repository.MediaFileRepositoryInterface {
//	       if u.mediaFileRepo == nil {
//	           u.mediaFileRepo = repository.NewMediaFileRepository(u.GetTx())
//	       }
//	       return u.mediaFileRepo
//	   }
//
// پرفورمنس و ایندکس‌ها
// --------------------
// - PreparedStatement: gorm.Config{PrepareStmt:true} فعال است تا هزینهٔ سرور کاهش یابد.
//
// ایندکس‌های وابسته
//
//	----------------
//	متدهای زیر از ایندکس‌های Migration ۴۴/۴۵ استفاده می‌کنند:
//	  • GetByCategory        → idx_media_files_category_created
//	  • GetByUser            → idx_media_files_uploaded_by_created
//
// تست‌پذیری
//
//	-------
//	برای هر ریپازیتوری فایل *_test.go وجود دارد که با SQLite in-memory اجرا می‌شود.
//	این تست‌ها توسط دستور `go test ./internal/...` در Pipeline CI فراخوانی می‌شوند.
//
// راهنما برای افزودن متد جدید
//
//	-------------------------
//	1. نام متد واضح (Rule 7)
//	2. فقط یک کار انجام دهد (Rule 8)
//	3. عدم استفاده از مقادیر جادویی (Rule 9)
//	4. تست واحد اضافه شود (Rule 22)
package repository
