package middleware

import (
	"encoding/json"
	global "server-fiber/model"
	"server-fiber/model/system"
	"server-fiber/utils"
	"strconv"
	"strings"
	"time"

	// json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"

	systemService "server-fiber/service/system"

	"go.uber.org/zap"
)

var operationRecordService = new(systemService.OperationRecordService)

// 写入操作历史

func OperationRecord(c *fiber.Ctx) error {
	var body []byte
	var userId int
	// var err error
	if c.Method() != fiber.MethodGet {
		body = c.Request().Body()
	} else {
		query := c.OriginalURL()
		split := strings.Split(query, "?")
		if len(split) > 1 {
			splitI := strings.Split(split[1], "&")
			m := make(map[string]string)
			for _, v := range splitI {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

	}
	claims, err := utils.GetClaims(c)
	if err != nil {
		return c.Status(403).JSON(map[string]string{"msg": err.Error()})
	}
	// fmt.Printf("%v\n", claims)
	if claims.BaseClaims.ID != 0 {
		userId = int(claims.BaseClaims.ID)
	} else {
		id, err := strconv.Atoi(c.Get("x-user-id"))
		if err != nil {
			userId = 0
		}
		userId = id
	}
	pathURL := c.Path()
	isBackend := system.Backend
	switch {
	case strings.HasPrefix(pathURL, "/backend"):
		isBackend = system.Backend
	case strings.HasPrefix(pathURL, "/api"):
		isBackend = system.Frontend
	case strings.HasPrefix(pathURL, "/mobile"):
		isBackend = system.Mobile
	default:
		isBackend = system.Backend
	}

	record := system.SysOperationRecord{
		Ip:       c.IP(),
		Method:   c.Method(),
		Path:     pathURL,
		Agent:    c.Get("User-Agent"),
		Body:     string(body),
		UserID:   userId,
		TypePort: isBackend,
	}
	// 上传文件时候 中间件日志进行裁断操作
	if strings.Contains(c.Get("Content-Type"), "multipart/form-data") {
		if len(record.Body) > 512 {
			record.Body = "File or Length out of limit"
		}
	}
	defer func() {
		record.Status = c.Response().StatusCode()
		if record.Status == fiber.StatusInternalServerError {
			record.ErrorMessage = string(c.Response().Body())
		} else {
			record.ErrorMessage = ""
		}
		record.Latency = time.Since(time.Now())
		record.Resp = string(c.Response().Body())
		if err := operationRecordService.CreateSysOperationRecord(&record); err != nil {
			global.LOG.Error("create operation record error:", zap.Error(err))
		}
	}()

	return c.Next()
}
