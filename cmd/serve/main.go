package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yangyj/internal/handler"
	"yangyj/pkg/config"
	"yangyj/pkg/i18n"

	// 初始化
	_ "yangyj/pkg/sys"
)

func main() {
	addr := fmt.Sprintf("localhost:%v", config.Config.Port)

	log.Println(i18n.Trans(&i18n.Option{
		ID: "app.listening",
	}), addr)

	log.Println(config.Config)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler.Router(),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println(i18n.Trans(&i18n.Option{
		ID: "app.shutdown",
	}))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(i18n.Trans(&i18n.Option{
			ID: "app.shutdown",
		}), err)
	}
	log.Println(i18n.Trans(&i18n.Option{
		ID: "app.exit",
	}))
}
