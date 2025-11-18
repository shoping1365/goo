// Package cqrs پیاده‌سازی الگوی Command Query Responsibility Segregation برای مدیریت رسانه
//
// CQRS یک الگوی معماری است که عملیات خواندن (Query) و نوشتن (Command) را از هم جدا می‌کند.
// این جداسازی مزایای زیادی دارد:
//
// مزایای CQRS:
//   - جداسازی منطق خواندن از نوشتن
//   - بهینه‌سازی مستقل هر بخش
//   - امکان scale کردن جداگانه read و write
//   - کد خواناتر و قابل نگهداری‌تر
//   - تست‌پذیری بهتر
//
// ساختار پکیج:
//   - commands/: تعریف دستورات (write operations)
//   - queries/: تعریف کوئری‌ها (read operations)
//   - handlers/: پردازنده‌های دستورات و کوئری‌ها
//   - bus/: مسیریابی دستورات و کوئری‌ها به handler مناسب
//
// نحوه استفاده:
//
// ایجاد و پیکربندی CQRS:
//
//	uowFactory := unitofwork.NewUnitOfWorkFactory(db)
//
//	// ایجاد handlers
//	cmdHandler := handlers.NewMediaCommandHandler(uowFactory)
//	queryHandler := handlers.NewMediaQueryHandler(uowFactory)
//
//	// ایجاد buses
//	cmdBus := bus.NewCommandBus(cmdHandler)
//	queryBus := bus.NewQueryBus(queryHandler)
//
// استفاده برای عملیات write:
//
//	cmd := commands.CreateMediaFileCommand{
//	    FileName: "example.jpg",
//	    FileType: "image/jpeg",
//	    Size: 1024000,
//	    URL: "/uploads/example.jpg",
//	}
//	mediaFile, err := cmdBus.CreateMediaFile(ctx, cmd)
//
// استفاده برای عملیات read:
//
//	files, total, err := queryBus.GetMediaFilesPaged(ctx, 1, 20, queries.MediaFileFilter{
//	    FileType: "image/jpeg",
//	    Category: "products",
//	})
//
// پیاده‌سازی handler جدید:
//
// برای اضافه کردن عملیات جدید:
// 1. تعریف command/query در پکیج مربوطه
// 2. اضافه کردن handler method
// 3. اضافه کردن به bus interface
// 4. پیاده‌سازی در bus implementation
//
// مثال command جدید:
//
//	// در commands/media_commands.go
//	type ArchiveMediaFileCommand struct {
//	    ID uint
//	    Reason string
//	}
//
//	// در handlers/command_handlers.go
//	func (h *MediaCommandHandler) HandleArchiveMediaFile(ctx context.Context, cmd commands.ArchiveMediaFileCommand) error {
//	    // پیاده‌سازی
//	}
//
//	// در bus/command_bus.go - interface
//	ArchiveMediaFile(ctx context.Context, cmd commands.ArchiveMediaFileCommand) error
//
//	// در bus/command_bus.go - implementation
//	func (b *commandBusImpl) ArchiveMediaFile(ctx context.Context, cmd commands.ArchiveMediaFileCommand) error {
//	    return b.handler.HandleArchiveMediaFile(ctx, cmd)
//	}
package cqrs
