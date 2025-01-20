package http

import (
	"Hexagonal_go/app/internal/core/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Handler สำหรับจัดการคำขอ HTTP
type Handler struct {
	userService       service.UserService
	landOfficeService service.LandOfficeService
}

// NewHandler สร้าง Handler ใหม่
func NewHandler() *Handler {
	userService := service.NewUserService()
	landOfficeService := service.NewLandOfficeService()                               // แก้ไขการประกาศตัวแปร landOfficeService
	return &Handler{userService: *userService, landOfficeService: *landOfficeService} // ส่งค่าไปยัง Handler
}

// GetUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching users")
	}
	return c.JSON(users)
}

// GetAllLandOffice ดึงข้อมูลสำนักงานที่ดินทั้งหมด
func (h *Handler) GetAllLandOffice(c *fiber.Ctx) error {
	landOffices, err := h.landOfficeService.GetAllLandOffice() // เรียกใช้ landOfficeService
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching land offices")
	}
	return c.JSON(landOffices)
}
